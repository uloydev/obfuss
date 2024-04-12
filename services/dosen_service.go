package services

import (
	"errors"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"skripsi.id/obfuss/entities"
	"skripsi.id/obfuss/models"
)

type DosenService struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewDosenService(db *gorm.DB, logger *zap.Logger) *DosenService {
	return &DosenService{
		db:     db,
		logger: logger.With(zap.String("type", "DosenService")),
	}
}

func (s *DosenService) GetAll() (
	[]models.GetAllDosen,
	error,
) {
	var response []models.GetAllDosen

	if err := s.db.Table(entities.Dosen{}.TableName()).Find(&response).Error; err != nil {
		s.logger.Error("error while find Dosen", zap.Error(err))

		return nil, errors.New("internal Server Error")
	}

	return response, nil
}
