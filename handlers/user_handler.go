package handlers

import (
	"movies-app/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserRepo *repositories.UserRepository
}

type CreateUserRequest struct {
	Name     string `json:"name" binding:"max=20"`
	Password string `json:"password" binding:"required,min=6"`
	Email    string `json:"email" binding:"required,email"`
}

func (handler *UserHandler) Register(c *gin.Context) {
	var req CreateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := handler.UserRepo.Create(req.Email, req.Password, req.Name); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}
