package handlers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"skripsi.id/obfuss/services"
)

type JadwalKuliahHandler struct {
	db      *gorm.DB
	logger  *zap.Logger
	service *services.TodoService
}

func NewJadwalKuliahHandler(db *gorm.DB, logger *zap.Logger) *JadwalKuliahHandler {
	return &JadwalKuliahHandler{
		db:      db,
		logger:  logger.With(zap.String("type", "JadwalKuliahHandler")),
		service: services.NewTodoService(db, logger),
	}
}

// @Summary		Get Todos
// @Description	Get Todos
// @Tags			Todo
// @Accept			json
// @Produce		json
// @Param			params	query	models.PaginationParams	true	"Pagination parameters"
// @Success		200	{object}	models.BaseResponse[[]entities.Todo]
// @Router			/todo [get]
func (j *JadwalKuliahHandler) GetJadwalKuliah(c *gin.Context) {

}
