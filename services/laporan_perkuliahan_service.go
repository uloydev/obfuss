package services

import (
	"errors"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"skripsi.id/obfuss/entities"
	"skripsi.id/obfuss/middlewares"
	"skripsi.id/obfuss/models"
	"skripsi.id/obfuss/queries"
	"skripsi.id/obfuss/utils"
)

type LaporanPerkuliahanService struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewLaporanPerkuliahanService(db *gorm.DB, logger *zap.Logger) *LaporanPerkuliahanService {
	return &LaporanPerkuliahanService{
		db,
		logger.With(zap.String("type", "LaporanPerkuliahanService")),
	}
}

func (s *LaporanPerkuliahanService) SaveTrans(payload *models.SaveTransLaporanPerkuliahan, user *middlewares.User) error {
	var angketDosen *entities.AngketDosen

	err := s.db.Model(&angketDosen).First(&angketDosen, map[string]any{
		"id_pertemuan": &payload.IDPertemuan,
	}).Error

	if err != nil {
		s.logger.Error("error when find angket dosen :", zap.Any("error", err))

		return errors.New("internal server error")
	}

	currentTime := time.Now()

	angketDosen.ModifiedDate = &currentTime
	angketDosen.ModifiedUser = &user.ActorID
	angketDosen.IDDosen = payload.IDDosen
	angketDosen.PertemuanKe = &payload.PertemuanKe
	angketDosen.JamKuliah = &payload.JamKuliah
	angketDosen.Keterangan = &payload.Keterangan
	angketDosen.RingkasanMateri = &payload.RingkasanMateri

	if payload.FileName != "" {
		angketDosen.FileGambar = &payload.FileName
	}

	if err := s.db.Model(&angketDosen).Save(angketDosen).Error; err != nil {
		s.logger.Error("error when update angket dosen :", zap.Any("error", err))

		return errors.New("internal server error")
	}

	return nil
}

func (s *LaporanPerkuliahanService) GetAllAngeketDosen(
	pageParams models.PaginationParams,
	user *middlewares.User,
) (
	[]map[string]any,
	*models.PaginationMeta,
	error,
) {
	var kelasMhs []map[string]interface{}

	query := queries.GetKelasIDFromTPK(s.db, user)
	query.Find(&kelasMhs).Limit(1)
	idKelas := kelasMhs[0]["id_kelas"].(int32)

	query = queries.GetAllLaporanKuliahByUsertype(s.db, int(idKelas), user)
	meta, data, err := utils.Paginate[map[string]interface{}](pageParams, query, s.logger)

	if err != nil {
		s.logger.Error("failed to get laporan perkuliahan", zap.Error(err))
		return nil, nil, errors.New("failed to get laporan perkuliahan")
	}

	return data, meta, err
}
