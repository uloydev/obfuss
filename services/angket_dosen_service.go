package services

import (
	"errors"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"skripsi.id/obfuss/entities"
)

type AngketDosenService struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewAngketDosenService(db *gorm.DB, logger *zap.Logger) *AngketDosenService {
	return &AngketDosenService{
		db:     db,
		logger: logger.With(zap.String("type", "AngketDosenService")),
	}
}

func (s *AngketDosenService) FineOneByPertemuan(idPertemuan int) (entities.AngketDosen, error) {
	var angket entities.AngketDosen
	err := s.db.Table(angket.TableName()).
		Where("id_pertemuan = ?", idPertemuan).
		First(&angket).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		s.logger.Error("failed to find angket by pertemuan", zap.Error(err))
		return angket, errors.New("failed to find angket by pertemuan")
	}

	return angket, err
}

func (s *AngketDosenService) SaveAngketDosen(payload *entities.AngketDosen) error {
	tx := s.db.Begin()
	err := tx.Table(payload.TableName()).Create(payload).Error
	if err != nil {
		s.logger.Error("failed to save angket dosen", zap.Error(err))
		tx.Rollback()
		return errors.New("failed to save angket dosen")
	}

	tx.Commit()
	return nil
}

func (s *AngketDosenService) DeleteByPertemuanID(idPertemuan int) error {
	tx := s.db.Begin()
	err := tx.Table(entities.AngketDosen{}.TableName()).
		Where("id_pertemuan = ?", idPertemuan).
		Delete(entities.AngketDosen{}).Error
	if err != nil {
		s.logger.Error("failed to delete angket dosen by pertemuan", zap.Error(err))
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}
