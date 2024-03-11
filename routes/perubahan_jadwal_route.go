package routes

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"skripsi.id/obfuss/handlers"
	"skripsi.id/obfuss/middlewares"
)

type PerubahanJadwalRoute struct {
	handler *handlers.PerubahanJadwalHandler
	prefix  string
}

func NewPerubahanJadwalRoute(db *gorm.DB, logger *zap.Logger) Route {
	return &PerubahanJadwalRoute{
		handler: handlers.NewPerubahanJadwalHandler(db, logger),
		prefix:  "/perubahan-jadwal",
	}
}

func (r *PerubahanJadwalRoute) Register(app *gin.RouterGroup) {
	api := app.Group(r.prefix, middlewares.RequestId(), middlewares.Authorization())
	api.GET("/", r.handler.GetPerubahanJadwal)
	api.PATCH("/:idJadwalPertemuan", r.handler.Update)
}
