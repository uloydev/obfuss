package routes

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"skripsi.id/obfuss/handlers"
	"skripsi.id/obfuss/middlewares"
)

type JadwalKuliahRoute struct {
	handler *handlers.JadwalKuliahHandler
	prefix  string
}

func NewJadwalKuliahRoute(db *gorm.DB, logger *zap.Logger) Route {
	return &JadwalKuliahRoute{
		handler: handlers.NewJadwalKuliahHandler(db, logger),
		prefix:  "/jadwal-kuliah",
	}
}

func (r *JadwalKuliahRoute) Register(app *gin.RouterGroup) {
	api := app.Group(r.prefix, middlewares.RequestId())
	api.GET("/", r.handler.GetJadwalKuliah)
	api.DELETE("/delete/:id", r.handler.DeleteJadwalKuliah)
	api.POST("/save-trans", r.handler.SaveTransJadwalKuliah)
}
