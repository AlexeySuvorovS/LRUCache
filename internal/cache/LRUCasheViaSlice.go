package cache

import "sync"

type LRUSlice struct {
	capacity int
	queue    []Key
	cache    map[Key]Val
}

var muS sync.RWMutex

func NewLRUSlice(capacity int) *LRUSlice {
	return &LRUSlice{
		capacity: capacity,
		cache:    make(map[Key]Val, capacity),
		queue:    make([]Key, capacity),
	}
}

func (lru *LRUSlice) exists(key Key) (int, bool) {
	idx := 0
	isFound := false

	for i, keyQ := range lru.queue {
		if keyQ == key {
			idx = i
			isFound = true
		}
	}

	return idx, isFound
}

func (lru *LRUSlice) Get(key Key) (Val, bool) {
	val, isFound := lru.cache[key]

	return val, isFound
}

func (lru *LRUSlice) Add(key Key, val Val) {
	muS.Lock()
	ePrev := lru.queue[0]
	idx, found := lru.exists(key)

	if !found { // если элемент не в списке
		// сдвигаем всю очередь на 1
		for i := 1; i < len(lru.queue); i++ {
			delete(lru.cache, lru.queue[lru.capacity-1])
			curr := lru.queue[i]
			lru.queue[i] = ePrev
			ePrev = curr
		}
	} else { // если элемент в списке
		//сдвигаем часть очереди на 1 и обновляем элемент в map
		for i := 1; i <= idx; i++ {
			curr := lru.queue[i]
			lru.queue[i] = ePrev
			ePrev = curr
		}
	}

	// вставляем элемент в начало
	lru.queue[0] = key
	lru.cache[key] = val
	muS.Unlock()
}

func (lru *LRUSlice) Len() int {
	lenQ := 0
	for _, v := range lru.queue {
		if v != nil {
			lenQ++
		} else {
			break
		}
	}

	return lenQ
}

func (lru *LRUSlice) GetQueue() []Key {
	var res []Key

	for _, v := range lru.queue {
		if v != nil {
			res = append(res, v)
		} else {
			break
		}
	}

	return res
}
