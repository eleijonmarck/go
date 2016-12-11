package cart

import (
	"github.com/pborman/uuid"
	"strings"
)

// Cart is a collection of the items and how to store it
type Cart struct {
	Items  map[string]*CartItem
	CartID string
}

// AddItem will add and item to the cart
func (c *Cart) AddItem(id string, name string, price float64, q int, attrs map[string]interface{}) *CartItem {
	c.Items[id] = &CartItem{
		ID:         id,
		Name:       name,
		Price:      price,
		Quantity:   q,
		Attributes: attrs,
	}

	// add to database

	return c.Items[id]
}

// Remove will remove an item from the existing cart if it exists
func (c *Cart) RemoveItem(id string) bool {
	if _, exists := c.Items[id]; exists {
		delete(c.Items, id)

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

// New returns a new Cart
func New(id string) *Cart {
	m := make(map[string]*CartItem)
	return &Cart{
		CartID: id,
		Items:  m,
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
	FindAll() []*Cart
}
