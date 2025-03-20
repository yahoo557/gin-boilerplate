package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type (
	UserController interface {
		Login(c *gin.Context)
		Register(c *gin.Context)
	}

	userController struct {
	}
)

func (ctrl *userController) UserList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "UserList "})

}

func (ctrl *userController) Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})

}

func (ctrl userController) Register(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Register successful"})

}
