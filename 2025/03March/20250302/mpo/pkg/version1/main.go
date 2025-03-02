package main

import (
	"fmt"
	"sync"
	"unsafe"
)

// 每个内存块的大小
const (
	blockSize = 64
	pageSize  = 256
	blockNum  = pageSize / blockSize
)

// MemPool 内存池结构
type MemPool struct {
	mutex     sync.Mutex
	free      *Block // 空闲链表
	pageCount int    // 记录当前总分配数
	pageLimit int    // 池子最大容量
}

type Page struct {
	next     *Page
	freeList *Block
	Begin    *Block
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
	}
	for i := 0; i < blockNum; i++ {
		NewBlock := &Block{}
		if i == 0 {
			NewPage.Begin = NewBlock
		}
		NewBlock.page = NewPage
		NewBlock.next = unsafe.Pointer(NewPage.freeList)
		NewPage.freeList = NewBlock
	}
	return NewPage
}

// NewMemPool 创建新的内存池
func NewMemPool(size, limit int) *MemPool {
	mp := &MemPool{
		free:      nil,
		mutex:     sync.Mutex{},
		pageLimit: limit,
	}
	for i := 0; i < size; i++ {
		page := AllocPage()
		page.Begin.next = unsafe.Pointer(mp.free)
		mp.free = page.freeList
		mp.pageCount++
	}
	return mp
}

// Alloc 分配内存
func (mp *MemPool) Alloc() *Block {
	mp.mutex.Lock()
	defer mp.mutex.Unlock()

	if mp.free == nil {
		if mp.pageCount == mp.pageLimit {
			fmt.Printf("\n alloc failed %v,%v", mp.pageCount, mp.pageLimit)
			return nil
		}
		fmt.Printf("\n alloc success %v,%v", mp.pageCount, mp.pageLimit)
		newpage := AllocPage()
		mp.pageCount++
		mp.free = newpage.freeList
	}

	block := mp.free
	mp.free = (*Block)(block.next)
	fmt.Printf("\n auto alloc success")
	return block
}

// Free 释放内存
func (mp *MemPool) Free(block *Block) {
	if block == nil {
		fmt.Printf("\n nil pointer forbidden")
		return
	}
	mp.mutex.Lock()
	defer mp.mutex.Unlock()

	block.next = unsafe.Pointer(mp.free)
	mp.free = block
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

	fmt.Printf("\nAllocated blocks:\n %p\n %p\n %p \n %p\n %p \n %p \n %p \n %p", b1, b2, b3, b4, b5, b6, b7, b8)

	//释放
	//pool.Free(b1)
	//pool.Free(b2)
	//pool.Free(b3)
	//pool.Free(b4)
	//
	////重新分配，应该复用刚刚释放的块
	//a1 := pool.Alloc()
	//a2 := pool.Alloc()
	//a3 := pool.Alloc()
	//
	//fmt.Printf("\nReallocated blocks:\n %p \n %p \n %p", a1, a2, a3)
}
