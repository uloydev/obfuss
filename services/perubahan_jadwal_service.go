package services

import (
	"errors"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"skripsi.id/obfuss/entities"
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

func (s *PerubahanJadwalService) GetJadwalById(id int) (entities.JadwalPertemuan, error) {
	var jadwalPertemuan entities.JadwalPertemuan
	err := s.db.Table(jadwalPertemuan.TableName()).
		Where("id = ?", id).
		First(&jadwalPertemuan).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		s.logger.Error("failed to find jadwal pertemuan by id", zap.Error(err))
		return jadwalPertemuan, errors.New("failed to find jadwal pertemuan by id")
	}

	return jadwalPertemuan, err
}

func (s *PerubahanJadwalService) SavePerubahanJadwal(payload *entities.JadwalPertemuan) error {
	tx := s.db.Begin()
	err := tx.Table(payload.TableName()).Where("id = ?", payload.ID).Updates(payload).Error
	if err != nil {
		s.logger.Error("failed to save perubahan jadwal", zap.Error(err))
		tx.Rollback()
		return errors.New("failed to save perubahan jadwal")
	}

	tx.Commit()
	return nil
}
