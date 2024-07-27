package handlers

import (
	"errors"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"skripsi.id/obfuss/models"
	"skripsi.id/obfuss/services"
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
// @Success		200		{object}	models.BaseResponse[[]models.GetAllPerubahanJadwal]
// @Router			/perubahan-jadwal [get]
// @Security BearerAuth
func (h *PerubahanJadwalHandler) GetPerubahanJadwal(c *gin.Context) {
	semesterId, err := strconv.Atoi(c.Query("smtId"))

	if err != nil {
		c.JSON(400, models.BaseResponse[any]{
			Message: "error",
			Errors:  []any{errors.New("invalid mahasiswa id").Error()},
		})
		return
	}

	perubahanJadwal, err := h.service.GetPerubahanJadwal(uint(semesterId))
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
	})
}

// @Summary		Update Perubahan Jadwal
// @Description	Update Perubahan Jadwal
// @Tags			Perubahan Jadwal
// @Accept			json
// @Produce		json
// @Param			idJadwalPertemuan	path	int	true	"ID Jadwal Pertemuan"
// @Param			payload		body	models.UpdateJadwalPertemuanRequest	true	"Update Jadwal Pertemuan Request"
// @Success		204		{object}	models.BaseResponse[any]
// @Router			/perubahan-jadwal/{idJadwalPertemuan} [patch]
// @Security BearerAuth
func (h *PerubahanJadwalHandler) Update(c *gin.Context) {
	var payload models.UpdateJadwalPertemuanRequest

	idPertemuanStr := c.Param("idJadwalPertemuan")
	idPertemuan, err := strconv.Atoi(idPertemuanStr)

	if err != nil {
		c.JSON(400, models.BaseResponse[any]{
			Message: "error",
			Errors:  []any{errors.New("invalid id pertemuan")},
		})
		return
	}

	mahasiswaId, err := strconv.Atoi(c.Query("mahasiswaId"))
	if err != nil {
		c.JSON(400, models.BaseResponse[any]{
			Message: "error",
			Errors:  []any{errors.New("invalid mahasiswa id")},
		})
		return
	}

	if err := c.BindJSON(&payload); err != nil {
		c.JSON(400, models.BaseResponse[any]{
			Message: "error",
			Errors:  []any{err.Error()},
		})
		return
	}

	tanggalUsulanGanti, err := time.Parse("2006-01-02 15:04:05", payload.TanggalUsulanGanti)
	if err != nil {
		c.JSON(400, models.BaseResponse[any]{
			Message: "error",
			Errors:  []any{errors.New("invalid date format")},
		})
		return
	}

	jadwalPertemuan, err := h.service.GetJadwalById(idPertemuan)
	if err != nil {
		c.JSON(500, models.BaseResponse[any]{
			Message: "error",
			Errors:  []any{err.Error()},
		})
		return
	}

	jadwalPertemuan.TanggalUsulanGanti = tanggalUsulanGanti
	jadwalPertemuan.UsulanMulaiJam = payload.UsulanMulaiJam
	jadwalPertemuan.UsulanSampaiJam = payload.UsulanSampaiJam
	jadwalPertemuan.StatusUsulan = payload.StatusUsulan
	jadwalPertemuan.IDJadwal = payload.IDJadwalKuliah
	jadwalPertemuan.AlasanPerubahan = payload.AlasanPerubahan
	jadwalPertemuan.ModifiedDate = time.Now()
	jadwalPertemuan.ModifiedUser = mahasiswaId

	err = h.service.SavePerubahanJadwal(&jadwalPertemuan)
	if err != nil {
		c.JSON(500, models.BaseResponse[any]{
			Message: "error",
			Errors:  []any{err.Error()},
		})
		return
	}

	c.JSON(200, models.BaseResponse[any]{
		Data:    jadwalPertemuan,
		Message: "success",
	})
}
