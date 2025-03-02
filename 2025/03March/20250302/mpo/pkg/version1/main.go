package main

import (
	"fmt"
	"sync"
	"unsafe"
)

// 每个内存块的大小
const (
	blockSize      = 64
	pageSize       = 256
	blockNum       = pageSize / blockSize
	freePageMaxNum = 0
)

// MemPool 内存池结构
type MemPool struct {
	mutex       sync.Mutex
	free        *Block // 空闲链表
	pageCount   int    // 记录当前总分配数
	pageLimit   int    // 池子最大容量
	pages       []*Page
	freePage    *Page
	freePageNum int
}

type Page struct {
	next     *Page
	freeList *Block
	blocks   [blockNum]Block
	used     int
}

// Block 内存块
type Block struct {
	next unsafe.Pointer
	data [blockSize]byte
	page *Page
}

// AllocPage 分配页
func AllocPage() *Page {
	NewPage := &Page{
		next:     nil,
		freeList: nil,
		used:     0,
	}
	for i := 0; i < blockNum; i++ {
		if i < blockNum-1 {
			//维护链表
			NewPage.blocks[i].next = unsafe.Pointer(&NewPage.blocks[i+1])
		} else {
			//如果此时是页中最后一个块
			NewPage.blocks[i].next = nil
		}
		//保持映射关系
		NewPage.blocks[i].page = NewPage
	}
	//链表头指向头，总体是这样的结构
	//头->next->next->nil
	NewPage.freeList = &NewPage.blocks[0]
	return NewPage
}

// NewMemPool 创建新的内存池
func NewMemPool(size, limit int) *MemPool {
	mp := &MemPool{
		free:      nil,
		mutex:     sync.Mutex{},
		pageLimit: limit,
		pages:     make([]*Page, 0, limit),
		freePage:  nil,
	}
	for i := 0; i < size; i++ {
		//此时开辟给定的页空间大小
		mp.pages = append(mp.pages, AllocPage())
		//同样维护页的链表结构
		mp.freePage = mp.pages[i]
		//维护页链表内部的内存block链表结构
		mp.pages[i].blocks[blockNum-1].next = unsafe.Pointer(mp.free)
		if i > 0 {
			//防止数组越界
			mp.pages[i].next = mp.pages[i-1]
		}
		//也是更新内存block的链表结构
		mp.free = mp.pages[i].freeList
		mp.freePageNum++
		mp.pageCount++
	}

	return mp
}

// Alloc 分配内存
func (mp *MemPool) Alloc() *Block {
	mp.mutex.Lock()
	defer mp.mutex.Unlock()

	if mp.free == nil {
		//如果超出了最大分配页的限制
		if mp.pageCount == mp.pageLimit {
			fmt.Printf("\n alloc failed %v,%v", mp.pageCount, mp.pageLimit)
			return nil
		}
		fmt.Printf("\n alloc success %v,%v", mp.pageCount, mp.pageLimit)
		//分配新的页
		newpage := AllocPage()
		//如果此时内存池的页数量大于0，更新刚分配的页的next元素，维护链表
		if len(mp.pages) > 0 {
			newpage.next = mp.pages[mp.pageCount-1]
		}
		//将页加入内存池的页切片中
		mp.pages = append(mp.pages, newpage)
		//更新空闲页的数量
		mp.freePageNum++
		//更新页的数量
		mp.pageCount++
		//将页中的区块加入空闲内存块链表
		mp.free = newpage.freeList
		fmt.Printf("\n 分配了新的页， 空闲页增加， %v", mp.freePageNum)
	}
	//从空闲块中取出新区块
	block := mp.free
	//更新原链表的结构
	mp.free = (*Block)(block.next)
	block.next = nil
	if block.page.used == 0 {
		mp.freePageNum--
		fmt.Printf("\n此时页被分配第一个block，空闲页减少%v", mp.freePageNum)
	}
	block.page.used++
	fmt.Printf("\n auto alloc success")
	return block
}

// Free 释放内存
func (mp *MemPool) Free(block *Block) {
	mp.mutex.Lock()
	defer mp.mutex.Unlock()
	if block == nil {
		fmt.Printf("\n nil pointer forbidden")
		return
	}
	block.page.used--
	if block.page.used == 0 {
		mp.freePageNum++
		fmt.Printf("\n释放页中block，此时空闲页增加,%v", mp.freePageNum)
		//释放页
		mp.releasePage(block.page)
		return
	}
	block.next = unsafe.Pointer(mp.free)
	mp.free = block
}

func (mp *MemPool) releasePage(page *Page) {
	// 清理 free 链表，防止已释放 Page 的 Block 仍然存在于 free
	var prev, curr *Block
	curr = mp.free
	for curr != nil {
		next := (*Block)(curr.next)
		if curr.page == page {
			fmt.Printf("\nfree中存在原来的区块%p", curr)
			if prev != nil {
				prev.next = curr.next
			} else {
				mp.free = next
			}
		} else {
			prev = curr
		}
		curr = next
	}
	//如果超出了最大空闲页的限制，则从 pages 数组里移除 Page
	if mp.freePageNum > freePageMaxNum {
		mp.freePageNum--
		fmt.Printf("\n此时空闲页超出范围，回收空闲页, %v", mp.freePageNum)
		var prevPage *Page = nil
		for i, p := range mp.pages {
			if p == page {
				if prevPage != nil {
					prevPage.next = page.next
				}
				mp.pages[i] = nil
				if len(mp.pages)-1 > i {
					mp.pages = append(mp.pages[:i], mp.pages[i+1:]...)
				} else {
					mp.pages = mp.pages[:i]
				}
				mp.pageCount--
				fmt.Printf("\n page %p released", page)
				break
			}
			prevPage = p
		}
	}
}

func main() {
	pool := NewMemPool(1, 2)

	// 分配 3 个块
	b1 := pool.Alloc()
	b2 := pool.Alloc()
	b3 := pool.Alloc()
	b4 := pool.Alloc()
	b5 := pool.Alloc()
	b6 := pool.Alloc()
	b7 := pool.Alloc()
	b8 := pool.Alloc()
	b9 := pool.Alloc()
	fmt.Printf("\n%v", pool.pages)
	fmt.Printf("\nAllocated blocks:\n %p\n %p\n %p \n %p\n %p \n %p \n %p \n %p \n %p", b1, b2, b3, b4, b5, b6, b7, b8, b9)

	//释放
	pool.Free(b1)
	pool.Free(b2)
	pool.Free(b3)
	pool.Free(b4)
	pool.Free(b5)
	pool.Free(b6)
	pool.Free(b7)
	pool.Free(b8)
	fmt.Printf("\n%v", pool.pages)
	//此处无法实现真正的释放，只是逻辑的释放，还是c语言更合适干这些...
	//重新分配，应该复用刚刚释放的块
	a1 := pool.Alloc()
	a2 := pool.Alloc()
	a3 := pool.Alloc()
	a4 := pool.Alloc()

	fmt.Printf("\nReallocated blocks:\n %p \n %p \n %p \n %p", a1, a2, a3, a4)
}
