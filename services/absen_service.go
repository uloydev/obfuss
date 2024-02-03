package services

import (
	"errors"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"skripsi.id/obfuss/models"
	"skripsi.id/obfuss/queries"
	"skripsi.id/obfuss/utils"
)

type AbsenService struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewAbsenService(db *gorm.DB, logger *zap.Logger) *AbsenService {
	return &AbsenService{
		db:     db,
		logger: logger.With(zap.String("type", "AbsenService")),
	}
}

func (s *AbsenService) GetAllAbsen(
	pageParams models.PaginationParams,
	userType string,
	smtId, kelasId, dosenId int,
) (
	[]models.GetAllAbsenResponse,
	*models.PaginationMeta,
	error,
) {
	query := queries.FindAllAbsenQuery(s.db, userType, smtId, kelasId, dosenId)
	meta, data, err := utils.Paginate[models.GetAllAbsenResponse](pageParams, query, s.logger)

	if err != nil {
		s.logger.Error("failed to get absen mahasiswa", zap.Error(err))
		return nil, nil, errors.New("failed to get absen mahasiswa")
	}

	return data, meta, err
}
