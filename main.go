package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("启动服务...")
	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20

	r.Run()
}
