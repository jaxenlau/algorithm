package lru

import (
	"container/list"
	"sync"
)

/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU缓存机制
 */

type Node struct {
	key int
	val int
}

// LRUCache 缓存对象结构
type LRUCache struct {
	mutex    sync.Mutex
	capacity int
	hashMap  *sync.Map
	nodeList *list.List
}

// Constructor 构建 LRU 缓存对象
func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		hashMap:  new(sync.Map),
		nodeList: list.New(),
	}
}

// Get LRU 缓存获取 key 值
func (lru *LRUCache) Get(key int) int {
	val, ok := lru.hashMap.Load(key)
	if !ok {
		return -1
	}

	elem := val.(*list.Element)

	lru.mutex.Lock()
	defer lru.mutex.Unlock()
	lru.nodeList.MoveToFront(elem)

	return elem.Value.(*Node).val
}

// Put key，value 写入 LRU 缓存
func (lru *LRUCache) Put(key int, value int) {

	val, ok := lru.hashMap.Load(key)
	if ok {
		elem := val.(*list.Element)
		node := elem.Value.(*Node)
		node.key = key
		node.val = value
		lru.mutex.Lock()
		defer lru.mutex.Unlock()
		lru.nodeList.MoveToFront(elem)
		lru.hashMap.Store(key, elem)
		return
	}

	if lru.IsFull() {
		lru.Flush()
	}

	lru.mutex.Lock()
	defer lru.mutex.Unlock()
	elem := lru.nodeList.PushFront(&Node{key: key, val: value})
	lru.hashMap.Store(key, elem)
}

// Capacity 缓存容量
func (lru *LRUCache) Capacity() int {
	return lru.capacity
}

// IsFull 缓存容量
func (lru *LRUCache) IsFull() bool {
	return lru.nodeList.Len() == lru.Capacity()
}

// Flush 缓存删除链表尾部元素
func (lru *LRUCache) Flush() {
	elem := lru.nodeList.Back()
	if elem == nil {
		return
	}
	lru.mutex.Lock()
	defer lru.mutex.Unlock()
	lru.nodeList.Remove(elem)
	lru.hashMap.Delete(elem.Value.(*Node).key)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
