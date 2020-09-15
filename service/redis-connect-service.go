package service

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

var client redis.Pool

func Init() {
	client = redis.Pool{
		MaxIdle:   5,
		MaxActive: 1000,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "127.0.0.1:6379")
		},
	}
	test()
}

func test() {
	conn := client.Get()
	defer conn.Close()
	_, err := conn.Do("set", "key", "hhhh")
	if err != nil {
		panic(err)
	}
	fmt.Println("redis ready.....")
}
func GetRedisConn() redis.Conn {
	return client.Get()
}
