package queries

import (
	"gorm.io/gorm"
)

func FindAllJadwalKuliahDosen(db *gorm.DB, semesterId int) *gorm.DB {
	return db.Table("v_jadwal_dosen_v1").
		Select("*").
		Where("id_semester = ?", semesterId)
}
