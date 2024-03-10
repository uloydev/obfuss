package services

import (
	"errors"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"skripsi.id/obfuss/middlewares"
	"skripsi.id/obfuss/models"
	"skripsi.id/obfuss/queries"
	"skripsi.id/obfuss/utils"
)

type LaporanPerkuliahanService struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewLaporanPerkuliahanService(db *gorm.DB, logger *zap.Logger) *LaporanPerkuliahanService {
	return &LaporanPerkuliahanService{
		db,
		logger.With(zap.String("type", "LaporanPerkuliahanService")),
	}
}

func (s *LaporanPerkuliahanService) GetAllAngeketDosen(
	pageParams models.PaginationParams,
	user *middlewares.User,
) (
	[]map[string]any,
	*models.PaginationMeta,
	error,
) {
	var kelasMhs []map[string]interface{}

	query := queries.GetKelasIDFromTPK(s.db, user)
	query.Find(&kelasMhs).Limit(1)
	idKelas := kelasMhs[0]["id_kelas"].(int32)

	query = queries.GetAllLaporanKuliahByUsertype(s.db, int(idKelas), user)
	meta, data, err := utils.Paginate[map[string]interface{}](pageParams, query, s.logger)

	if err != nil {
		s.logger.Error("failed to get laporan perkuliahan", zap.Error(err))
		return nil, nil, errors.New("failed to get laporan perkuliahan")
	}

	return data, meta, err
}
