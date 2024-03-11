package handlers

import (
	"errors"
	"fmt"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"skripsi.id/obfuss/entities"
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

// @Summary		SaveTrans
// @Description	SaveTrans
// @Tags			Laporan Perkuliahan
// @Accept			x-www-form-urlencoded
// @Produce		json
// @Param file formData models.SaveTransLaporanPerkuliahan true "formData"
// @Param foto formData file false "file"
// @Success		200		{object}	models.BaseResponse[any]
// @Router		/laporan-perkuliahan/save-trans [post]
// @Security BearerAuth
func (h *LaporanPerkuliahanHandler) SaveTransLaporanPerkuliahan(c *gin.Context) {
	var formData *models.SaveTransLaporanPerkuliahan

	if err := c.ShouldBind(&formData); err != nil {
		c.JSON(400, models.BaseResponse[any]{
			Message: "error",
			Errors:  []any{err.Error()},
		})
		return
	}

	user, err := utils.GetUser(c)

	if err != nil {
		c.JSON(401, models.BaseResponse[any]{
			Message: "error",
			Errors:  []any{err.Error()},
		})

		return
	}

	if formData.Foto != nil {
		ext := filepath.Ext(formData.Foto.Filename)
		fotoName := fmt.Sprintf("%d%s", formData.IDPertemuan, ext)
		formData.FileName = fotoName

		if err := c.SaveUploadedFile(formData.Foto, filepath.Join("storage/foto-kuliah/", fotoName)); err != nil {
			h.logger.Error("error:", zap.Any("while store data", err.Error()))

			c.JSON(500, models.BaseResponse[any]{
				Message: "error",
				Errors:  []any{"internal server error"},
			})

			return
		}
	}

	err = h.service.SaveTrans(formData, user)

	if err != nil {
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

// @Summary		Delete Angket Dosen
// @Description	Delete Angket Dosen
// @Tags		Laporan Perkuliahan
// @Accept			json
// @Produce		json
// @Param			id-pertemuan	path		int	true	"pertemuan ID"
// @Success		204	{object}	models.BaseResponse[any]
// @Router			/laporan-perkuliahan/{id-pertemuan}/pertemuan [delete]
func (h *LaporanPerkuliahanHandler) Delete(c *gin.Context) {
	idPertemuan, err := strconv.Atoi(c.Param("id-pertemuan"))
	if err != nil {
		c.JSON(500, models.BaseResponse[any]{
			Message: "error",
			Errors:  []any{err.Error()},
		})
		return
	}

	if err := h.service.Delete(idPertemuan); err != nil {
		c.JSON(500, models.BaseResponse[any]{
			Message: "error",
			Errors:  []any{err.Error()},
		})
		return
	}

	c.JSON(200, models.BaseResponse[entities.Todo]{
		Message: "success",
	})
}
