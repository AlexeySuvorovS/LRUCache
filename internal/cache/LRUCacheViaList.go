package cache

import (
	"container/list"
	"sync"
)

var muL sync.RWMutex

type LRUList struct {
	capacity int
	queue    list.List
	cache    map[Key]Val
}

func NewLRUList(capacity int) *LRUList {
	return &LRUList{
		capacity: capacity,
		queue:    list.List{},
		cache:    make(map[Key]Val, capacity),
	}
}

func (lru *LRUList) Get(key Key) (Val, bool) {
	val, isFound := lru.cache[key]
	return val, isFound
}

func (lru *LRUList) Add(key Key, val Val) {
	_, ok := lru.cache[key]

	if ok {
		// элемент в списке, перемещаем его в начало
		for e := lru.queue.Back(); e != nil; e = e.Prev() {
			if e.Value == key {
				muL.Lock()
				lru.queue.MoveToBack(e)
				lru.cache[key] = val
				muL.Unlock()
				break
			}
		}
	} else {
		// элемент не в списке, вставляем в начало
		lru.queue.PushBack(key)
		lru.cache[key] = val

		// если выходим за размер, удаляем последний элемент
		if lru.queue.Len() > lru.capacity {
			lru.deleteLast()
		}
	}
}

func (lru *LRUList) deleteLast() {
	muL.Lock()
	lastE := lru.queue.Front()
	delete(lru.cache, lastE.Value)
	lru.queue.Remove(lastE)
	muL.Unlock()
}

func (lru *LRUList) Len() int {
	return lru.queue.Len()
}

func (lru *LRUList) GetQueue() []Key {

	m := min(lru.Len())
	res := make([]Key, m)

	i := 0
	for e := lru.queue.Front(); e != nil; e = e.Next() {
		res[i] = e.Value
		i++
	}

	return res
}
