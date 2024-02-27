package middlewares

import (
	"errors"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"skripsi.id/obfuss/models"
)

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

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			fmt.Println(os.Getenv("JWT_SECRET"))

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

		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok {
			c.JSON(401, models.BaseResponse[map[string]any]{
				Message: "error",
				Errors:  []any{errors.New("unauthorized").Error()},
			})
			c.Abort()

			return
		}

		actorId, ok := (claims["actorId"].(float64))
		if !ok {
			c.JSON(500, models.BaseResponse[map[string]any]{
				Message: "error",
				Errors:  []any{errors.New("internal server error").Error()},
			})
			c.Abort()

			return
		}

		userType, ok := claims["userType"].(string)
		if !ok {
			c.JSON(500, models.BaseResponse[map[string]any]{
				Message: "error",
				Errors:  []any{errors.New("internal server error").Error()},
			})
			c.Abort()

			fmt.Println("Semester id not found in claims")
			return
		}

		semesterId, ok := claims["semesterId"].(float64)
		if !ok {
			c.JSON(500, models.BaseResponse[map[string]any]{
				Message: "error",
				Errors:  []any{errors.New("internal server error").Error()},
			})
			c.Abort()

			fmt.Println("Semester id not found in claims")
			return
		}

		c.Set("x-actor-id", int(actorId))
		c.Set("x-user-type", userType)
		c.Set("x-semester-id", int(semesterId))
		c.Set("x-kelas-id", 0)
		c.Next()
	}
}
