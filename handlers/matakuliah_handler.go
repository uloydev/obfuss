package handlers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"skripsi.id/obfuss/models"
	"skripsi.id/obfuss/services"
)

type MatakuliahHandler struct {
	db      *gorm.DB
	logger  *zap.Logger
	service *services.MataKuliahService
}

func NewMatakuliahHandler(db *gorm.DB, logger *zap.Logger) *MatakuliahHandler {
	return &MatakuliahHandler{
		db:      db,
		logger:  logger.With(zap.String("type", "MatakuliahHandler")),
		service: services.NewMataKuliahService(db, logger),
	}
}

// @Summary		Get Matakuliah
// @Description	Get Matakuliah
// @Tags			Matakuliah
// @Produce		json
// @Success		200			{object}	models.BaseResponse[[]models.MataKuliah]
// @Router			/mata-kuliah [get]
// @Security BearerAuth
func (h *MatakuliahHandler) GetAll(c *gin.Context) {
	matakuliah, err := h.service.GetAll()

	if err != nil {
		c.JSON(500, models.BaseResponse[map[string]any]{
			Message: "error",
			Errors:  []any{err.Error()},
		})
		return
	}

	c.JSON(200, models.BaseResponse[[]models.MataKuliah]{
		Message: "success",
		Data:    matakuliah,
	})
}
