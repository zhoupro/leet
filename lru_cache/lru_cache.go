package main

func main() {
	obj := Constructor(2)
	obj.Put(1, 1)
	obj.Put(2, 2)
	_ = obj.Get(1)
	obj.Put(3, 3)
	_ = obj.Get(2)
	obj.Put(4, 4)

}

//local end

type DLinkedList struct {
	key  int
	val  int
	prev *DLinkedList
	next *DLinkedList
}

type LRUCache struct {
	capacity int
	head     *DLinkedList
	tail     *DLinkedList
	cache    map[int]*DLinkedList
}

func Constructor(capacity int) LRUCache {
	head := &DLinkedList{}
	tail := &DLinkedList{}
	head.next = tail
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*DLinkedList),
		head:     head,
		tail:     tail,
	}
}

func (lru *LRUCache) Get(key int) int {
	element, ok := lru.cache[key]

	if !ok {
		return -1
	}

	lru.removeFromChain(element)
	lru.addToChain(element)

	return element.val
}

func (lru *LRUCache) Put(key int, value int) {
	list, ok := lru.cache[key]

	if !ok {
		list = &DLinkedList{key: key, val: value}
		lru.cache[key] = list
	} else {
		list.val = value

		lru.removeFromChain(list)
	}

	lru.addToChain(list)

	if len(lru.cache) > lru.capacity {
		last := lru.tail.prev

		lru.removeFromChain(last)

		delete(lru.cache, last.key)
	}
}

func (lru *LRUCache) addToChain(list *DLinkedList) {
	list.next = lru.head.next
	lru.head.next = list
	list.prev = lru.head
	list.next.prev = list
}

func (lru *LRUCache) removeFromChain(list *DLinkedList) {
	list.prev.next = list.next
	list.next.prev = list.prev
}
