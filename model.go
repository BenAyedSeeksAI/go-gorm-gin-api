package main

import (
	"gorm.io/gorm"
)

type standardizer interface {
}
type Task struct {
	gorm.Model
	Id        int    `gorm:"primaryKey"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(&Task{})
}
func seed(db *gorm.DB) {
	db.Create(&Task{Id: 1, Name: "Go to the university ..", Completed: false})
	db.Create(&Task{Id: 2, Name: "contact the recruiter", Completed: false})
	db.Create(&Task{Id: 3, Name: "get the diploma ..", Completed: true})
	db.Create(&Task{Id: 4, Name: "give him a statement", Completed: true})
	db.Create(&Task{Id: 5, Name: "go abroad ..", Completed: false})
}
func getTasks(db *gorm.DB) []Task {
	var tasks []Task
	db.Find(&tasks)
	return tasks
}
func getTasksFalse(db *gorm.DB) []Task {
	var tasks []Task
	db.Where("completed = ?", "0").Find(&tasks)
	return tasks
}
func getTaskById(db *gorm.DB, id int) Task {
	var task = Task{Id: id}
	db.First(&task)
	return task
}
