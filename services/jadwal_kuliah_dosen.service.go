package services

import (
	"database/sql"
	"errors"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"skripsi.id/obfuss/entities"
	"skripsi.id/obfuss/models"
)

type JadwalKuliahDosenService struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewJadwalKuliahDosenService(db *gorm.DB, logger *zap.Logger) *JadwalKuliahDosenService {
	return &JadwalKuliahDosenService{
		db,
		logger.With(zap.String("type", "JadwalKuliahDosenService")),
	}
}

func (s *JadwalKuliahDosenService) SaveTrans(payload models.SaveTransJadwalKuliahDosen) error {
	var (
		errorCreate = errors.New("failed when create jadwal kuliah")
		errorUpdate = errors.New("failed when update jadwal kuliah dosen")
	)

	var jadwalKuliah entities.JadwalKuliah

	err := s.db.
		Table(jadwalKuliah.TableName()).
		Model(&jadwalKuliah).
		FirstOrCreate(&jadwalKuliah, map[string]int{"id": payload.ID}).Error
	if err != nil {
		s.logger.Error(err.Error())

		return errorCreate
	}

	jadwalKuliah.Hari = &payload.Hari
	jadwalKuliah.IDDosen = &payload.IDDosen

	jadwalKuliah.IDJamMulai = &payload.IDJamMulai
	jadwalKuliah.IDJamSelesai = &payload.IDJamSelesai
	jadwalKuliah.IDKelas = &payload.IDKelas
	jadwalKuliah.IDMk = &payload.IDMK

	if err := s.db.Save(&jadwalKuliah).Error; err != nil {
		s.logger.Error(err.Error())

		return errorUpdate
	}

	return nil
}

func (s *JadwalKuliahDosenService) Delete(id int) error {
	var jadwalPertemuan *entities.JadwalPertemuan

	err := s.db.Table(jadwalPertemuan.TableName()).
		Where("id_kelas = @id_kelas", sql.Named("id_kelas", id)).
		Delete(jadwalPertemuan).Error

	if err != nil {
		return errors.New("error when delete jadwal kuliah")
	}

	return nil
}
