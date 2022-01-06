package list

import "container/list"

type kv struct {
	// key is only needed to delete entry in data map
	// when removing LRU item
	k int
	v int
}

type LRUCache struct {
	capacity int
	data     map[int]*list.Element
	hits     *list.List
}

func NewLRU(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		data:     make(map[int]*list.Element),
		hits:     list.New(),
	}
}

func (c *LRUCache) Get(key int) int {
	if el, ok := c.data[key]; ok {
		c.hits.MoveToFront(el)
		return el.Value.(kv).v
	}
	return -1
}

func (c *LRUCache) Put(key int, value int) {
	// just an update?
	if el, ok := c.data[key]; ok {
		el.Value = kv{k: key, v: value}
		c.hits.MoveToFront(el)
		return
	}
	// full?
	if c.hits.Len() == c.capacity {
		last := c.hits.Back()
		delete(c.data, last.Value.(kv).k)
		c.hits.Remove(last)
	}
	// put
	c.data[key] = c.hits.PushFront(kv{k: key, v: value})
}
