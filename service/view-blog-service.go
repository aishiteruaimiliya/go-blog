package service

import (
	"github.com/garyburd/redigo/redis"
)

func AddViewTimes(bid string) (int, error) {
	conn := GetRedisConn()
	defer conn.Close()
	return redis.Int(conn.Do("incr", bid))
}

func GetViewTimes(bid string) (int, error) {
	conn := GetRedisConn()
	defer conn.Close()
	return redis.Int(conn.Do("get", bid))
}
