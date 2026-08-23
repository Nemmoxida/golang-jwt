package services

import (
	"context"
	"fmt"
	"net/http"
	"office-expense-management-backend/database"
	"office-expense-management-backend/pkg"
	"time"

	"github.com/gin-gonic/gin"
)

type ExpenseQuery struct {
	Claim_id          string    `json:"id"`
	Username          string    `json:"username"`
	Claim_title       string    `json:"claim_title"`
	Claim_description string    `json:"claim_description"`
	Merchant          string    `json:"merchant"`
	Category          string    `json:"category"`
	Image_id          string    `json:"image_id"`
	Creation_date     time.Time `json:"creation_date"`
	Amount            int       `json:"amount"`
	Status            string    `json:"status"`
	Departement       string    `json:"departement"`
}

type ExpenseList struct {
}

func NewExpenseList() *ExpenseList {
	return &ExpenseList{}
}

func (e *ExpenseList) GetExpenseList(c *gin.Context) {
	token := c.GetHeader("authorization")

	verifyToken, err := pkg.VerifyToken(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	userDepartement := verifyToken.Departement

	pool, err := database.Connect()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"DatabaseError": err})
		return
	}

	defer pool.Close()

	fmt.Println(userDepartement)

	rows, err := pool.Query(context.Background(), "SELECT ec.id, ec.claim_title, ec.claim_description, ec.category, ec.image_id, ec.creation_date, ec.amount, ec.status, u.departement, u.username, m.name FROM expenses_claim ec JOIN users u ON u.id = ec.user_id JOIN merchants m ON m.id = ec.merchant_id WHERE u.departement = $1 LIMIT 100", userDepartement)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"DatabaseErrorQuery": err})
		return
	}
	defer rows.Close()

	expenses := make([]ExpenseQuery, 0)

	for rows.Next() {
		var expense ExpenseQuery

		err := rows.Scan(
			&expense.Claim_id,
			&expense.Claim_title,
			&expense.Claim_description,
			&expense.Category,
			&expense.Image_id,
			&expense.Creation_date,
			&expense.Amount,
			&expense.Status,
			&expense.Departement,
			&expense.Username,
			&expense.Merchant,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		expenses = append(expenses, expense)
	}

	c.JSON(http.StatusOK, expenses)
}
