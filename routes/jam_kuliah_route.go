package routes

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"skripsi.id/obfuss/handlers"
	"skripsi.id/obfuss/middlewares"
)

type JamKuliahRoute struct {
	handler *handlers.JamKuliahHandler
	prefix  string
}

func NewJamKuliahRoute(db *gorm.DB, logger *zap.Logger) Route {
	return &JamKuliahRoute{
		handler: handlers.NewJamKuliahHandler(db, logger),
		prefix:  "/jam-kuliah",
	}
}

func (r *JamKuliahRoute) Register(app *gin.RouterGroup) {
	api := app.Group(r.prefix, middlewares.RequestId(), middlewares.Authorization())
	api.GET("/", r.handler.GetAll)
}
