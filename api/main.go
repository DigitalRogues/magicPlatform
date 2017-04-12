package main

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	rC, err := redis.Dial("tcp", "redis:6379")
	if err != nil {
		fmt.Println(err)
	}
	defer rC.Close()

	r := gin.Default()

	// gin.H is a shortcut for map[string]interface{}
	r.GET("/", func(c *gin.Context) {
		//var data []byte
		data := getData(rC)
		//jString := fmt.Sprintf("%s", data)

		c.IndentedJSON(http.StatusOK, data)
	})
	// gin.H{"message": "hey", "status": http.StatusOK}
	// Listen and serve on 0.0.0.0:8080
	r.Run(":8888")
}

func getData(rC redis.Conn) interface{} {

	n, err := redis.Bytes(rC.Do("GET", "magic-data"))
	if err != nil {
		fmt.Printf("redis GET error:%v \n", err)
	}
	var v interface{}
	json.Unmarshal(n, &v)
	fmt.Printf("redis GET response:%v \n", v)
	return &v
}
