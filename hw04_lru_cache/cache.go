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

func (c *lruCache) Clear() {
	c.queue = NewList()
	c.items = make(map[Key]*ListItem, c.capacity)
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	// Если элемент присутствует в словаре
	if val, ok := c.items[key]; ok {
		c.queue.MoveToFront(val)
		li := c.queue.Front()
		c.items[key] = li
		val.Value.(*cacheItem).value = value
		return true
	}

	// Если элемента нет и достигли ёмкости кэша - выталкиваем последний элемент
	if c.queue.Len() == c.capacity {
		if back := c.queue.Back(); back != nil {
			c.queue.Remove(back)
			delete(c.items, back.Value.(*cacheItem).key)
		}
	}

	// Добавляем новый элемент
	item := &cacheItem{
		key:   key,
		value: value,
	}
	li := c.queue.PushFront(item)
	c.items[key] = li
	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	if val, ok := c.items[key]; ok {
		c.queue.MoveToFront(val)
		li := c.queue.Front()
		c.items[key] = li
		return li.Value.(*cacheItem).value, true
	}
	return nil, false
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
