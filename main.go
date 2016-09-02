package main

import (
	"log"
	"time"

	"github.com/Louch2010/gochat-server/gui"
)

func init() {
	log.Println("初始化Gochat服务端...")
}

func main() {
	startTime := time.Now().Unix()
	log.Println("启动中...")

	gui.Init()

	log.Println("启动完成，耗时：", time.Now().Unix()-startTime)
}
