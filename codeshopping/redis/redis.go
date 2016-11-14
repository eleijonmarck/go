package redis

import (
	"github.com/garyburd/redigo/redis"
)

type cartRepository struct {
	db   string
	conn *redis.Conn
}

//func (r *cartRepository) Store(cart *cart.Cart) error {
//	conn := r.conn.Do("SET", cart.Cart.
//}
