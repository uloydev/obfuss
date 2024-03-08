package middlewares

import (
	"errors"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"skripsi.id/obfuss/models"
)

type User struct {
	ActorID    int    `json:"actorId"`
	UserType   string `json:"userType"`
	SemesterId int    `json:"semesterId"`
	KelasId    int    `json:"kelasId"`
}

type UserToken struct {
	jwt.MapClaims
	*User
}

func Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")

		if tokenString == "" {
			c.JSON(401, models.BaseResponse[map[string]any]{
				Message: "error",
				Errors:  []any{errors.New("Unauthorized").Error()},
			})
			c.Abort()

			return
		}

		token, err := jwt.ParseWithClaims(tokenString, &UserToken{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil {
			fmt.Print(err)
			c.JSON(500, models.BaseResponse[map[string]any]{
				Message: "error",
				Errors:  []any{errors.New("internal Server Error").Error()},
			})
			c.Abort()

			return
		}

		if !token.Valid {
			c.JSON(401, models.BaseResponse[map[string]any]{
				Message: "error",
				Errors:  []any{errors.New("unauthorized").Error()},
			})
			c.Abort()

			return
		}

		claims, ok := token.Claims.(*UserToken)
		claims.SemesterId = 1

		if !ok {
			c.JSON(401, models.BaseResponse[map[string]any]{
				Message: "error",
				Errors:  []any{errors.New("unauthorized").Error()},
			})
			c.Abort()

			return
		}

		c.Set("user", claims.User)
		c.Next()
	}
}
