package main

import (
	"log"
	"os"
	_ "github.com/Lovelcp/go-in-action-learning/chapter2/matchers"
	"github.com/Lovelcp/go-in-action-learning/chapter2/search"
)

// init 在 main之前调用
func init() {
	log.SetOutput(os.Stdout)
}

// main 是整个程序的入口
func main() {
	// 使用特定的项做搜索
	search.Run("president")
}