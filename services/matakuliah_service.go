package services

import (
	"errors"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"skripsi.id/obfuss/entities"
	"skripsi.id/obfuss/models"
)

type MataKuliahService struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewMataKuliahService(db *gorm.DB, logger *zap.Logger) *MataKuliahService {
	return &MataKuliahService{
		db:     db,
		logger: logger.With(zap.String("type", "MataKuliahService")),
	}
}

func (s *MataKuliahService) GetAll() (
	[]models.MataKuliah,
	error,
) {
	var entitiy []*entities.MataKuliah

	if err := s.db.Find(&entitiy).Error; err != nil {
		s.logger.Error("error while find matakuliah", zap.Error(err))

		return nil, errors.New("internal Server Error")
	}

	var responses []models.MataKuliah

	for _, v := range entitiy {
		responses = append(responses,
			models.MataKuliah{
				ID:             v.ID,
				KodeMk:         v.KodeMk,
				NamaMk:         v.NamaMk,
				Keterangan:     v.Keterangan,
				SemesterKuliah: v.SemesterKuliah,
				SKS:            v.SKS,
				Status:         v.Status,
				JenisMk:        v.JenisMk,
				IsAgama:        v.IsAgama,
				IsTaPkl:        v.IsTaPkl,
				SksPraktek:     v.SksPraktek,
				SksLapangan:    v.SksLapangan,
				SksSimulasi:    v.SksSimulasi,
			},
		)

	}

	return responses, nil
}
