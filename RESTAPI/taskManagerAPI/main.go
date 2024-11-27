package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Task represents a task with its properties.
type Task struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	Status      string    `json:"status"`
}

// Mock data for tasks
var tasks = []Task{
	{ID: "1", Title: "Task 1", Description: "First task", DueDate: time.Now(), Status: "Pending"},
	{ID: "2", Title: "Task 2", Description: "Second task", DueDate: time.Now().AddDate(0, 0, 1), Status: "In Progress"},
	{ID: "3", Title: "Task 3", Description: "Third task", DueDate: time.Now().AddDate(0, 0, 2), Status: "Completed"},
}

func main() {
	router := gin.Default()                      // create the router
	router.GET("/ping", func(ctx *gin.Context) { //
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/tasks", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"tasks": tasks,
		})
	})

	router.GET("/tasks/:id", func(ctx *gin.Context) {
		id := ctx.Param("id");
		
		for _, task := range tasks{
			if task.ID == id{
				ctx.JSON(http.StatusOK, task)
				return
			}
		}

		ctx.JSON(http.StatusNotFound, gin.H{
			"error" : "Task not found",
		})
	})

	router.PUT("/tasks/:id", func(ctx *gin.Context){
		id := ctx.Param("id");

		var updatedTask Task;
		if err := ctx.ShouldBindJSON(&updatedTask); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H {"error" : err.Error()})
			return
		}

		for i, task := range tasks{
			if task.ID == id {
				if updatedTask.Title != "" {
					tasks[i].Title = updatedTask.Title
				}
				if updatedTask.Description != "" {
					tasks[i].Description = updatedTask.Description
				}

				ctx.JSON(http.StatusOK, gin.H {
					"message" : "Task Updated",
				})
				return
			}
		}

		ctx.JSON(http.StatusNotFound, gin.H {"message" : "Task not found"})
	})

	router.Run()
}
