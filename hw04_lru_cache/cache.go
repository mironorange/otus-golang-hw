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

// Получить ранее установленное значение кеша и флаг его присутствия
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

// Установить значение кеша
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

// Очистить кеш
func (cache *lruCache) Clear() {
	item := cache.queue.Front()
	if item == nil {
		return
	}
	for i := item; i != nil; i = i.Next {
		key := i.Value.(*cacheItem).key
		delete(cache.items, key)
	}
}

// Инициировать новое значение кеша
func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewDoublyLinkedList(capacity),
		items:    make(map[Key]*ListItem, capacity),
	}
}
