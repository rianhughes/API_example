package handlers

import (
	"net/http"
	"nethgateapi/nethgateapi"

	"github.com/gin-gonic/gin"
)

// Assuming you have a models.User type and service.NethGateServiceInterface defined.

// UserHandler holds the methods for handling user-related requests
type UserHandler struct {
	Service nethgateapi.NethGateServiceInterface
}

// NewUserHandler creates a new UserHandler with the given service.
func NewUserHandler(svc nethgateapi.NethGateServiceInterface) *UserHandler {
	return &UserHandler{Service: svc}
}

// CreateUser is a gin handler function to create a new user.
func (h *UserHandler) CreateUser(c *gin.Context) {
	var newUser nethgateapi.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdUser, err := h.Service.CreateUser(newUser.Login, "todo")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, createdUser)
}

// ... Define other handler functions for SearchUser, UpdateUser, DeleteUser, etc.
