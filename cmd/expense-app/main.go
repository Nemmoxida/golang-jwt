package main

import (
	"office-expense-management-backend/internals/router"
	"office-expense-management-backend/internals/services"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()


	r := router.Router()

	r.Run(":3001")
}
