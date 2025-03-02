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
	pages     []*Page
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
			NewPage.blocks[i].next = unsafe.Pointer(&NewPage.blocks[i+1])
		} else {
			NewPage.blocks[i].next = nil
		}
		NewPage.blocks[i].page = NewPage
	}
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
	}
	for i := 0; i < size; i++ {
		mp.pages = append(mp.pages, AllocPage())
		mp.pages[i].blocks[blockNum-1].next = unsafe.Pointer(mp.free)
		if i > 0 {
			mp.pages[i].next = mp.pages[i-1]
		}
		mp.free = mp.pages[i].freeList
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
		if len(mp.pages) > 0 {
			newpage.next = mp.pages[mp.pageCount-1]
		}
		mp.pages = append(mp.pages, newpage)
		mp.pageCount++
		mp.free = newpage.freeList
	}
	block := mp.free
	mp.free = (*Block)(block.next)
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

	// 从 pages 数组里移除 Page
	var prevPage *Page = nil
	for i, p := range mp.pages {
		if p == page {
			if prevPage != nil {
				prevPage.next = page.next
			}
			mp.pages = append(mp.pages[:i], mp.pages[i+1:]...)
			mp.pageCount--
			fmt.Printf("\n page %p released", page)
			break
		}
		prevPage = p
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

	fmt.Printf("\nAllocated blocks:\n %p\n %p\n %p \n %p\n %p \n %p \n %p \n %p", b1, b2, b3, b4, b5, b6, b7, b8)

	//释放
	pool.Free(b1)
	pool.Free(b2)
	pool.Free(b3)
	pool.Free(b4)

	//重新分配，应该复用刚刚释放的块
	a1 := pool.Alloc()
	a2 := pool.Alloc()
	a3 := pool.Alloc()
	a4 := pool.Alloc()

	fmt.Printf("\nReallocated blocks:\n %p \n %p \n %p \n %p", a1, a2, a3, a4)
}
