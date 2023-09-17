package handlers

import (
	"fyno/server/internal/models"
	"fyno/server/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandlers struct {
	userService models.UserService
}

func NewUserHandlers(us models.UserService) models.UserHandlers {
	return &userHandlers{
		userService: us,
	}
}

func (uh *userHandlers) GetUser(c *gin.Context) {
	id := c.Param("id")

	user, err := uh.userService.GetUser(utils.StringToUUID(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func (uh *userHandlers) GetUserByName(c *gin.Context) {
	name := c.Param("name")

	user, err := uh.userService.GetUserByName(name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func (uh *userHandlers) CreateUser(c *gin.Context) {
	var input *models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := uh.userService.CreateUser(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}
