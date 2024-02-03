package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"skripsi.id/obfuss/models"
	"skripsi.id/obfuss/services"
)

type AbsenHandler struct {
	db               *gorm.DB
	logger           *zap.Logger
	absenService     *services.AbsenService
	plotKelasService *services.PlotKelasService
}

func NewAbsenHandler(db *gorm.DB, logger *zap.Logger) *AbsenHandler {
	return &AbsenHandler{
		db:               db,
		logger:           logger.With(zap.String("type", "AbsenHandler")),
		absenService:     services.NewAbsenService(db, logger),
		plotKelasService: services.NewPlotKelasService(db, logger),
	}
}

// @Summary		Get Absen Mahasiswa
// @Description	Get Absen Mahasiswa
// @Tags			Absen
// @Accept			json
// @Produce		json
// @Param			params	query	models.PaginationParams	true	"Pagination parameters"
// @Param			userType	query	string	true	"User type"
// @Param			smtId	query	int	true	"Semester ID"
// @Param			kelasId	query	int	true	"Kelas ID"
// @Success		200	{object}	models.BaseResponse[models.GetAllAbsenResponse]
// @Router			/absen/mahasiswa [get]
func (h *AbsenHandler) GetAbsenMhs(c *gin.Context) {
	var params models.PaginationParams
	if err := c.BindQuery(&params); err != nil {
		c.JSON(400, models.BaseResponse[map[string]any]{
			Message: "error",
			Errors:  []any{err.Error()},
		})
		return
	}

	// START TODO: remove this params when user auth is implemented
	userType := c.Query("userType")
	smtId, err := strconv.Atoi(c.Query("smtId"))
	if err != nil {
		c.JSON(400, models.BaseResponse[map[string]any]{
			Message: "error",
			Errors:  []any{"smtId must be a number"},
		})
		return
	}
	// END TODO

	kelasId, err := h.plotKelasService.GetIdKelas(map[string]any{
		"id_semester":  smtId,
		"id_mahasiswa": 1, // TODO: replace this with session('siap_id') /  userId / mhsId
	})
	if err != nil {
		c.JSON(500, models.BaseResponse[map[string]any]{
			Message: "error",
			Errors:  []any{err.Error()},
		})
		return
	}

	absen, meta, err := h.absenService.GetAllAbsen(params, userType, smtId, kelasId, 0)
	if err != nil {
		c.JSON(500, models.BaseResponse[map[string]any]{
			Message: "error",
			Errors:  []any{err.Error()},
		})
		return
	}

	c.JSON(200, models.BaseResponse[[]models.GetAllAbsenResponse]{
		Message: "success",
		Data:    absen,
		Meta:    meta,
	})
}
