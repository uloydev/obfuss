package services

import (
	"errors"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"skripsi.id/obfuss/models"
	"skripsi.id/obfuss/queries"
	"skripsi.id/obfuss/utils"
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
	pageParams models.PaginationParams,
	idKelas int,
	idPertemuan int,
) (
	[]map[string]any,
	*models.PaginationMeta,
	error,
) {
	query := queries.GetMahasiswaRaw(s.db, idKelas, idPertemuan)
	meta, data, err := utils.Paginate[map[string]interface{}](pageParams, query, s.logger)

	if err != nil {
		s.logger.Error("failed to get data mahasiswa", zap.Error(err))
		return nil, nil, errors.New("failed to get jadwal kuliah dosen")
	}

	return data, meta, err
}
