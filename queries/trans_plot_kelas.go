package queries

import (
	"gorm.io/gorm"
	"skripsi.id/obfuss/middlewares"
)

func GetKelasIDFromTPK(db *gorm.DB, user *middlewares.User) *gorm.DB {
	return db.Table("trans_plot_kelas").
		Where("id_semester = ?", user.SemesterId).
		Where("id_mahasiswa = ?", 1).
		Select("id_kelas")
}
