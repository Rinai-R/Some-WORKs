### **核心逻辑与设计**
1. **竞态检测处理**：
    - 当竞态检测启用时（`race.Enabled`），`Add` 方法会处理与 `Wait` 的同步：
        - **减少计数时**（`delta < 0`）：通过 `race.ReleaseMerge` 确保递减操作与 `Wait` 正确同步。
        - **增减计数时**：临时禁用竞态检测（`race.Disable`），并在函数返回时重新启用（`defer race.Enable()`）。

2. **原子更新计数器与等待数**：
    - `state` 是 `uint64` 类型，高 32 位为计数器（`v`），低 32 位为等待的 goroutine 数量（`w`）。
    - 通过原子操作 `wg.state.Add(uint64(delta) << 32)` 更新计数器，确保并发安全。

3. **错误检查**：
    - **计数器为负**：若 `v < 0`，触发 `panic("negative WaitGroup counter")`。
    - **并发调用 `Add` 和 `Wait`**：若 `w > 0`（已有等待者）且 `delta > 0`（增加计数）且 `v == delta`（计数器从 0 开始增加），触发 `panic("Add called concurrently with Wait")`。

4. **唤醒等待者**：
    - 当计数器归零（`v == 0`）且存在等待者（`w > 0`）时：
        - 检查状态是否被篡改（`wg.state.Load() != state`），防止并发滥用。
        - 重置 `state` 为 0，并通过 `runtime_Semrelease` 逐个释放信号量，唤醒所有等待的 goroutine。

---

### **关键代码段解析**
```go
// 更新计数器与等待者数量
state := wg.state.Add(uint64(delta) << 32)
v := int32(state >> 32)  // 高32位为计数器
w := uint32(state)       // 低32位为等待者数量

// 首次递增需同步Wait（竞态检测逻辑）
if race.Enabled && delta > 0 && v == int32(delta) {
    race.Read(unsafe.Pointer(&wg.sema))
}

// 错误处理：计数器为负
if v < 0 { panic("negative counter") }

// 错误处理：Add与Wait并发调用
if w != 0 && delta > 0 && v == int32(delta) { 
    panic("Add called concurrently with Wait") 
}

// 计数器未归零或无等待者，直接返回
if v > 0 || w == 0 { return }

// 唤醒所有等待者前检查状态一致性
if wg.state.Load() != state { 
    panic("Add called concurrently with Wait") 
}

// 重置状态并释放信号量
wg.state.Store(0)
for ; w != 0; w-- {
    runtime_Semrelease(&wg.sema, false, 0)
}
```

---

### **并发场景与错误处理**
1. **计数器管理**：
    - `Add(1)` 在启动新 goroutine **前** 调用，`Done()`（即 `Add(-1)`）在任务完成时调用。
    - 计数器归零时唤醒所有 `Wait` 阻塞的 goroutine。

2. **错误使用**：
    - **负数计数器**：多次调用 `Done()` 或 `Add` 负数导致 `v < 0`。
    - **并发调用 `Add` 和 `Wait`**：在 `Wait` 执行期间调用 `Add` 增加计数，违反“所有 `Add` 需在 `Wait` 前完成”的约定。

3. **状态一致性检查**：
    - 唤醒等待者前，验证 `state` 未被并发修改，避免数据竞争。

---

### **信号量机制**
- `runtime_Semrelease` 和 `runtime_Semacquire` 是 Go 运行时内部的信号量操作。
- `Wait` 调用 `Semacquire` 阻塞，直到 `Add` 在计数器归零时调用 `Semrelease` 唤醒。

---

### **总结**
- **原子操作**：通过 `state` 的原子读写，保证计数器与等待数的线程安全。
- **错误防御**：严格检查并发滥用，触发 panic 防止未定义行为。
- **同步机制**：信号量实现高效阻塞与唤醒，确保 `Wait` 与 `Add` 正确协作。

这段代码通过精细的状态管理和错误检查，实现了 `WaitGroup` 的核心功能：协调多个 goroutine 的完成事件。