package services

import "github.com/gin-gonic/gin"

type UserUsage struct{}

func NewUserUsage() *UserUsage {
	return &UserUsage{}
}

func (u *UserUsage) GetUserUsage(c *gin.Context) {

}
