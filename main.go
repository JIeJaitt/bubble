package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Todo Model
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func main() {
	r := gin.Default()

	r.Static("/static", "static")
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	v1Group := r.Group("v1")

	// 代办事项
	// 添加
	v1Group.POST("todo", func(c *gin.Context) {

	})

	// 查看
	v1Group

	// 修改
	// 删除

	r.Run()
}
