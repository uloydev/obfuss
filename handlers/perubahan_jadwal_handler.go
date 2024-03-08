package handlers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"skripsi.id/obfuss/entities"
	"skripsi.id/obfuss/models"
	"skripsi.id/obfuss/services"
	"skripsi.id/obfuss/utils"
)

type PerubahanJadwalHandler struct {
	db      *gorm.DB
	logger  *zap.Logger
	service *services.PerubahanJadwalService
}

func NewPerubahanJadwalHandler(db *gorm.DB, logger *zap.Logger) *PerubahanJadwalHandler {
	return &PerubahanJadwalHandler{
		db:      db,
		logger:  logger.With(zap.String("type", "PerubahanJadwalHandler")),
		service: services.NewPerubahanJadwalService(db, logger),
	}
}

// @Summary		Get Perubahan Jadwal
// @Description	Get Perubahan Jadwal
// @Tags			Perubahan Jadwal
// @Accept			json
// @Produce		json
// @Param			params	query		models.PaginationParams	true	"Pagination parameters"
// @Success		200		{object}	models.BaseResponse[[]models.GetAllPerubahanJadwal]
// @Router			/perubahan-jadwal [get]
// @Security BearerAuth
func (h *PerubahanJadwalHandler) GetPerubahanJadwal(c *gin.Context) {
	var params models.PaginationParams

	user, error := utils.GetUser(c)

	if error != nil {
		c.JSON(500, models.BaseResponse[entities.Todo]{
			Message: "error",
			Errors:  []any{error.Error()},
		})
		return
	}

	perubahanJadwal, meta, err := h.service.GetPerubahanJadwal(params, uint(user.SemesterId))

	if err != nil {
		c.JSON(500, models.BaseResponse[any]{
			Message: "error",
			Errors:  []any{err.Error()},
		})
		return
	}

	c.JSON(200, models.BaseResponse[[]models.GetAllPerubahanJadwal]{
		Message: "success",
		Data:    perubahanJadwal,
		Meta:    meta,
	})
}
