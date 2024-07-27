package services

import (
	"errors"
	"fmt"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"skripsi.id/obfuss/entities"
	"skripsi.id/obfuss/models"
	"skripsi.id/obfuss/queries"
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
	userType string,
	smtId, kelasId, dosenId int,
) (
	[]models.GetAllAbsenResponse,
	error,
) {
	query := queries.FindAllAbsenQuery(s.db, userType, smtId, kelasId, dosenId)
	data := []models.GetAllAbsenResponse{}
	err := query.Find(&data).Error

	if err != nil {
		s.logger.Error("failed to get absen mahasiswa", zap.Error(err))
		return nil, err
	}

	return data, err
}

func (s *AbsenService) DeleteAbsenByPertemuan(idPertemuan int) error {
	tx := s.db.Begin()
	err := tx.Table(entities.AbsenMahasiswa{}.TableName()).
		Where("id_pertemuan = ?", idPertemuan).
		Delete(entities.AbsenMahasiswa{}).Error
	if err != nil {
		s.logger.Error("failed to delete absen by pertemuan", zap.Error(err))
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}

func (s *AbsenService) SaveAbsenTrans(
	req models.SaveAbsenTransRequest,
	actorId int,
) ([]entities.AbsenMahasiswa, error) {
	tx := s.db.Begin()
	data := []entities.AbsenMahasiswa{}

	for _, absen := range req.Absen {
		data = append(data, entities.AbsenMahasiswa{
			IDPertemuan: req.IdPertemuan,
			IDMahasiswa: absen.IdMahasiswa,
			IsHadir:     absen.IsHadir,
			Keterangan:  absen.Keterangan,
			AddUser:     actorId,
		})
	}
	results := tx.Table(entities.AbsenMahasiswa{}.TableName()).Create(&data)
	err := results.Error

	if err != nil {
		tx.Rollback()
		s.logger.Error("failed to save absen trans", zap.Error(err))
		return nil, errors.New("failed to save absen trans")
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		s.logger.Error("failed to save absen trans", zap.Error(err))
		return nil, errors.New("failed to save absen trans")
	}
	return data, nil
}

func (s *AbsenService) CountAbsen(idPertemuan int) (models.AbsenCountResult, error) {
	var data models.AbsenCountResult
	err := queries.CountAbsenQuery(s.db, idPertemuan).Scan(&data).Error
	if err != nil {
		s.logger.Error("failed to count absen", zap.Error(err))
		return data, errors.New("failed to count absen")
	}

	return data, nil
}

func (s *AbsenService) DeleteByPertemuanID(idPertemuan int) error {
	results := s.db.Table(entities.AbsenMahasiswa{}.TableName()).
		Where("id_pertemuan = ?", idPertemuan).
		Delete(entities.AbsenMahasiswa{})

	err := results.Error

	if err != nil {
		s.logger.Error("failed to delete absen by pertemuan", zap.Error(err))
		return err
	} else if results.RowsAffected < 1 {
		s.logger.Log(zap.InfoLevel, fmt.Sprintf("cannot find any record with give id %d", idPertemuan))

		return gorm.ErrRecordNotFound
	}

	return nil
}
