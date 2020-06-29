package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"os"
)

func HandleError(err error, when string) {
	if err != nil {
		fmt.Println(err, when)
		os.Exit(1)
	}
}

func main() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	HandleError(err, "redis.Dial")
	reply, err := conn.Do("get", "name")
	HandleError(err, "connDo")
	fmt.Printf("type = %T,value=%v\n", reply, reply)
	nameStr, _ := redis.String(reply, err)
	fmt.Println(nameStr)
}
