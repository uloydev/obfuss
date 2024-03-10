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
	"skripsi.id/obfuss/utils"
)

type JadwalKuliahDosenHandler struct {
	db         *gorm.DB
	logger     *zap.Logger
	service    *services.JadwalKuliahDosenService
	mhsService *services.MahasiswaService
}

func NewJadwalKuliahDosenHandler(db *gorm.DB, logger *zap.Logger) *JadwalKuliahDosenHandler {
	return &JadwalKuliahDosenHandler{
		db:         db,
		logger:     logger.With(zap.String("type", "JadwalKuliahHandler")),
		service:    services.NewJadwalKuliahDosenService(db, logger),
		mhsService: services.NewMahasiswaService(db, logger),
	}
}

// @Summary		Get All Jadwal Kuliah Dosen
// @Description	Get All Jadwal Kuliah Dosen
// @Tags			Jadwal Kuliah
// @Accept			json
// @Produce		json
// @Param			page	query	int	false	"Page"
// @Param			limit	query	int	false	"Limit"
// @Success		200		{object}	models.BaseResponse[any]
// @Router		/jadwal-kuliah-dosen/ [get]
// @Security BearerAuth
func (h *JadwalKuliahDosenHandler) GetAll(c *gin.Context) {
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

	data, meta, err := h.service.GetAllJadwalKuliahDosen(params, user.SemesterId)

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

	user, err := utils.GetUser(c)

	if err != nil {
		c.JSON(401, models.BaseResponse[entities.JadwalKuliahDosen]{
			Message: "error",
			Errors:  []any{errors.New("unauthorize")},
		})
		return
	}

	if err := h.service.SaveTrans(&body, user.ActorID); err != nil {
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
// @Success		204		{object}	models.BaseResponse[any]
// @Router			/jadwal-kuliah-dosen/{id} [delete]
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

// @Summary		Get All Mahasiswa
// @Description	Get All Mahasiswa
// @Tags			Jadwal Kuliah
// @Accept			json
// @Produce		json
// @Param			data	body	models.GetMahasiswaRaw  true  "data"
// @Param			page	query	int	false	"Page"
// @Param			limit	query	int	false	"Limit"
// @Success		200		{object}	models.BaseResponse[any]
// @Router		/jadwal-kuliah-dosen/mahasiswa [post]
// @Security BearerAuth
func (h *JadwalKuliahDosenHandler) GetAllMahasiswa(c *gin.Context) {
	var params models.PaginationParams
	var body models.GetMahasiswaRaw

	if err := c.BindQuery(&params); err != nil {
		c.JSON(400, models.BaseResponse[any]{
			Message: "error",
			Errors:  []any{err.Error()},
		})
		return
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(400, models.BaseResponse[any]{
			Message: "error",
			Errors:  []any{err.Error()},
		})
		return
	}

	data, meta, err := h.mhsService.GetMahasiswaRaw(params, body.IDKelas, body.IDPertemuan)

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
