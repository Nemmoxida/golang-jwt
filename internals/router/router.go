package router

import (
	"office-expense-management-backend/internals/services"
	"office-expense-management-backend/middleware"

	"github.com/gin-gonic/gin"
)

func Router(expenseList *services.ExpenseList, userUsage *services.UserUsage) *gin.Engine {
	r := gin.Default()

	r.GET("/expense", middleware.AuthHandler(), expenseList.GetExpenseList)
	r.GET("/userusage", middleware.AuthHandler(), userUsage.GetUserUsage)
	r.POST("/login", services.Login)

	return r

}
