package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB
var router *gin.Engine

func start(db *gorm.DB) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("tasks.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
	// migrate(db)
	// seed(db)
}

func getAll(ctxt *gin.Context) {
	db := start(db)
	tasks := getTasks(db)
	ctxt.IndentedJSON(http.StatusOK, tasks)
}
func getFalse(ctxt *gin.Context) {
	db := start(db)
	Ftasks := getTasksFalse(db)
	ctxt.IndentedJSON(http.StatusOK, Ftasks)
}
func create(ctxt *gin.Context) {
	var newTask Task
	fmt.Println("create called ...")
	if err := ctxt.BindJSON(&newTask); err != nil {
		return
	}
	db := start(db)
	createTask(db, newTask)
}
func main() {
	db = start(db)
	router = gin.Default()
	router.GET("/tasks/all/", getAll)
	router.GET("/tasks/uncompleted/", getFalse)
	router.POST("/tasks/create/", create)
	router.Run("localhost:8080")
}
