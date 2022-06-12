package cache

import (
	"L0/pkg/model"
	"errors"
	"sync"
)

type Cache struct {
	data map[string]model.Order
	mu   sync.RWMutex
}

// NewCache возвращает новую структуру с инициализированным кэшом
func NewCache() *Cache {
	return &Cache{data: make(map[string]model.Order), mu: sync.RWMutex{}}
}

// Get возвращает заказ из кэша по его id
func (c *Cache) Get(id string) (model.Order, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	order, ok := c.data[id]
	if !ok {
		return model.Order{}, errors.New("the order isn't in cache")
	}

	return order, nil
}

// Put добавляет заказ в кэш
func (c *Cache) Put(order model.Order) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[order.OrderUID] = order
}

// IsExist проверяет наличие заказа в кэше по его id
func (c *Cache) IsExist(id string) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()

	_, ok := c.data[id]

	return ok
}
