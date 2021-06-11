package gomutimod

import (
	"fmt"
	_ "github.com/zhangxu-ai/gomutimod/mods/a"
	_ "github.com/zhangxu-ai/gomutimod/mods/b"
)

func init() {
	fmt.Println("hello mutimod")
}

func Say() {
	fmt.Println("hello")
}
