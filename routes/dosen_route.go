package routes

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"skripsi.id/obfuss/handlers"
	"skripsi.id/obfuss/middlewares"
)

type DosenRoute struct {
	handler *handlers.DosenHandler
	prefix  string
}

func NewDosenRoute(db *gorm.DB, logger *zap.Logger) Route {
	return &DosenRoute{
		handler: handlers.NewDosenHandler(db, logger),
		prefix:  "/dosen",
	}
}

func (r *DosenRoute) Register(app *gin.RouterGroup) {
	api := app.Group(r.prefix, middlewares.RequestId(), middlewares.Authorization())
	api.GET("/", r.handler.GetAll)
}
