package routes

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"skripsi.id/obfuss/handlers"
	"skripsi.id/obfuss/middlewares"
)

type JadwalKuliahDosenRoute struct {
	handler *handlers.JadwalKuliahDosenHandler
	prefix  string
}

func NewJadwalKuliahDosenRoute(db *gorm.DB, logger *zap.Logger) Route {
	return &JadwalKuliahDosenRoute{
		handler: handlers.NewJadwalKuliahDosenHandler(db, logger),
		prefix:  "/jadwal-kuliah-dosen",
	}
}

func (r *JadwalKuliahDosenRoute) Register(app *gin.RouterGroup) {
	api := app.Group(r.prefix, middlewares.RequestId(), middlewares.Authorization())
	api.GET("/", r.handler.GetAll)
	api.DELETE("/delete/:id", r.handler.Delete)
	api.POST("/save-trans", r.handler.SaveTrans)
	api.POST("/mahasiswa", r.handler.GetAllMahasiswa)
}
