package kredis

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/spf13/viper"
)

// Pool *redis.Pool
var (
	Pool *redis.Pool
)

// NewConfig generates a mysql configuration object which
// This mysql instance will becomes the single source of
// truth for the app configuration.
func NewConfig() *redis.Pool {
	conn := getConnection()
	Pool = &redis.Pool{
		// Maximum number of idle connections in the pool.
		MaxIdle:     40,
		IdleTimeout: 1000 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", conn)
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
	return Pool
}

// GetConn Get Pool
func GetConn() *redis.Pool {
	return Pool
}

func getConnection() string {
	host := viper.GetString("APP.DB.REDIS.HOST")
	port := viper.GetString("APP.DB.REDIS.PORT")
	conn := fmt.Sprintf("%s:%s", host, port)
	return conn
}

// Set executes the redis SET command
func Set(k string, v interface{}) error {
	c, err := Pool.Dial()
	if err != nil {
		return err
	}
	_v, err := json.Marshal(v)
	if err != nil {
		return err
	}
	_, err = c.Do("SET", k, _v)
	if err != nil {
		return err
	}
	return nil
}

// Get executes the redis GET command
func Get(k string) (v interface{}, err error) {
	c, err := Pool.Dial()
	if err != nil {
		return nil, err
	}
	s, err := redis.String(c.Do("GET", k))
	if err != nil {
		return s, err
	}
	err = json.Unmarshal([]byte(s), &v)
	if err != nil {
		return s, err
	}
	return s, err
}
