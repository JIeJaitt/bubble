## 未采用逻辑分层的代码
```go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

var (
	DB *gorm.DB
)

// Todo Model
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func initMySQL() (err error) {
	dsn := "root:12345678@tcp(127.0.0.1:3306)/bubble?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return
	}
	return DB.DB().Ping()
}

func main() {
	err := initMySQL()
	if err != nil {
		panic(err)
	}
	defer DB.Close()

	DB.AutoMigrate(&Todo{})

	r := gin.Default()

	r.Static("/static", "static")
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	v1Group := r.Group("v1")

	// ADD
	v1Group.POST("todo", func(c *gin.Context) {
		var todo Todo
		c.BindJSON(&todo)

		if err = DB.Create(&todo).Error; err != nil {
			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, todo)
		}
	})

	// SEARCH
	v1Group.GET("/todo", func(c *gin.Context) {
		var todoList []Todo
		if err = DB.Find(&todoList).Error; err != nil {
			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, todoList)
		}
	})

	// SEARCH A
	v1Group.GET("/todo:id", func(c *gin.Context) {

	})

	// UPDATE
	v1Group.PUT("/todo/:id", func(c *gin.Context) {
		id, ok := c.Params.Get("id")
		if !ok {
			c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
			return
		}

		var todo Todo
		if err = DB.Where("id=?", id).First(&todo).Error; err != nil {
			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			return
		}
		c.BindJSON(&todo)
		if err = DB.Save(&todo).Error; err != nil {
			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, todo)
		}
	})

	// DELETE
	v1Group.DELETE("/todo/:id", func(c *gin.Context) {
		id, ok := c.Params.Get("id")
		if !ok {
			c.JSON(http.StatusOK, gin.H{"error": "无效的 id"})
			return
		}
		if err = DB.Where("id=?", id).Delete(Todo{}).Error; err != nil {
			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{id: "deleted"})
		}

	})

	r.Run()
}

```