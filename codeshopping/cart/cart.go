package cart

import (
	"github.com/pborman/uuid"
)

// Cart is a collection of the items and how to store it
type Cart struct {
	Items map[string]*CartItem
	CartID    string
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
	if (len(c.Items)) > 0 {
		return false
	}
	return true
}

// GetContent for mutator approaches
func (c *Cart) GetContent() map[string]*CartItem {
	return c.Items
}

func New(id string) *Cart {
	return &Cart {
		CartId : id,
		Items : []map[string]*CartItem,
	}
}

// NextTrackingID generates a new tracking ID.
// TODO: Move to infrastructure(?)
func NewCartID() string {
	return strings.Split(strings.ToUpper(uuid.New()), "-")[0]
}

// Repository provides access to a cart storage
type Repository interface {
	Store(cart *Cart) error
	Find(id string) (*Cart, error)
	FindAll() []*CartItem
}
