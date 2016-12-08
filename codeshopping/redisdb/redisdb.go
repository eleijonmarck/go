package redisdb

import (
	"encoding/json"
	"fmt"
	"github.com/eleijonmarck/codeshopping/cart"
	"github.com/garyburd/redigo/redis"
)

type cartRepository struct {
	db   string
	Pool *redis.Pool
}

func (r *cartRepository) Store(cart *cart.Cart) error {

	// makes a copy of the struct that is pointed to by the pointer
	c := r.Pool.Get()
	defer c.Close()
	serialized, err := json.Marshal(&cart)
	if err != nil {
		// error handle
		panic(err)
	}
	_, err2 := c.Do("SET", cart.CartID, string(serialized))
	if err2 != nil {
		panic(err2)
	}
	defer c.Close()
	return err
}

func (r *cartRepository) Find(key string) (*cart.Cart, error) {
	c := r.Pool.Get()
	defer c.Close()
	values, err := redis.Bytes(c.Do("GET", key))
	carty := cart.Cart{}
	err2 := json.Unmarshal(values, &carty)
	if err2 != nil {
		//
		fmt.Println("lol")
	}
	return &carty, err
}

func (r *cartRepository) FindAll() []*cart.Cart {
	c := r.Pool.Get()
	// TODO: return slices of bytes and return them
	keys, err := redis.Strings(c.Do("KEYS", "test*"))
	if err != nil {
		// handle it
	}
	var result = make([]*cart.Cart, len(keys))
	carts, err2 := redis.ByteSlices(c.Do("MGET", keys))
	fmt.Println(carts)
	if err2 != nil {
		//
		panic(err)
	}
	for i := 0; i < len(carts); i++ {
		json.Unmarshal(carts[i], result[i])
	}
	return result
}

// NewCartRepository creates a repository for storage of the carts
func NewCartRepository(db string, pool *redis.Pool) (cart.Repository, error) {
	r := &cartRepository{
		db:   db,
		Pool: pool,
	}
	return r, nil
}
