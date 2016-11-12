package cart

import (
	"github.com/eleijonmarck/codeshopping/item"
)

type Cart struct {
	items   map[string]*CartItem
	storage CartBucket
}

type mutator func(*CartItem)

func Boot(storage CartBucket) (*Cart, error) {
	restored, err := storage.Restore()

	if err != nil {
		return nil, err
	}

	bucket := &Cart{
		items:   restored,
		storage: storage,
	}

	return bucket, nil
}

// A cart will have many items on it and this add to it
func (c *Cart) Add(id, name string, price float64, q int, attrs map[string]interface{}) *Cart {
	c.items[id] = &CartItem{
		Id:         id,
		Name:       name,
		Price:      price,
		Quantity:   q,
		Attributes: attrs,
	}

	c.storage.Save(c.items)

	return c.items[id]
}

// Remove, will remove an item from the existing cart if it exists
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
	} else {
		return true
	}
}
