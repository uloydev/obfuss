package queries

import (
	"gorm.io/gorm"
	"skripsi.id/obfuss/entities"
)

func FindAllJadwalKuliah(
	db *gorm.DB,
	smtId uint,
) *gorm.DB {

	return db.Table(entities.JadwalKuliah{}.TableName()).
		Select("mst_kurikulum.kode_kurikulum, trans_jadwal_kuliah.id_kelas, COUNT(trans_jadwal_kuliah.id_mk) AS jumlah_mk, "+
			"mst_kelas.inisial_kelas, mst_kelas.nama_kelas, mst_semester.semester, "+
			"mst_tahun_ajaran.tahun, mst_konsentrasi.nama_prodi, mst_konsentrasi.nama_konsentrasi").
		Joins("JOIN mst_kelas ON mst_kelas.id = trans_jadwal_kuliah.id_kelas").
		Joins("JOIN mst_kurikulum ON mst_kurikulum.id = mst_kelas.id_kurikulum").
		Joins("JOIN mst_konsentrasi ON mst_konsentrasi.id = mst_kelas.id_konsentrasi").
		Joins("JOIN mst_semester ON mst_semester.id = mst_kelas.id_semester").
		Joins("JOIN mst_tahun_ajaran ON mst_tahun_ajaran.id = mst_semester.id_tahun").
		Where("mst_semester.id = ?", smtId).
		Group("mst_kurikulum.kode_kurikulum, trans_jadwal_kuliah.id_kelas, " +
			"mst_kelas.inisial_kelas, mst_kelas.nama_kelas, mst_semester.semester, " +
			"mst_tahun_ajaran.tahun, mst_konsentrasi.nama_prodi, mst_konsentrasi.nama_konsentrasi")
}
