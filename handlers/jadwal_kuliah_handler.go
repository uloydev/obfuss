package handlers

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"skripsi.id/obfuss/entities"
	"skripsi.id/obfuss/middlewares"
	"skripsi.id/obfuss/models"
	"skripsi.id/obfuss/services"
)

type JadwalKuliahHandler struct {
	db      *gorm.DB
	logger  *zap.Logger
	service *services.JadwalKuliahService
}

func NewJadwalKuliahHandler(db *gorm.DB, logger *zap.Logger) *JadwalKuliahHandler {
	return &JadwalKuliahHandler{
		db:      db,
		logger:  logger.With(zap.String("type", "JadwalKuliahHandler")),
		service: services.NewJadwalKuliahService(db, logger),
	}
}

// @Summary		Get Jadwal Kuliah
// @Description	Get jadwal Kuliah
// @Tags			Jadwal Kuliah
// @Accept			json
// @Produce		json
// @Param			params	query		models.PaginationParams	true	"Pagination parameters"
// @Success		200		{object}	models.BaseResponse[[]entities.JadwalKuliah]
// @Router			/jadwal-kuliah [get]
// @Security BearerAuth
func (h *JadwalKuliahHandler) GetJadwalKuliah(c *gin.Context) {
	var params models.PaginationParams

	userContext, exist := c.Get("user")

	if !exist {
		c.JSON(401, models.BaseResponse[entities.Todo]{
			Message: "error",
			Errors:  []any{errors.New("unauthorize")},
		})
		return
	}

	user, ok := userContext.(*middlewares.User)

	if !ok {
		h.logger.Error("error when parse user data")

		c.JSON(500, models.BaseResponse[entities.Todo]{
			Message: "error",
			Errors:  []any{errors.New("internal server error")},
		})

		return
	}

	if err := c.BindQuery(&params); err != nil {
		c.JSON(400, models.BaseResponse[entities.Todo]{
			Message: "error",
			Errors:  []any{err.Error()},
		})
		return
	}

	fmt.Println(user)

	jadwalKuliah, meta, err := h.service.GetJadwalKuliah(params, uint(user.SemesterId))
	if err != nil {
		c.JSON(500, models.BaseResponse[entities.JadwalKuliah]{
			Message: "error",
			Errors:  []any{err.Error()},
		})
		return
	}

	c.JSON(200, models.BaseResponse[[]entities.JadwalKuliah]{
		Message: "success",
		Data:    jadwalKuliah,
		Meta:    meta,
	})
}

// @Summary		Save Trans Jadwal Kuliah
// @Description	Save Trans Jadwal Kuliah
// @Tags			Jadwal Kuliah
// @Accept			json
// @Produce		json
// @Param			jadwalKuliah	body		models.JadwalKuliahRequest	true "Save Trans request"
// @Success		200		{object}	models.BaseResponse[any]
// @Router			/jadwal-kuliah/save-trans [post]
// @Security BearerAuth
func (j *JadwalKuliahHandler) SaveTransJadwalKuliah(c *gin.Context) {
	var body models.JadwalKuliahRequest

	if err := c.BindJSON(&body); err != nil {
		c.JSON(400, models.BaseResponse[entities.Todo]{
			Message: "error",
			Errors:  []any{err.Error()},
		})
		return
	}

	err := j.service.SaveTransJadwalKuliah(&body, 1) // replace this value to session(semester_id)
	if err != nil {
		c.JSON(500, models.BaseResponse[entities.JadwalKuliah]{
			Message: "error",
			Errors:  []any{err.Error()},
		})
		return
	}

	c.JSON(200, models.BaseResponse[[]entities.JadwalKuliah]{
		Message: "success",
		Data:    nil,
	})
}

func (j *JadwalKuliahHandler) DeleteJadwalKuliah(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(500, models.BaseResponse[entities.Todo]{
			Message: "error",
			Errors:  []any{err.Error()},
		})
		return
	}

	err = j.service.DeleteJadwalKuliah(id) // replace this value to session(semester_id)
	if err != nil {
		c.JSON(500, models.BaseResponse[entities.JadwalKuliah]{
			Message: "error",
			Errors:  []any{err.Error()},
		})
		return
	}

	c.JSON(200, models.BaseResponse[[]entities.JadwalKuliah]{
		Message: "success",
		Data:    nil,
	})
}
