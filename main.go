package main

import (
	"fmt"

	"webframeworkV2.0/app/http"
	"webframeworkV2.0/framework"
)

func main() {

	// 初始化服务容器
	container := framework.NewWebContainer()

	engine, err := http.NewHttpEngine()
	if err != nil {
		// 报错
		fmt.Printf("NewHttpEngine error %v", err)
	}

	container.Bind(engine)

}
