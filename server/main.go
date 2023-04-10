package main

import (
	"blog/router"
	"log"

	"golang.org/x/sync/errgroup"
)

var g errgroup.Group

func main() {
	// 初始化全局变量
	router.InitGlobalVariable()

	// 前台接口服务
	g.Go(func() error {
		return router.FrontendServer().ListenAndServe()
	})

	// 后台接口服务
	g.Go(func() error {
		return router.BackendServer().ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
