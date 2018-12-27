package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("mysql", "sample:samplepass@/sample?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Todo{})
}

func main() {

	router := gin.Default()

	v1 := router.Group("/api/v1/todos")
	{
		v1.GET("/", ping)
		v1.POST("/", createTodo)
	}
	router.Run()

}

type Todo struct {
	gorm.Model
	Title     string `json:"title"`
	Completed int    `json:"completed"`
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pondapsodjpaosd",
	})
}

func createTodo(c *gin.Context) {
	fmt.Println("creating!!!")

	type RequestBody struct {
		Title     string `json:"title" binding:"required"`
		Completed int    `json:"completed"`
	}
	var requestBody RequestBody

	if err := c.BindJSON(&requestBody); err != nil {
		fmt.Println("reason", err)
		c.AbortWithStatus(400)
		return
	}

	todo := Todo{Title: requestBody.Title, Completed: requestBody.Completed}

	db.NewRecord(todo)
	db.Create(&todo)

	// completed, _ := strconv.Atoi(c.PostForm("completed"))
	// fmt.Println("completed", completed)
	// todo := Todo{Title: c.PostForm("title"), Completed: completed}
	// db.Save(&todo)
	// c.JSON(http.StatusCreated, gin.H{
	// 	"status":     http.StatusCreated,
	// 	"message":    "Todo Item Created Successfully",
	// 	"resourceId": todo.ID,
	// })

	c.JSON(http.StatusOK, gin.H{
		"message": "OK!",
	})
}
