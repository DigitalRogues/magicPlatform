package main

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func dbSet(json []byte) {
	c, err := redis.Dial("tcp", "redis:6379")
	if err != nil {
		fmt.Println(err)
	}
	defer c.Close()
	n, err := c.Do("SET", "magic-data", json)
	if err != nil {
		fmt.Printf("redis SET error:%v \n", err)
	} else {
		fmt.Printf("redis SET response:%v \n", n)
	}
}

func dbGet() {
	c, err := redis.Dial("tcp", "redis:6379")
	if err != nil {
		fmt.Println(err)
	}
	defer c.Close()

	n, err := redis.Bytes(c.Do("GET", "magic-data"))
	if err != nil {
		fmt.Printf("redis GET error:%v \n", err)
	} else {
		var st Magic
		json.Unmarshal(n, &st)
		fmt.Printf("redis GET response:%+v \n", st)
	}
}
