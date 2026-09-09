package router

import (
	"office-expense-management-backend/internals/services"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.POST("/login", services.Login)
	r.POST("/signup", services.Signup)

	return r

}
