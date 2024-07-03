package routes

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"skripsi.id/obfuss/handlers"
	"skripsi.id/obfuss/middlewares"
)

type AbsenRoute struct {
	handler *handlers.AbsenHandler
	prefix  string
}

func NewAbsenRoute(db *gorm.DB, logger *zap.Logger) Route {
	return &AbsenRoute{
		handler: handlers.NewAbsenHandler(db, logger),
		prefix:  "/mahasiswa/absen",
	}
}

func (r *AbsenRoute) Register(app *gin.RouterGroup) {
	api := app.Group(r.prefix, middlewares.RequestId(), middlewares.Authorization())
	api.GET("/", r.handler.GetAbsenMhs)
	api.POST("/save-trans", r.handler.SaveTrans)
	api.POST("/:idPertemuan/delete", r.handler.Delete)
	api.GET("/pertemuan", r.handler.ListAbsen)
}
