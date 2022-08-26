package Controllers

import (
	Models "TaskManagement/Models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GET

func GetAllTasks(context *gin.Context) {
	var tasks []Models.Task
	Models.DB.Find(&tasks)
	context.JSON(http.StatusOK, gin.H{"data": tasks})
}

func GetTaskById(context *gin.Context) {
	var task Models.Task
	err := Models.DB.Where("id = ?", context.Param("id")).First(&task).Error
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"ERROR": "Record not found!"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": task})
}

func GetTasksByType(context *gin.Context) {
	var tasks []Models.Task
	err := Models.DB.Find(&tasks, "type = ?", context.Param("type")).Error
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"ERROR": "Record not found!"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": tasks})
}

// POST

type CreateTaskInput struct {
	Type    string `json:"type" binding:"required"`
	Title   string `json:"title" binding:"required"`
	Context string `json:"context" binding:"required"`
}

func CreateTask(context *gin.Context) {
	// Validate Input
	var input CreateTaskInput
	err := context.ShouldBindJSON(&input)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"ERROR: ": err.Error()})
		return
	}

	// Create task
	task := Models.Task{Type: input.Type, Title: input.Title, Context: input.Context}
	Models.DB.Create(&task)

	context.JSON(http.StatusOK, gin.H{"data": task})
}

// UPDATE

type UpdateTaskInput struct {
	Input string `json:"input" binding:"required"`
}

func UpdateTaskType(context *gin.Context) {
	// Validate Input
	var input UpdateTaskInput
	err := context.ShouldBindJSON(&input)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"ERROR: ": err.Error()})
		return
	}

	// Update Task
	var task Models.Task
	err = Models.DB.Where("id = ?", context.Param("id")).First(&task).Error
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"ERROR": "Record not found!"})
		return
	}
	task.Type = input.Input
	Models.DB.Update(&task)
	context.JSON(http.StatusOK, gin.H{"data": task})
}

func UpdateTaskTitle(context *gin.Context) {
	// Validate Input
	var input UpdateTaskInput
	err := context.ShouldBindJSON(&input)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"ERROR: ": err.Error()})
		return
	}

	// Update Task
	var task Models.Task
	err = Models.DB.Where("id = ?", context.Param("id")).First(&task).Error
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"ERROR": "Record not found!"})
		return
	}
	task.Title = input.Input
	Models.DB.Update(&task)
	context.JSON(http.StatusOK, gin.H{"data": task})
}

func UpdateTaskContext(context *gin.Context) {
	// Validate Input
	var input UpdateTaskInput
	err := context.ShouldBindJSON(&input)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"ERROR: ": err.Error()})
		return
	}

	// Update Task
	var task Models.Task
	err = Models.DB.Where("id = ?", context.Param("id")).First(&task).Error
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"ERROR": "Record not found!"})
		return
	}
	task.Context = input.Input
	Models.DB.Update(&task)
	context.JSON(http.StatusOK, gin.H{"data": task})
}

// DELETE

func DeleteTask(c *gin.Context) {
	var task Models.Task
	if err := Models.DB.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	Models.DB.Delete(&task)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
