package redis

import (
	"errors"
	"github.com/garyburd/redigo/redis"
	"github.com/Unknwon/goconfig"
	"time"
)

func init() {
	c, err := goconfig.LoadConfigFile("conf/conf.ini")
	if err != nil {
		panic(err)
	}

	ip, err := c.GetValue("Redis", "address")
	if err != nil {
		panic(err)
	}

	Redis = &myRedis{newPool(ip)}
}

var (
	Redis   *myRedis
	NotFind = errors.New("Not Find")
)

type myRedis struct {
	*redis.Pool
}

func (r *myRedis) Get(key string) (string, error) {
	conn, err := r.Dial()
	if err != nil {
		return "", err
	}
	result, err := conn.Do("GET", key)
	var s []byte
	if err == nil {
		if result == nil {
			return "", NotFind
		}
		s = result.([]byte)
	}
	return string(s), err
}

func (r *myRedis) Set(key, value string) error {
	conn, err := r.Dial()
	if err != nil {
		return err
	}
	_, err = conn.Do("SET", key, value)
	return err
}

func (r *myRedis) Del(key string) error {
	conn, err := r.Dial()
	if err != nil {
		return err
	}
	_, err = conn.Do("DEL", key)
	return err
}

func newPool(server string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			//			if _, err := c.Do("AUTH", password); err != nil {
			//				c.Close()
			//				return nil, err
			//			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}
