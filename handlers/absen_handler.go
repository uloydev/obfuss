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

type AbsenHandler struct {
	db                     *gorm.DB
	logger                 *zap.Logger
	absenService           *services.AbsenService
	plotKelasService       *services.PlotKelasService
	jadwalPertemuanService *services.JadwalPertemuanService
	angketDosenService     *services.AngketDosenService
	jadwalKuliahService    *services.JadwalKuliahService
}

func NewAbsenHandler(db *gorm.DB, logger *zap.Logger) *AbsenHandler {
	return &AbsenHandler{
		db:                     db,
		logger:                 logger.With(zap.String("type", "AbsenHandler")),
		absenService:           services.NewAbsenService(db, logger),
		plotKelasService:       services.NewPlotKelasService(db, logger),
		jadwalPertemuanService: services.NewJadwalPertemuanService(db, logger),
		angketDosenService:     services.NewAngketDosenService(db, logger),
		jadwalKuliahService:    services.NewJadwalKuliahService(db, logger),
	}
}

// @Summary		Get Absen Mahasiswa
// @Description	Get Absen Mahasiswa
// @Tags			Absen
// @Accept			json
// @Produce		json
// @Param			params		query		models.PaginationParams	true	"Pagination parameters"
// @Param			userType	query		string					true	"User type"
// @Param			smtId		query		int						true	"Semester ID"
// @Param			kelasId		query		int						true	"Kelas ID"
// @Success		200			{object}	models.BaseResponse[models.GetAllAbsenResponse]
// @Router			/mahasiswa/absen [get]
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

// @Summary		Save Absen Trans
// @Description	Save Absen
// @Tags			Absen
// @Accept			json
// @Produce		json
// @Param			absen	body		models.SaveAbsenTransRequest	true	"Absen request"
// @Success		200		{object}	models.BaseResponse[any]
// @Router			/mahasiswa/absen/save-trans [post]
func (h *AbsenHandler) SaveTrans(c *gin.Context) {
	var req models.SaveAbsenTransRequest
	actorId := 1 // TODO: replace this with session('siap_userid') /  userId / mhsId
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, models.BaseResponse[map[string]any]{
			Message: "error",
			Errors:  []any{err.Error()},
		})
		return
	}

	jadwalKuliah, err := h.jadwalKuliahService.GetById(req.IdJadwal)
	if err != nil {
		c.JSON(500, models.BaseResponse[map[string]any]{
			Message: "error",
			Errors:  []any{err.Error()},
		})
		return
	}

	err = h.absenService.DeleteAbsenByPertemuan(req.IdPertemuan)
	if err != nil {
		c.JSON(500, models.BaseResponse[map[string]any]{
			Message: "error",
			Errors:  []any{err.Error()},
		})
		return
	}

	err = h.absenService.SaveAbsenTrans(req, actorId)
	if err != nil {
		c.JSON(500, models.BaseResponse[map[string]any]{
			Message: "error",
			Errors:  []any{err.Error()},
		})
		return
	}

	pertemuan, err := h.jadwalPertemuanService.GetById(req.IdPertemuan)
	if err != nil {
		c.JSON(500, models.BaseResponse[map[string]any]{
			Message: "error",
			Errors:  []any{err.Error()},
		})
		return
	}

	pertemuan.IsTerlaksana = "yes"
	err = h.jadwalPertemuanService.UpdateById(req.IdPertemuan, pertemuan)
	if err != nil {
		c.JSON(500, models.BaseResponse[map[string]any]{
			Message: "error",
			Errors:  []any{err.Error()},
		})
		return
	}

	angket, err := h.angketDosenService.FineOneByPertemuan(req.IdPertemuan)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {

		c.JSON(500, models.BaseResponse[map[string]any]{
			Message: "error",
			Errors:  []any{err.Error()},
		})
		return
	}

	absenCount, err := h.absenService.CountAbsen(req.IdPertemuan)
	if err != nil {
		c.JSON(500, models.BaseResponse[map[string]any]{
			Message: "error",
			Errors:  []any{err.Error()},
		})
		return
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		angket = entities.AngketDosen{
			IDPertemuan:      req.IdPertemuan,
			IDDosen:          *jadwalKuliah.IDDosen,
			JumlahHadir:      &absenCount.JumlahHadir,
			JumlahTidakHadir: &absenCount.JumlahTidakHadir,
			JumlahLuring:     &absenCount.JumlahLuring,
			JumlahDaring:     &absenCount.JumlahDaring,
			Status:           "Draft",
			AddUser:          &actorId,
		}
	} else {
		angket.JumlahHadir = &absenCount.JumlahHadir
		angket.JumlahTidakHadir = &absenCount.JumlahTidakHadir
		angket.JumlahLuring = &absenCount.JumlahLuring
		angket.JumlahDaring = &absenCount.JumlahDaring
		angket.Status = "Draft"
		angket.PertemuanKe = nil
		angket.JamKuliah = nil
	}

	if err := h.angketDosenService.SaveAngketDosen(&angket); err != nil {
		c.JSON(500, models.BaseResponse[map[string]any]{
			Message: "error",
			Errors:  []any{err.Error()},
		})
		return
	}

	c.JSON(200, models.BaseResponse[map[string]any]{
		Message: "success",
	})
}
