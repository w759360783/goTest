package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"strconv"
	"time"
)

func main() {
	pool := &redis.Pool{
		Dial: func() (conn redis.Conn, e error) {
			return redis.Dial("tcp", "127.0.0.1:6379")
		},                             // 拨号函数
		MaxActive:   1000,             // 最大活动数
		MaxIdle:     20,               // 最大闲置数
		IdleTimeout: time.Second * 60, // 最大等待连接时间
	}
	defer pool.Close() // 延迟关闭连接
	for i := 0; i < 10; i++ {
		go func(pool *redis.Pool, i int) {
			conn := pool.Get()
			defer conn.Close() // 延迟关闭连接
			reply, err := conn.Do("set", "conn"+strconv.Itoa(i), i)
			strings, err := redis.String(reply, err)
			fmt.Println(strings)
		}(pool, i)
	}
	time.Sleep(time.Second * 3) // 防止主线程直接退出
}
