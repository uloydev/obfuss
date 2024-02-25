package services

import (
	"errors"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"skripsi.id/obfuss/entities"
)

type JadwalPertemuanService struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewJadwalPertemuanService(db *gorm.DB, logger *zap.Logger) *JadwalPertemuanService {
	return &JadwalPertemuanService{
		db:     db,
		logger: logger.With(zap.String("type", "JadwalPertemuanService")),
	}
}

func (s *JadwalPertemuanService) GetById(id int) (entities.JadwalPertemuan, error) {
	var data entities.JadwalPertemuan
	err := s.db.Table(entities.JadwalPertemuan{}.TableName()).First(&data, id).Error
	if err != nil {
		s.logger.Error("failed to get jadwal pertemuan by id", zap.Error(err))
		return data, errors.New("failed to get jadwal pertemuan by id")
	}

	return data, nil
}

func (s *JadwalPertemuanService) UpdateById(id int, data entities.JadwalPertemuan) error {
	tx := s.db.Begin()
	err := tx.Table(entities.JadwalPertemuan{}.TableName()).Where("id = ?", id).Updates(data).Error
	if err != nil {
		s.logger.Error("failed to update jadwal pertemuan by id", zap.Error(err))
		tx.Rollback()
		return errors.New("failed to update jadwal pertemuan by id")
	}

	tx.Commit()
	return nil
}
