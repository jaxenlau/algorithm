package _46_lru

import (
	"container/list"
	"sync"
)

// [146] 运用你所掌握的数据结构，设计和实现一个 LRU (最近最少使用) 缓存机制。它应该支持以下操作： 获取数据 get 和 写入数据 put 。
//
// 获取数据 get(key) - 如果密钥 (key) 存在于缓存中，则获取密钥的值（总是正数），否则返回 -1。
// 写入数据 put(key, value) - 如果密钥不存在，则写入其数据值。当缓存容量达到上限时，它应该在写入新数据之前删除最近最少使用的数据值，从而为新的数据值留出空间。
//
// 进阶:
//
// 你是否可以在 O(1) 时间复杂度内完成这两种操作？
//
// 示例:
//
// LRUCache cache = new LRUCache( 2 /* 缓存容量 */ );
//
// cache.put(1, 1);
// cache.put(2, 2);
// cache.get(1);       // 返回  1
// cache.put(3, 3);    // 该操作会使得密钥 2 作废
// cache.get(2);       // 返回 -1 (未找到)
// cache.put(4, 4);    // 该操作会使得密钥 1 作废
// cache.get(1);       // 返回 -1 (未找到)
// cache.get(3);       // 返回  3
// cache.get(4);       // 返回  4

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
