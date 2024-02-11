package services

import (
	"database/sql"
	"errors"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"skripsi.id/obfuss/entities"
	"skripsi.id/obfuss/models"
	"skripsi.id/obfuss/queries"
	"skripsi.id/obfuss/utils"
)

type JadwalKuliahService struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewJadwalKuliahService(db *gorm.DB, logger *zap.Logger) *JadwalKuliahService {
	return &JadwalKuliahService{
		db,
		logger.With(zap.String("type", "TodoService")),
	}
}

func (j *JadwalKuliahService) GetJadwalKuliah(pageParams models.PaginationParams, idSemester uint) ([]entities.JadwalKuliah, *models.PaginationMeta, error) {
	query := queries.FindAllJadwalKuliah(j.db, idSemester)

	meta, jadwalKuliah, err := utils.Paginate[entities.JadwalKuliah](pageParams, query, j.logger)

	if err != nil {
		return nil, nil, err
	}

	return jadwalKuliah, meta, nil
}

func (j *JadwalKuliahService) DeleteJadwalKuliah(id int) error {
	var jadwalKuliah entities.JadwalKuliah

	err := j.db.Table(jadwalKuliah.TableName()).
		Where("id_kelas = @id_kelas", sql.Named("id_kelas", id)).
		Delete(jadwalKuliah).Error

	if err != nil {
		return errors.New("error when delete jadwal kuliah")
	}

	return nil
}

func (j *JadwalKuliahService) SaveTransJadwalKuliah(payload *models.JadwalKuliahRequest, userId int) error {
	var (
		jadwalKuliah      entities.JadwalKuliah
		batchJadwalKuliah []entities.JadwalKuliah
	)

	for _, v := range payload.JadwalKuliah {
		var data = entities.JadwalKuliah{}

		data.IDKelas = &payload.IdKelas
		data.IDMk = &v.IDMk
		data.IDDosen = &v.IDDosen
		data.Hari = &v.Hari
		data.IDJamMulai = &v.IDJamMulai
		data.IDJamSelesai = &v.IDJamSelesai
		data.AddUser = userId

		batchJadwalKuliah = append(batchJadwalKuliah, data)
	}

	if err := j.db.Table(jadwalKuliah.TableName()).Create(&batchJadwalKuliah).Error; err != nil {
		j.logger.Error(err.Error())
		return errors.New("failed when save trans jadwal kuliah")
	}

	return nil
}
