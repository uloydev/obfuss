package handlers

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"skripsi.id/obfuss/entities"
	"skripsi.id/obfuss/models"
	"skripsi.id/obfuss/services"
)

type JadwalKuliahDosenHandler struct {
	db      *gorm.DB
	logger  *zap.Logger
	service *services.JadwalKuliahDosenService
}

func NewJadwalKuliahDosenHandler(db *gorm.DB, logger *zap.Logger) *JadwalKuliahDosenHandler {
	return &JadwalKuliahDosenHandler{
		db:      db,
		logger:  logger.With(zap.String("type", "JadwalKuliahHandler")),
		service: services.NewJadwalKuliahDosenService(db, logger),
	}
}

// @Summary		Save Trans Jadwal Kuliah
// @Description	Save Trans Jadwal Kuliah
// @Tags			Jadwal Kuliah
// @Accept			json
// @Produce		json
// @Success		204		{object}	models.BaseResponse[any]
// @Router		/jadwal-kuliah-dosen/save-trans [post]
// @Security BearerAuth
func (h *JadwalKuliahDosenHandler) SaveTrans(c *gin.Context) {
	var body models.SaveTransJadwalKuliahDosen

if err := c.BindJSON(&body); err != nil {
		c.JSON(400, models.BaseResponse[entities.JadwalKuliahDosen]{
			Message: "error",
			Errors:  []any{err.Error()},
		})
		return
	}

	if err := h.service.SaveTrans(&body); err != nil {
		c.JSON(500, models.BaseResponse[entities.JadwalKuliahDosen]{
			Message: "error",
			Errors:  []any{err.Error()},
		})

		return
	}

	c.JSON(204, models.BaseResponse[*entities.JadwalKuliahDosen]{
		Message: "success",
		Data:    nil,
	})
}

// @Summary		delete Trans Jadwal Kuliah
// @Description	delete Trans Jadwal Kuliah
// @Tags			Jadwal Kuliah
// @Accept			json
// @Produce		json
// @Param			id	path		int	true	"Jadwal Kuliah ID"
// @Success		200		{object}	models.BaseResponse[any]
// @Router			/jadwal-kuliah-dosen/ [delete]
// @Security BearerAuth

func (h *JadwalKuliahDosenHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(400, models.BaseResponse[any]{
			Message: "error",
			Errors:  []any{errors.New("bad request").Error()},
		})

		return
	}

	err = h.service.Delete(id)

	if err != nil {
		c.JSON(500, models.BaseResponse[any]{
			Message: "error",
			Errors:  []any{"error"},
		})

		return
	}

	c.JSON(204, models.BaseResponse[any]{
		Message: "success  delete jadwal kuliah",
		Data:    nil,
	})
}
