package service

import (
	"blog/model/blogs"
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

func TokenToRedis(token,aid string,expire time.Duration)error{
	conn := GetRedisConn()
	defer conn.Close()
	r, err := redis.String(conn.Do("set", token, aid,"ex",expire.Seconds()))
	if r=="OK"{
		return nil
	}
	return err
}

func RedisToToken(token string)(string,error){
	conn:=GetRedisConn()
	defer conn.Close()
	r,err:=redis.String(conn.Do("get",token))
	return r,err
}

func SendMsg(receiver string,msg blogs.Message)error {
	conn:=GetRedisConn()
	defer conn.Close()
	msgjson, err := json.Marshal(msg)
	if err != nil {
		fmt.Println(err)
		return err
	}
	r,err:=redis.String(conn.Do("lpush",receiver,msgjson))
	if r!="OK"{
		return err
	}
	return nil
}

func RecvMsg(receiver string,msgs *[]blogs.Message)error{
	conn:=GetRedisConn()
	defer conn.Close()
	r,err:=redis.String(conn.Do("rpop" ,receiver))
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = json.Unmarshal([]byte(r), msgs)
	if err != nil {
		return err
	}
	return nil
}