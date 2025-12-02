package main

import "sync"

type Cache struct {
	mu    sync.RWMutex
	items map[int]PostCard
}

func NewCache() *Cache {
	return &Cache{
		items: make(map[int]PostCard),
	}
}

func (c *Cache) Set(card PostCard) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.items[card.ID] = card
}

func (c *Cache) Get(id int) (PostCard, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	card, ok := c.items[id]
	return card, ok
}

func (c *Cache) LoadFromDB(dbCards []PostCard) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for _, card := range dbCards {
		c.items[card.ID] = card
	}
}