package cart

import (
	"github.com/garyburd/redigo/redis"
)

type cartRespository struct {
	db   string
	conn *redis.Conn
}

type Repository struct {
	Store(cart *Cart) error
	Find(id ID) (*Cart, error)
	AddItem(id string, item *CartItem) error
}


// Cart is a collection of the items and how to store it
type Cart struct {
	items   map[string]*CartItem
	storage cartRespository
}

type mutator func(*CartItem)

// NewCartRespository returns a new instance of a redis cargo respository
func NewCartRespository(db string, conn *redis.Conn) (cartRespository, error) {

	r := &cartRespository{
		db:   db,
		conn: conn,
	}

	return r, nil
}

// Add will add and item to the cart
func (c *Cart) Add(id string, name string, price float64, q int, attrs map[string]interface{}) *CartItem {
	c.items[id] = &CartItem{
		ID:         id,
		Name:       name,
		Price:      price,
		Quantity:   q,
		Attributes: attrs,
	}

	c.storage.Save(c.items)

	return c.items[id]
}

// Remove will remove an item from the existing cart if it exists
func (c *Cart) Remove(id string) bool {
	if _, exists := c.items[id]; exists {
		delete(c.items, id)
		c.storage.Save(c.items)

		return true
	}

	return false
}

// IsEmpty checks if there is no items in the cart
func (c *Cart) IsEmpty() bool {
	if (len(c.items)) > 0 {
		return false
	}
	return true
}

// GetContent for mutator approaches
func (c *Cart) GetContent() map[string]*CartItem {
	return c.items
}

// EachItem mutates each item in the bucket as the mutator passed to the
// method
func (c *Cart) EachItem(callback mutator) {

	for _, item := range c.items {

		// Execute mutator passed to Each method
		callback(item)
	}

	c.storage.Save(c.items)
}
