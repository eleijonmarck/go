package redis

import (
	"github.com/garyburd/redigo/redis"
)

type cartRepository struct {
	db string
	session *redigo.Session
}

}
