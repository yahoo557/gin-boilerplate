package user

import (
	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.RouterGroup) {

	userRouter := router.Group("/user")
	UserController := &userController{}
	{
		// User list
		userRouter.GET("/user", UserController.UserList)

		// Login router
		userRouter.POST("/login", UserController.Login)

		// Register router
		userRouter.POST("/register", UserController.Register)
	}

}
