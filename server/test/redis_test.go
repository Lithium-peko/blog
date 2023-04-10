package test

import (
	"blog/utils"
	"fmt"
)

func ExampleClient() {
	rdb := utils.InitRedis()
	fmt.Println("redis success: ", rdb)
}
