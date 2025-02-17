package sync

import (
	"internal/race"
	"sync/atomic"
	"unsafe"
)

type Mutex struct {
	// [阻塞的goroutine个数, starving标识, woken标识, locked标识]
	// [0~28, 1, 1, 1]
	state int32
	sema  uint32
}

const (
	mutexLocked      = 1 << iota // mutex is locked
	mutexWoken                   // 唤醒标记
	mutexStarving                // 饥饿模式
	mutexWaiterShift = iota      // 位移数

	starvationThresholdNs = 1e6 // 阻塞时间阀值1ms
)

func (m *Mutex) Lock() {
	// Fast path: grab unlocked mutex.
	// 尝试CAS上锁
	if atomic.CompareAndSwapInt32(&m.state, 0, mutexLocked) {
		if race.Enabled {
			race.Acquire(unsafe.Pointer(m))
		}
		// 上锁成功，直接返回
		return
	}

	var waitStartTime int64
	starving := false
	awoke := false
	iter := 0
	old := m.state
	for {

		// 进入到这个循环的，有两种角色goroutine
		// 一种是新来的goroutine。另一种是被唤醒的goroutine。所以它们可能在这个地方再一起竞争锁
		// 如果新来的goroutine抢成功了，那另一个只能再阻塞着等待。但超过1ms后，锁会转换成饥饿模式
		// 在这个模式下，所有新来的goroutine必须排在队伍的后面。没有抢锁资格

		// 饥饿模式下，不能自旋
		// 锁被占用了，不能自旋
		if old&(mutexLocked|mutexStarving) == mutexLocked && runtime_canSpin(iter) {
			// woken位没有被设置；被阻塞等待goroutine的个数大于0
			if !awoke && old&mutexWoken == 0 && old>>mutexWaiterShift != 0 &&
				atomic.CompareAndSwapInt32(&m.state, old, old|mutexWoken) {
				// 可以自旋了，那就设置上woken位，在unlock时，如果发现有别的goroutine在自旋，就立即返回，有被阻塞的goroutine也不唤醒了
				awoke = true
			}
			// runtime_doSpin -> sync_runtime_doSpin
			// 每次自旋30个时钟周期，最多120个周期
			runtime_doSpin()
			iter++
			old = m.state
			continue
		}

		// 自旋完了还是等不到锁 或 可以上锁

		new := old
		// 饥饿模式下的锁不抢
		if old&mutexStarving == 0 {
			// 非饥饿模式下，可以抢锁
			new |= mutexLocked
		}
		if old&(mutexLocked|mutexStarving) != 0 {
			// 已经被上锁了，或锁处于饥饿模式下，就阻塞当前的goroutine
			new += 1 << mutexWaiterShift
		}
		if starving && old&mutexLocked != 0 {
			// 当前的goroutine已经被饿着了，所以要把锁设置为饥饿模式
			new |= mutexStarving
		}
		if awoke {
			// 当前的goroutine有自旋过，但现在已经自旋结束了。所以要取消woken模式
			if new&mutexWoken == 0 {
				panic("sync: inconsistent mutex state")
			}
			// 取消woken标志
			new &^= mutexWoken
		}
		if atomic.CompareAndSwapInt32(&m.state, old, new) {
			if old&(mutexLocked|mutexStarving) == 0 {
				// 成功上锁
				break // locked the mutex with CAS
			}

			// 主要是为了和第一次调用的Lock的g划分不同的优先级
			queueLifo := waitStartTime != 0
			if waitStartTime == 0 {
				waitStartTime = runtime_nanotime()
			}
			// 使用信号量阻塞当前的g
			// 如果当前g已经阻塞等待过一次了，queueLifo被赋值true
			// 即，陷入睡眠状态
			runtime_SemacquireMutex(&m.sema, queueLifo)
			// 判断当前g是否被饿着了
			starving = starving || runtime_nanotime()-waitStartTime > starvationThresholdNs
			old = m.state
			if old&mutexStarving != 0 {
				// 饥饿模式下，被手递手喂信号量唤醒的
				if old&(mutexLocked|mutexWoken) != 0 || old>>mutexWaiterShift == 0 {
					panic("sync: inconsistent mutex state")
				}
				delta := int32(mutexLocked - 1<<mutexWaiterShift) // -7(111)
				if !starving || old>>mutexWaiterShift == 1 {
					// 退出饥饿模式
					// 饥饿模式会影响自旋
					delta -= mutexStarving
				}
				atomic.AddInt32(&m.state, delta)
				break
			}
			// 不是手递手的信号量，那就自己继续竞争锁
			// 必须设置为true，这样新一轮的CAS之前，就可以取消woken模式。
			// 因为通过信号量释放锁时，为了保持公平性，会同时设置woken模式。
			awoke = true
			iter = 0
		} else {
			old = m.state
		}
	}

	if race.Enabled {
		race.Acquire(unsafe.Pointer(m))
	}
}

func (m *Mutex) Unlock() {
	if race.Enabled {
		_ = m.state
		race.Release(unsafe.Pointer(m))
	}

	// Fast path: drop lock bit.
	new := atomic.AddInt32(&m.state, -mutexLocked)
	if (new+mutexLocked)&mutexLocked == 0 {
		// 不能多次执行unclock()
		panic("sync: unlock of unlocked mutex")
	}
	if new&mutexStarving == 0 {
		// 非饥饿模式
		old := new
		for {
			// 没有被阻塞的goroutine。直接返回
			// 有阻塞的goroutine，但处于woken模式，直接返回
			// 有阻塞的goroutine，但被上锁了。可能发生在此for循环内，第一次CAS不成功。因为CAS前可能被新的goroutine抢到锁。直接返回
			// 有阻塞的goroutine，但锁处于饥饿模式。可能发生在被阻塞的goroutine不是被唤醒调度的，而是被正常调度运行的。直接返回
			if old>>mutexWaiterShift == 0 || old&(mutexLocked|mutexWoken|mutexStarving) != 0 {
				return
			}

			// 有阻塞的goroutine，唤醒一个或变为没有阻塞的goroutine了就退出
			// 这个被唤醒的goroutine还需要跟新来的goroutine竞争
			// 如果只剩最后一个被阻塞的goroutine。唤醒它之后，state就变成0。
			// 如果此刻来一个新的goroutine抢锁，它有可能在goroutine被重新调度之前抢锁成功。
			// 这样就失去公平性了，不能让它那么干，所以这里也要设置为woken模式。
			// 因为Lock方法开始的fast path，CAS操作的old值是0。这里设置woken模式成功后，后来者就只能乖乖排队。保持了锁的公平性
			new = (old - 1<<mutexWaiterShift) | mutexWoken
			if atomic.CompareAndSwapInt32(&m.state, old, new) {
				runtime_Semrelease(&m.sema, false)
				return
			}
			old = m.state
		}
	} else {
		// 饥饿模式
		// 手递手唤醒一个goroutine
		runtime_Semrelease(&m.sema, true)
	}
}
