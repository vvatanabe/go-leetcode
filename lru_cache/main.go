package main

import "container/list"

type LRUCache struct {
	capacity int
	ll       *list.List
	cache    map[interface{}]*list.Element
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		ll:       list.New(),
		cache:    make(map[interface{}]*list.Element),
	}
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.cache[key]; ok {
		this.ll.MoveToFront(v)
		return v.Value.(*entry).value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if v, ok := this.cache[key]; ok {
		this.ll.MoveToFront(v)
		v.Value.(*entry).value = value
		return
	}
	v := this.ll.PushFront(&entry{key, value})
	this.cache[key] = v
	if this.ll.Len() > this.capacity {
		v := this.ll.Back()
		this.ll.Remove(v)
		kv := v.Value.(*entry)
		delete(this.cache, kv.key)
	}
}

type entry struct {
	key   int
	value int
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
