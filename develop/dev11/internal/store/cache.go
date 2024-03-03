package store

import (
	"errors"
	"sync"

	"github.com/ositlar/go-http/internal/models"
)

var Ch = NewStore()

type Cache struct {
	data  map[string]models.Event
	mutex sync.Mutex
}

func NewStore() *Cache {
	return &Cache{
		data: make(map[string]models.Event),
	}
}

func (c *Cache) NewEvent(e models.Event) {
	Ch.mutex.Lock()
	Ch.data[e.Id] = e
	Ch.mutex.Unlock()
}

func (c *Cache) UpdateEvent(e models.Event) error {
	Ch.mutex.Lock()
	_, ok := Ch.data[e.Id]
	if !ok {
		return errors.New("there's no event to update")
	}
	Ch.data[e.Id] = e
	Ch.mutex.Unlock()
	return nil
}
