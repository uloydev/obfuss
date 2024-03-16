package services

import (
	"database/sql"
	"errors"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"skripsi.id/obfuss/entities"
	"skripsi.id/obfuss/models"
	"skripsi.id/obfuss/queries"
	"skripsi.id/obfuss/utils"
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

func (s *JadwalKuliahDosenService) GetGridJadwalKuliahDosen() ([]entities.Kelas, error) {
	var results []entities.Kelas

	err := s.db.Joins("Kurikulum").
		Joins("Semester").
		Joins("Konsentrasi").
		Find(&results).Error

	if err != nil {
		s.logger.Error("failed to get grid jadwal kuliah dosen", zap.Error(err))
		return nil, errors.New("failed to get grid jadwal kuliah dosen")
	}

	return results, nil
}

func (s *JadwalKuliahDosenService) GetAllJadwalKuliahDosen(
	pageParams models.PaginationParams,
	idDosen int,
) (
	[]map[string]any,
	*models.PaginationMeta,
	error,
) {
	var data []map[string]any

	query := queries.FindAllJadwalKuliahDosen(s.db, idDosen)

	query.Find(&data)

	s.logger.Info("logging", zap.Any("data", data))

	meta, data, err := utils.Paginate[map[string]interface{}](pageParams, query, s.logger)

	if err != nil {
		s.logger.Error("failed to get jadwal kuliah dosen", zap.Error(err))
		return nil, nil, errors.New("failed to get jadwal kuliah dosen")
	}

	s.logger.Info("logging", zap.Any("data", data))

	return data, meta, err
}

func (s *JadwalKuliahDosenService) GetById(id int) (entities.JadwalKuliahDosen, error) {
	var data entities.JadwalKuliahDosen

	err := s.db.Table(entities.JadwalKuliahDosen{}.TableName()).First(&data, id).Error
	if err != nil {
		s.logger.Error("failed to get jadwal kuliah dosen by id", zap.Error(err))
		return data, errors.New("failed to get jadwal kuliah dosen by id")
	}

	return data, nil
}

func (s *JadwalKuliahDosenService) SaveTrans(payload *models.SaveTransJadwalKuliahDosen, userId int) error {
	var (
		errorFind   = errors.New("cannot find jadwal kuliah dosen")
		errorUpdate = errors.New("failed when update jadwal kuliah dosen")
	)

	var jadwalKuliahDosen entities.JadwalKuliahDosen

	err := s.db.
		Table(jadwalKuliahDosen.TableName()).
		Find(&jadwalKuliahDosen, map[string]any{
			"id": &payload.ID,
		}).Error

	if err != nil {
		s.logger.Error(err.Error())
		return errorFind
	}

	jadwalKuliahDosen.Hari = &payload.Hari
	jadwalKuliahDosen.IDDosen = &payload.IDDosen
	jadwalKuliahDosen.IDSemester = &payload.IDSemester
	jadwalKuliahDosen.IDJamMulai = &payload.IDJamMulai
	jadwalKuliahDosen.IDJamSelesai = &payload.IDJamSelesai
	jadwalKuliahDosen.IDKelas = &payload.IDKelas
	jadwalKuliahDosen.IDMk = &payload.IDMK

	// meta
	currentTime := time.Now()
	jadwalKuliahDosen.ModifiedDate = currentTime
	jadwalKuliahDosen.ModifiedUser = userId

	if err := s.db.Save(&jadwalKuliahDosen).Error; err != nil {
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
