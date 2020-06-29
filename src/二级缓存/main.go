package main

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"os"
)

/**
√ 实现二级缓存
√ 程序运行起来之后,提示"请输入命令:",输入getall,查询并显示所有人员信息
√ 第一次时查询mysql并将结果缓存在redis,设置60秒的过期时间
√ 以后的每次查询,如果redis有数据就从redis加载,没有则重复上一步的操作
 */

type People struct {
	Id       int     `db:"id"`       // id
	Name     string  `db:"name"`     // 姓名
	Age      int     `db:"age"`      // 年龄
	Rmb      float64 `db:"rmb"`      // 资产
	Gender   bool    `db:"gender"`   // 性别
	Birthday string  `db:"birthday"` // 生日
}

/*
封装的错误处理的方法
*/
func HandlerError(err error, when string) {
	if err != nil {
		fmt.Println(err, when)
		os.Exit(1)
	}
}
func main() {
	var cmd string
project:
	for {
		fmt.Println("请输入命令:")
		fmt.Scan(&cmd)
		switch cmd {
		case "getAll":
			//fmt.Println("getAll")
			getPeople()
		case "exit":
			break project
		default:
			fmt.Println("命令输入错误，请重新输入")
		}
	}
	fmt.Println("GAME OVER(ಥ_ಥ) ")
}

// 获取People的数据信息
func getPeople() {
	// 优先在缓存中读取需要的信息
	peopleStr := cacheGetPeople()
	// 判断缓存中是否有需要的信息
	if peopleStr == nil || len(peopleStr) <= 0 {
		// 当前缓存中没有缓存我们需要的信息,需要从mysql数据库中读取我们需要的信息然后缓存在redis当中
		people := getPeople2Mysql()
		//打印读出来的数据
		fmt.Println("mysql读出来的数据：", people)
		//将从mysql中读出来的数据缓存在redis当中以方便下次读取实现二级缓存
		cacheRedis(&people)

	} else {
		//缓存中缓存了我们需要的信息
		fmt.Println("读取到缓存的信息：", peopleStr)
	}

}

// 读取redis中的缓存数据
func cacheGetPeople() (people []string) {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	defer conn.Close()
	HandlerError(err, "cacheGerPeople:redis.Dial")
	reply, err := conn.Do("lrange", "peoples", "0", "-1")
	people, err = redis.Strings(reply, err)
	HandlerError(err, "cacheGerPeople:redis.Strings")
	return
}
func getPeople2Mysql() (p []People) {
	db, err := sqlx.Open("mysql", "root:123@tcp(localhost:3306)/study")
	defer db.Close()
	HandlerError(err, "getPeople2Mysql:sqlx.Open")
	err = db.Select(&p, "select * from person")
	HandlerError(err, "getPeople2Mysql:db.Select")
	return
}

/**
将从mysql中读取的数据缓存在redis当中
 */
func cacheRedis(p *[]People) {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	defer conn.Close()
	HandlerError(err, "cacheRedis:redis.Dial")
	_, err = conn.Do("del", "peoples")
	HandlerError(err, "cacheRedis:conn.Do:del")
	for _, people := range *p {
		bytes, err := json.Marshal(people)
		HandlerError(err, "cacheRedis:json.Marshal")
		_, err = conn.Do("rpush", "peoples", string(bytes))
		HandlerError(err, "cacheRedis:conn.Do"+string(bytes))
	}
	_, err = conn.Do("expire", "peoples", 60)
	HandlerError(err, "cacheRedis:conn.Do:expire")
	fmt.Println("缓存成功！！！")
}
