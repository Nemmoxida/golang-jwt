package services

import (
	"context"
	"net/http"
	"office-expense-management-backend/database"
	"office-expense-management-backend/pkg"
	"time"

	"github.com/gin-gonic/gin"
)

type UserUsageList struct {
	Id         int
	UserId     string
	Period     time.Time
	UsedBudget int
	Budget     int
	Username   string
	Category   string
}

type UserUsage struct{}

func NewUserUsage() *UserUsage {
	return &UserUsage{}
}

func (u *UserUsage) GetUserUsage(c *gin.Context) {
	claims := pkg.ExtrackClaims(c)

	userId := claims.UserId

	pool, err := database.Connect()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"DatabaseError": err})
		return
	}

	defer pool.Close()

	rows, err := pool.Query(context.Background(), "SELECT uu.id, uu.user_id, uu.period, uu.used_budget, uu.budget, u.username, c.name FROM user_usage uu JOIN users u ON u.id = uu.user_id JOIN category c ON c.id = uu.category_id WHERE u.id = $1 ORDER BY uu.used_budget DESC", userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"DatabaseErrorQuery": err})
		return
	}
	defer rows.Close()

	userUsageList := make([]UserUsageList, 0)

	for rows.Next() {
		var usage UserUsageList

		err := rows.Scan(
			&usage.Id,
			&usage.UserId,
			&usage.Period,
			&usage.UsedBudget,
			&usage.Budget,
			&usage.Username,
			&usage.Category,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		userUsageList = append(userUsageList, usage)
	}

	resp := &Respond{Created_at: time.Now(), Data: userUsageList}

	c.JSON(http.StatusOK, resp)
}
