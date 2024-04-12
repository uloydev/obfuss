package handlers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"skripsi.id/obfuss/models"
	"skripsi.id/obfuss/services"
)

type JamKuliahHandler struct {
	db      *gorm.DB
	logger  *zap.Logger
	service *services.JamKuliahService
}

func NewJamKuliahHandler(db *gorm.DB, logger *zap.Logger) *JamKuliahHandler {
	return &JamKuliahHandler{
		db:      db,
		logger:  logger.With(zap.String("type", "JamKuliahHandler")),
		service: services.NewJamKuliahService(db, logger),
	}
}

// @Summary		Get JamKuliah
// @Description	Get JamKuliah
// @Tags			JamKuliah
// @Produce		json
// @Success		200			{object}	models.BaseResponse[[]models.GetAllJamKuliah]
// @Router			/jam-kuliah [get]
// @Security BearerAuth
func (h *JamKuliahHandler) GetAll(c *gin.Context) {
	JamKuliah, err := h.service.GetAll()

	if err != nil {
		c.JSON(500, models.BaseResponse[map[string]any]{
			Message: "error",
			Errors:  []any{err.Error()},
		})
		return
	}

	c.JSON(200, models.BaseResponse[[]models.GetAllJamKuliah]{
		Message: "success",
		Data:    JamKuliah,
	})
}
