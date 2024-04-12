package handlers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"skripsi.id/obfuss/models"
	"skripsi.id/obfuss/services"
)

type DosenHandler struct {
	db      *gorm.DB
	logger  *zap.Logger
	service *services.DosenService
}

func NewDosenHandler(db *gorm.DB, logger *zap.Logger) *DosenHandler {
	return &DosenHandler{
		db:      db,
		logger:  logger.With(zap.String("type", "DosenHandler")),
		service: services.NewDosenService(db, logger),
	}
}

// @Summary		Get Dosen
// @Description	Get Dosen
// @Tags			Dosen
// @Produce		json
// @Success		200			{object}	models.BaseResponse[[]models.GetAllDosen]
// @Router			/dosen [get]
// @Security BearerAuth
func (h *DosenHandler) GetAll(c *gin.Context) {
	dosen, err := h.service.GetAll()

	if err != nil {
		c.JSON(500, models.BaseResponse[map[string]any]{
			Message: "error",
			Errors:  []any{err.Error()},
		})
		return
	}

	c.JSON(200, models.BaseResponse[[]models.GetAllDosen]{
		Message: "success",
		Data:    dosen,
	})
}
