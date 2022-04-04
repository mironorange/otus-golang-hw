package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
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
	isExists := false
	queueItem, ok := cache.items[key]
	if ok {
		isExists = true
		item := queueItem.Value
		value = item.(*cacheItem).value
		cache.queue.MoveToFront(queueItem)
	}
	return value, isExists
}

func (cache *lruCache) Set(key Key, value interface{}) bool {
	queue := cache.queue
	isExists := false
	queueItem, ok := cache.items[key]
	if ok {
		isExists = true
		queue.Remove(queueItem)
	}
	item := &cacheItem{
		key:   key,
		value: value,
	}
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
	cache.items[key] = queueItem
	return isExists
}

func (cache *lruCache) Clear() {
	for i := cache.queue.Front(); i != nil; i = i.Next {
		key := i.Value.(*cacheItem).key
		delete(cache.items, key)
	}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
