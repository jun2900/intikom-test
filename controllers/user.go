package controllers

import (
	"intikom-interview/dal"
	"intikom-interview/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// UserInput represents the input for creating a new user.
type UserInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

func RegisterUser(c *gin.Context) {
	var userInput UserInput

	var user model.User
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user = model.User{
		Name:     userInput.Name,
		Email:    userInput.Email,
		Password: userInput.Password,
	}

	if err := dal.User.Create(&user); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, user)
}

func GetAllUsers(c *gin.Context) {
	resp, err := dal.User.Find()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, resp)
}

// GetUserByID retrieves a user by ID.
func GetUserByID(c *gin.Context) {
	userID := c.Param("id")
	// Convert the parameter to an integer
	intID, err := strconv.Atoi(userID)
	if err != nil {
		// Handle the error if the conversion fails (e.g., invalid input)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid parameter: " + userID,
		})
		return
	}

	u := dal.User
	user, err := u.Where(u.ID.Eq(uint(intID))).First()
	if err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	c.JSON(200, user)
}

// UpdateUser updates an existing user by ID.
func UpdateUser(c *gin.Context) {
	userID := c.Param("id")

	intID, err := strconv.Atoi(userID)
	if err != nil {
		// Handle the error if the conversion fails (e.g., invalid input)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid parameter: " + userID,
		})
		return
	}

	u := dal.User

	user, err := u.Where(u.ID.Eq(uint(intID))).First()
	if err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	var userInput UserInput
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	_, err = u.Where(u.ID.Eq(uint(intID))).Updates(model.User{
		Name:     userInput.Name,
		Email:    userInput.Email,
		Password: userInput.Password,
	})
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, user)
}

// DeleteUser deletes a user by ID.
func DeleteUser(c *gin.Context) {
	userID := c.Param("id")
	intID, err := strconv.Atoi(userID)
	if err != nil {
		// Handle the error if the conversion fails (e.g., invalid input)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid parameter: " + userID,
		})
		return
	}

	u := dal.User

	_, err = u.Where(u.ID.Eq(uint(intID))).Delete()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(204, nil)
}
