package controllers

import (
	"intikom-interview/dal"
	"intikom-interview/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// TaskInput represents the input for creating a new task.
type TaskInput struct {
	UserID      uint   `json:"user_id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

func CreateTask(c *gin.Context) {
	var taskInput TaskInput

	var task model.Task
	if err := c.ShouldBindJSON(&taskInput); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	u := dal.User
	_, err := u.Where(u.ID.Eq(uint(taskInput.UserID))).First()
	if err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	task = model.Task{
		UserID:      taskInput.UserID,
		Title:       taskInput.Title,
		Description: taskInput.Description,
		Status:      taskInput.Status,
	}

	if err := dal.Task.Create(&task); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, task)
}

func GetAllTasks(c *gin.Context) {
	resp, err := dal.Task.Find()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, resp)
}

func GetTaskByID(c *gin.Context) {
	taskID := c.Param("id")
	intID, err := strconv.Atoi(taskID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid parameter: " + taskID,
		})
		return
	}

	t := dal.Task
	task, err := t.Where(t.ID.Eq(uint(intID))).First()
	if err != nil {
		c.JSON(404, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(200, task)
}

func UpdateTask(c *gin.Context) {
	taskID := c.Param("id")

	intID, err := strconv.Atoi(taskID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid parameter: " + taskID,
		})
		return
	}

	t := dal.Task

	_, err = t.Where(t.ID.Eq(uint(intID))).First()
	if err != nil {
		c.JSON(404, gin.H{"error": "Task not found"})
		return
	}

	var taskInput TaskInput
	if err := c.ShouldBindJSON(&taskInput); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	u := dal.User
	_, err = u.Where(u.ID.Eq(uint(taskInput.UserID))).First()
	if err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	_, err = t.Where(t.ID.Eq(uint(intID))).Updates(model.Task{
		UserID:      taskInput.UserID,
		Title:       taskInput.Title,
		Description: taskInput.Description,
		Status:      taskInput.Status,
	})
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "successfully update task"})
}

func DeleteTask(c *gin.Context) {
	taskID := c.Param("id")
	intID, err := strconv.Atoi(taskID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid parameter: " + taskID,
		})
		return
	}

	t := dal.Task

	_, err = t.Where(t.ID.Eq(uint(intID))).Delete()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(204, nil)
}
