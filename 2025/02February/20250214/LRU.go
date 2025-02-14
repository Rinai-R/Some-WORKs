package main

import "container/list"

type LRUCache struct {
	capacity int
	kv       map[int]*list.Element
	KeyList  *list.List
}

type KV struct {
	Key   int
	Value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		kv:       make(map[int]*list.Element),
		KeyList:  list.New(),
	}
}

func (this *LRUCache) Get(key int) int {
	if elem, ok := this.kv[key]; ok {
		this.KeyList.MoveToFront(elem)
		return elem.Value.(KV).Value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node := this.kv[key]; node != nil {
		node.Value = KV{
			Key:   key,
			Value: value,
		}
		this.KeyList.MoveToFront(node)
		return
	}
	if this.KeyList.Len() >= this.capacity {
		oldest := this.KeyList.Back()
		this.KeyList.Remove(oldest)
		delete(this.kv, oldest.Value.(KV).Key)
	}
	KeyValue := KV{Key: key, Value: value}
	elem := this.KeyList.PushFront(KeyValue)
	this.kv[key] = elem
}
