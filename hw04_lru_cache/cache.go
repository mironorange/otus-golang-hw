package hw04lrucache

import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	mutex    sync.Mutex
	capacity int
	queue    List
	items    map[Key]*ListItem
}

type cacheItem struct {
	key   Key
	value interface{}
}

func (cache *lruCache) Get(key Key) (interface{}, bool) {
	var value interface{}

	cache.mutex.Lock()
	isExists := false
	queueItem, ok := cache.items[key]
	if ok {
		isExists = true
		item := queueItem.Value
		value = item.(*cacheItem).value
		cache.queue.MoveToFront(queueItem)
	}
	cache.mutex.Unlock()
	return value, isExists
}

func (cache *lruCache) Set(key Key, value interface{}) bool {
	cache.mutex.Lock()
	queue := cache.queue
	isExists := false
	item := &cacheItem{
		key:   key,
		value: value,
	}
	queueItem, ok := cache.items[key]
	if ok {
		isExists = true
		queueItem.Value = item
		queue.MoveToFront(queueItem)
	} else {
		// Если операция добавления элемента приведет к переполнению списка
		// Перед тем как добавлять новый элемент необходимо
		// Вытолкнуть последний элемент списка
		if queue.Len() >= cache.capacity {
			back := queue.Back()
			if back != nil {
				key := back.Value.(*cacheItem).key
				queue.Remove(back)
				delete(cache.items, key)
			}
		}
		queueItem = queue.PushFront(item)
	}
	cache.items[key] = queueItem
	cache.mutex.Unlock()
	return isExists
}

func (cache *lruCache) Clear() {
	cache.mutex.Lock()
	for i := cache.queue.Front(); i != nil; i = i.Next {
		key := i.Value.(*cacheItem).key
		delete(cache.items, key)
	}
	cache.mutex.Unlock()
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
