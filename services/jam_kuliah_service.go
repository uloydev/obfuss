package services

import (
	"errors"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"skripsi.id/obfuss/entities"
	"skripsi.id/obfuss/models"
)

type JamKuliahService struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewJamKuliahService(db *gorm.DB, logger *zap.Logger) *JamKuliahService {
	return &JamKuliahService{
		db,
		logger.With(zap.String("type", "JamKuliahService")),
	}
}

func (s *JamKuliahService) GetAll() ([]models.GetAllJamKuliah, error) {
	var response []models.GetAllJamKuliah

	if err := s.db.Table(entities.JamKuliah{}.TableName()).Find(&response).Error; err != nil {
		s.logger.Error("error while find jam kuliah", zap.Error(err))
		return nil, errors.New("internal Server Error")
	}

	return response, nil
}
