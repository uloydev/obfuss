package routes

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"skripsi.id/obfuss/handlers"
	"skripsi.id/obfuss/middlewares"
)

type TokenRoute struct {
	handler *handlers.TokenHandler
	prefix  string
}

func NewTokenRoute(db *gorm.DB, logger *zap.Logger) Route {
	return &TokenRoute{
		handler: handlers.NewTokenHandler(db, logger),
		prefix:  "/token",
	}
}

func (r *TokenRoute) Register(app *gin.RouterGroup) {
	api := app.Group(r.prefix, middlewares.RequestId())
	api.GET("/", r.handler.GetToken)
}
