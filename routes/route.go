package routes

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Route interface {
	Register(app *gin.RouterGroup)
}

type RouteInitFunc func(*gorm.DB, *zap.Logger) Route

// Routes is a list of Route
var Routes = []RouteInitFunc{
	NewTodoRoute,
	NewAbsenRoute,
	NewJadwalKuliahRoute,
	NewTokenRoute,
	NewPerubahanJadwalRoute,
	NewJadwalKuliahDosenRoute,
	NewLaporanPerkuliahanRoute,
}

func Setup(app *gin.RouterGroup, db *gorm.DB, logger *zap.Logger) {
	for _, route := range Routes {
		route(db, logger).Register(app)
	}
}
