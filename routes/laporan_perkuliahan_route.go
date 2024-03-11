package routes

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"skripsi.id/obfuss/handlers"
	"skripsi.id/obfuss/middlewares"
)

type LaporanPerkuliahanRoute struct {
	handler *handlers.LaporanPerkuliahanHandler
	prefix  string
}

func NewLaporanPerkuliahanRoute(db *gorm.DB, logger *zap.Logger) Route {
	return &LaporanPerkuliahanRoute{
		handler: handlers.NewLaporanPerkuliahanHandler(db, logger),
		prefix:  "/laporan-perkuliahan",
	}
}

func (r *LaporanPerkuliahanRoute) Register(app *gin.RouterGroup) {
	api := app.Group(r.prefix, middlewares.RequestId(), middlewares.Authorization())
	api.GET("/", r.handler.GetAll)
	api.POST("/save-trans", r.handler.SaveTransLaporanPerkuliahan)
	api.DELETE("/:id-pertemuan/pertemuan", r.handler.Delete)
}
