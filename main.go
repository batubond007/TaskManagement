package main

import (
	"TaskManagement/Controllers"
	"TaskManagement/Models"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	Models.ConnectDatabase()
	router.GET("/tasks", Controllers.GetAllTasks)
	router.GET("/tasks/GetById/:id", Controllers.GetTaskById)
	router.GET("/tasks/GetByType/:type", Controllers.GetTasksByType)
	router.POST("/tasks", Controllers.CreateTask)
	router.PATCH("/tasks/UpdateType/:id", Controllers.UpdateTaskType)
	router.PATCH("/tasks/UpdateTitle/:id", Controllers.UpdateTaskTitle)
	router.PATCH("/tasks/UpdateContext/:id", Controllers.UpdateTaskContext)
	router.DELETE("/tasks/DeleteTask/:id", Controllers.DeleteTask)
	router.Run("localhost:8080")
}
