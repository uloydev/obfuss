package services

import (
	"errors"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"skripsi.id/obfuss/queries"
)

type MahasiswaService struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewMahasiswaService(db *gorm.DB, logger *zap.Logger) *MahasiswaService {
	return &MahasiswaService{
		db,
		logger.With(zap.String("type", "MahasiswaService")),
	}
}

func (s *MahasiswaService) GetMahasiswaRaw(
	idKelas int,
	idPertemuan int,
) (
	[]map[string]any,
	error,
) {
	data := []map[string]any{}
	query := queries.GetMahasiswaRaw(s.db, idKelas, idPertemuan)

	err := query.Find(&data).Error

	if err != nil {
		s.logger.Error("failed to get data mahasiswa", zap.Error(err))
		return nil, errors.New("failed to get jadwal kuliah dosen")
	}

	return data, err
}
