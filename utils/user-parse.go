package utils

import (
	"errors"

	"github.com/gin-gonic/gin"
	"skripsi.id/obfuss/middlewares"
)

func GetUser(c *gin.Context) (*middlewares.User, error) {
	userCtx, exist := c.Get("user")

	if !exist {
		return nil, errors.New("user data not found")
	}

	user, ok := userCtx.(*middlewares.User)

	if !ok {
		return nil, errors.New("error when parse user")
	}

	return user, nil
}
