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
// @Param			smtId		query		int						true	"Semester ID"
// @Param			kelasId		query		int						true	"Kelas ID"
// @Success		200			{object}	models.BaseResponse[models.GetAllAbsenResponse]
// @Router			/mahasiswa/absen [get]
// @Security BearerAuth
func (h *AbsenHandler) GetAbsenMhs(c *gin.Context) {
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
		c.JSON(400, models.BaseResponse[map[string]any]{
			Message: "error",
			Errors:  []any{err.Error()},
		})
		return
	}

	smtId, err := strconv.Atoi(c.Query("smtId"))
	if err != nil {
		c.JSON(400, models.BaseResponse[map[string]any]{
			Message: "error",
			Errors:  []any{"smtId must be a number"},
		})
		return
	}

	kelasId, err := h.plotKelasService.GetIdKelas(map[string]any{
		"id_semester":  smtId,
		"id_mahasiswa": user.ActorID,
	})
	if err != nil {
		c.JSON(500, models.BaseResponse[map[string]any]{
			Message: "error",
			Errors:  []any{err.Error()},
		})
		return
	}

	absen, meta, err := h.absenService.GetAllAbsen(params, user.UserType, smtId, kelasId, 0)
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
// @Security BearerAuth
func (h *AbsenHandler) SaveTrans(c *gin.Context) {
	var req models.SaveAbsenTransRequest
	user, err := utils.GetUser(c)
	if err != nil {
		c.JSON(401, models.BaseResponse[any]{
			Message: "error",
			Errors:  []any{errors.New("unauthorize")},
		})
		return
	}
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

	err = h.absenService.SaveAbsenTrans(req, user.ActorID)
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
			AddUser:          &user.ActorID,
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

// @Summary		Delete Absen
// @Description	Delete Absen
// @Tags			Absen
// @Accept			json
// @Produce		json
// @Param			idPertemuan	path	int	true	"Pertemuan ID"
// @Success		204		{object}	models.BaseResponse[any]
// @Router			/mahasiswa/absen/{idPertemuan} [delete]
// @Security BearerAuth
func (h *AbsenHandler) Delete(c *gin.Context) {
	idPertemuanStr := c.Param("idPertemuan")
	idPertemuan, err := strconv.Atoi(idPertemuanStr)
	if err != nil {
		c.JSON(400, models.BaseResponse[any]{
			Message: "error",
			Errors:  []any{errors.New("idPertemuan must be a number")},
		})
		return
	}
	if err := h.absenService.DeleteByPertemuanID(idPertemuan); err != nil {
		c.JSON(500, models.BaseResponse[any]{
			Message: "internal error",
			Errors:  []any{err.Error()},
		})
	}

	if err := h.angketDosenService.DeleteByPertemuanID(idPertemuan); err != nil {
		c.JSON(500, models.BaseResponse[any]{
			Message: "internal error",
			Errors:  []any{err.Error()},
		})
	}

	c.JSON(204, models.BaseResponse[any]{
		Data:    nil,
		Message: "success",
	})

}

func (h *AbsenHandler) ListAbsen(c *gin.Context) {
	idKelas, err := strconv.Atoi(c.Query("idKelas"))
	if err != nil {
		c.JSON(400, models.BaseResponse[map[string]any]{
			Message: "error",
			Errors:  []any{"smtId must be a number"},
		})
		return
	}

	idPertemuan, err := strconv.Atoi(c.Query("idPertemuan"))
	if err != nil {
		c.JSON(400, models.BaseResponse[map[string]any]{
			Message: "error",
			Errors:  []any{"smtId must be a number"},
		})
		return
	}

	var mst []models.Result

	h.db.Table("mst_mahasiswa").
		Select("DISTINCT trans_absen_mahasiswa.is_hadir AS kehadiran, trans_absen_mahasiswa.keterangan, mst_mahasiswa.user_name, mst_mahasiswa.nama_lengkap").
		Joins("JOIN trans_plot_kelas ON trans_plot_kelas.id_mahasiswa = mst_mahasiswa.id AND trans_plot_kelas.id_kelas = ?", idKelas).
		Joins("LEFT JOIN trans_absen_mahasiswa ON trans_absen_mahasiswa.id_mahasiswa = mst_mahasiswa.id AND trans_absen_mahasiswa.id_pertemuan = ?", idPertemuan).
		Find(&mst)

	c.JSON(200, models.BaseResponse[[]models.Result]{
		Data:    mst,
		Message: "success",
	})
}
