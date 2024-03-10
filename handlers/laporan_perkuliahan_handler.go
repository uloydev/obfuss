package handlers

import (
	"errors"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"skripsi.id/obfuss/models"
	"skripsi.id/obfuss/services"
	"skripsi.id/obfuss/utils"
)

type LaporanPerkuliahanHandler struct {
	db      *gorm.DB
	logger  *zap.Logger
	service *services.LaporanPerkuliahanService
}

func NewLaporanPerkuliahanHandler(db *gorm.DB, logger *zap.Logger) *LaporanPerkuliahanHandler {
	return &LaporanPerkuliahanHandler{
		db:      db,
		logger:  logger.With(zap.String("type", "JadwalKuliahHandler")),
		service: services.NewLaporanPerkuliahanService(db, logger),
	}
}

// @Summary		Get All Laporan Perkuliahan
// @Description	Get All Laporan Perkuliahan
// @Tags			Laporan Perkuliahan
// @Accept			json
// @Produce		json
// @Param			page	query	int	false	"Page"
// @Param			limit	query	int	false	"Limit"
// @Success		200		{object}	models.BaseResponse[any]
// @Router		/laporan-perkuliahan/ [get]
// @Security BearerAuth
func (h *LaporanPerkuliahanHandler) GetAll(c *gin.Context) {
	var params models.PaginationParams
	user, err := utils.GetUser(c)

	if err != nil {
		c.JSON(401, models.BaseResponse[any]{
			Message: "error",
			Errors:  []any{errors.New("unauthorize")},
		})
		return
	}

	if err := c.BindQuery(&params); err != nil {
		c.JSON(400, models.BaseResponse[any]{
			Message: "error",
			Errors:  []any{err.Error()},
		})
		return
	}

	data, meta, err := h.service.GetAllAngeketDosen(params, user)

	if err != nil {
		c.JSON(500, models.BaseResponse[any]{
			Message: "error",
			Errors:  []any{err.Error()},
		})
		return
	}

	c.JSON(200, models.BaseResponse[[]map[string]any]{
		Message: "success",
		Data:    data,
		Meta:    meta,
	})

}
