package redis

import (
	"encoding/json"

	"github.com/gomodule/redigo/redis"
)

type IRedisClient interface {
	SetDataRedis(key string, value interface{}, duration int64) error
	GetDataRedis(key string) ([]byte, error)
}

type RedisClent struct {
	conn redis.Conn
}

func NewRedisClient(conn redis.Conn) IRedisClient {
	return &RedisClent{
		conn: conn,
	}
}

func (r *RedisClent) SetDataRedis(key string, value interface{}, duration int64) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}

	if duration == 0 {
		if _, err := r.conn.Do("SET", key, data); err != nil {
			return err
		}
	} else {
		if _, err := r.conn.Do("SETEX", key, duration, data); err != nil {
			return err
		}
	}
	return nil
}

func (r *RedisClent) GetDataRedis(key string) ([]byte, error) {
	dataRedis, err := redis.Bytes(r.conn.Do("GET", key))
	if err != nil {
		return nil, err
	}
	return dataRedis, nil
}
