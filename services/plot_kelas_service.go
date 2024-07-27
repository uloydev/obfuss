package services

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"skripsi.id/obfuss/entities"
)

type PlotKelasService struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewPlotKelasService(db *gorm.DB, logger *zap.Logger) *PlotKelasService {
	return &PlotKelasService{
		db:     db,
		logger: logger.With(zap.String("type", "PlotKelasService")),
	}
}

func (s *PlotKelasService) GetIdKelas(conds map[string]any) (int, error) {
	var plotKelas entities.PlotKelas
	err := s.db.Table(plotKelas.TableName()).Where(conds).First(&plotKelas).Error

	if err != nil {
		s.logger.Error("failed to get id kelas", zap.Error(err))
		return 0, err
	}

	return plotKelas.IDKelas, nil
}
