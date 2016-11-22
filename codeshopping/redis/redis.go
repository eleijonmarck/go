package redis

import (
	"encoding/json"
	"errors"
	"github.com/eleijonmarck/codeshopping/cart"
	"github.com/garyburd/redigo/redis"
)

type cartRepository struct {
	db   string
	conn *redis.Conn
}

func (r *cartRepository) Store(cart *cart.Cart) error {

	serialized, err := json.Marshal(&cart)
	if err != nil {
		// error handle
	}
	r.conn.DO("SET", cart.CartID, string(serialized))
}

var (
	ErrNotFound = errors.New("not found")
)

// DB is the interface to a simple key/value store
type DB interface {
	// Get returns the value for the given key, ErrNotFound if the key doesn't exist,
	// or another error if the get failed
	Get(key string) ([]byte, error)
	// Set sets the value for the given key. Returns an error if the set failed.
	// If non-nil error is returned, the value was not updated
	Set(key string, val []byte) error
}

func NewCartRepository(db string, conn *redis.Conn) (cart.Repository, err) {
}
