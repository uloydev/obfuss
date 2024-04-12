package routes

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"skripsi.id/obfuss/handlers"
	"skripsi.id/obfuss/middlewares"
)

type MatakuliahRoute struct {
	handler *handlers.MatakuliahHandler
	prefix  string
}

func NewMatakuliahRoute(db *gorm.DB, logger *zap.Logger) Route {
	return &MatakuliahRoute{
		handler: handlers.NewMatakuliahHandler(db, logger),
		prefix:  "/mata-kuliah",
	}
}

func (r *MatakuliahRoute) Register(app *gin.RouterGroup) {
	api := app.Group(r.prefix, middlewares.RequestId(), middlewares.Authorization())
	api.GET("/", r.handler.GetAll)
}
