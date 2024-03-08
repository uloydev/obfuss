package services

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"skripsi.id/obfuss/models"
	"skripsi.id/obfuss/queries"
	"skripsi.id/obfuss/utils"
)

type PerubahanJadwalService struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewPerubahanJadwalService(db *gorm.DB, logger *zap.Logger) *PerubahanJadwalService {
	return &PerubahanJadwalService{
		db:     db,
		logger: logger.With(zap.String("type", "PerubahanJadwalService")),
	}
}

func (s *PerubahanJadwalService) GetPerubahanJadwal(pageParams models.PaginationParams, idSemester uint) ([]models.GetAllPerubahanJadwal, *models.PaginationMeta, error) {
	query := queries.FindAllPerubahanData(s.db, idSemester)
	meta, perubahanJadwal, err := utils.Paginate[models.GetAllPerubahanJadwal](pageParams, query, s.logger)

	if err != nil {
		return nil, nil, err
	}

	return perubahanJadwal, meta, nil
}
