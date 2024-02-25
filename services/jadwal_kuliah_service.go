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

	j.logger.Info("debug pagination: ", zap.Any("meta", meta), zap.Any("jadwal", jadwalKuliah))

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

func (j *JadwalKuliahService) saveAuto(tx *gorm.DB, dayName string, idJadwal int, idJamMulai int, idJamSelesai int) error {
	day := utils.GetDayName(dayName)

	if err := queries.DeleteTransJadwalKuliah(tx, idJadwal).Error; err != nil {
		j.logger.Warn(err.Error())
		tx.Rollback()

		return errors.New("error when delete trans jadwal kuliah")
	}

	if err := tx.Exec("SET @row_number=0;").Error; err != nil {
		j.logger.Warn(err.Error())
		tx.Rollback()

		return errors.New("error when set row number")
	}

	subQuery := tx.Table("date_calendar").
		Select("date").
		Joins("LEFT JOIN mst_tanggal_off ON mst_tanggal_off.tanggal = date_calendar.date").
		Where("DAYNAME(date_calendar.date) = ?", day)

	err := tx.Exec("INSERT INTO trans_jadwal_pertemuan (id_jadwal, pertemuan_ke, tanggal_pertemuan, mulai_jam, sampai_jam, keterangan, add_date, add_user)"+
		"SELECT ?, (@row_number := @row_number +1), cal.date, ?, ?, mst_tanggal_off.keterangan, NOW(), 'SYSTEM' "+
		"FROM (?) AS cal", idJadwal, idJamMulai, idJamSelesai, subQuery).Error

	if err != nil {
		j.logger.Warn(err.Error())
		tx.Rollback()
		return errors.New("error when insert trans jadwal kuliah pertemuan")
	}

	return nil
}

func (j *JadwalKuliahService) SaveTransJadwalKuliah(payload *models.JadwalKuliahRequest, userId int) error {
	var (
		jadwalKuliah entities.JadwalKuliah
	)

	var jadwalKuliahs []entities.JadwalKuliah
	var trx = j.db.Begin()

	err := trx.Table(jadwalKuliah.TableName()).
		Where("id_kelas = ?", &payload.IdKelas).
		Delete(&jadwalKuliahs).Error

	if err != nil {
		j.logger.Error(err.Error())

		trx.Rollback()
		return errors.New("failed when delete jadwal kuliah by id kelas")
	}

	for _, v := range payload.JadwalKuliah {
		var data = entities.JadwalKuliah{}

		data.IDKelas = &payload.IdKelas
		data.IDMk = &v.IDMk
		data.IDDosen = &v.IDDosen
		data.Hari = &v.Hari
		data.IDJamMulai = &v.IDJamMulai
		data.IDJamSelesai = &v.IDJamSelesai
		data.AddUser = userId

		if err := trx.Table(jadwalKuliah.TableName()).Create(&data).Error; err != nil {
			j.logger.Error(err.Error())
			trx.Rollback()

			return errors.New("failed when save trans jadwal kuliah")
		}

		// Function gak jelas
		if err := j.saveAuto(trx, v.Hari, data.ID, v.IDJamMulai, v.IDJamSelesai); err != nil {
			return err
		}
	}

	trx.Commit()

	return nil
}

func (j *JadwalKuliahService) GetById(id int) (entities.JadwalKuliah, error) {
	var data entities.JadwalKuliah
	err := j.db.Table(entities.JadwalKuliah{}.TableName()).First(&data, id).Error
	if err != nil {
		j.logger.Error("failed to get jadwal kuliah by id", zap.Error(err))
		return data, errors.New("failed to get jadwal kuliah by id")
	}

	return data, nil
}
