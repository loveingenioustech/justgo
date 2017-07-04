package main

// go get github.com/garyburd/redigo/redis
import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/garyburd/redigo/redis"
)

var (
	Pool *redis.Pool
)

func init() {
	redisHost := "127.0.0.1:6379"
	Pool = newPool(redisHost)
	close()
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
			return c, err
		},

		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func close() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	signal.Notify(c, syscall.SIGKILL)
	go func() {
		<-c
		Pool.Close()
		os.Exit(0)
	}()
}

func Get(key string) ([]byte, error) {
	conn := Pool.Get()
	defer conn.Close()

	var data []byte
	data, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return data, fmt.Errorf("error get key %s: %v", key, err)
	}
	return data, err
}

// 调用String，将返回结果转成String
func GetString(key string)(string, error){
	conn := Pool.Get()
	defer conn.Close()

	s, err := redis.String(conn.Do("GET", key))

	if err != nil {
		return s, fmt.Errorf("error get key %s: %v", key, err)
	}
	return s, err
}

// 设置
func SetString(key, value string)(interface{}, error){
	conn := Pool.Get()
	defer conn.Close()

	data, err := conn.Do("SET", key, value)

	if err != nil {
		return data, fmt.Errorf("error set key %s: %v", key, err)
	}
	return data, err
}

func main() {
	test, err := Get("test")
	fmt.Println(test, err)

	str, err := GetString("test")
	fmt.Println(str, err)

	setResult, err := SetString("setbygo", "hahaha")
	fmt.Println(setResult)
}
