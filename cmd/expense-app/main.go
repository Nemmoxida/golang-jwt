package main

import (
	"office-expense-management-backend/internals/router"
	"office-expense-management-backend/internals/services"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	ExpenseList := services.NewExpenseList()
	UserUsage := services.NewUserUsage()

	r := router.Router(ExpenseList, UserUsage)

	r.Run(":3001")
}
