package queries

import (
	"gorm.io/gorm"
	"skripsi.id/obfuss/middlewares"
)

func GetAllLaporanKuliahByUsertype(db *gorm.DB, kelasMhs int, user *middlewares.User) *gorm.DB {
	query := db.Table("trans_angket_dosen").
		Select("trans_angket_dosen.*, mst_semester.semester, mst_tahun_ajaran.tahun, mst_kelas.id as id_kelas, mst_kelas.inisial_kelas, mst_mata_kuliah.nama_mk, mst_mata_kuliah.id as id_mk, trans_jadwal_pertemuan.id as id_pertemuan, DATE_FORMAT(trans_jadwal_pertemuan.tanggal_pertemuan,'%d/%m/%Y') AS tanggal_pertemuan, mulai_jam.jam_mulai, selesai_jam.jam_selesai, mst_dosen.nama_lengkap").
		Joins("JOIN trans_jadwal_pertemuan ON trans_angket_dosen.id_pertemuan = trans_jadwal_pertemuan.id").
		Joins("JOIN mst_jam_kuliah as mulai_jam ON trans_jadwal_pertemuan.mulai_jam = mulai_jam.no_urut").
		Joins("JOIN mst_jam_kuliah as selesai_jam ON trans_jadwal_pertemuan.sampai_jam = selesai_jam.no_urut").
		Joins("JOIN trans_jadwal_kuliah ON trans_jadwal_pertemuan.id_jadwal = trans_jadwal_kuliah.id").
		Joins("JOIN mst_kelas ON trans_jadwal_kuliah.id_kelas = mst_kelas.id").
		Joins("JOIN mst_semester ON mst_semester.id=mst_kelas.id_semester").
		Joins("JOIN mst_tahun_ajaran ON mst_tahun_ajaran.id=mst_semester.id_tahun").
		Joins("JOIN mst_mata_kuliah ON trans_jadwal_kuliah.id_mk = mst_mata_kuliah.id").
		Joins("JOIN mst_dosen ON trans_jadwal_kuliah.id_dosen = mst_dosen.id")

	if user.UserType == "dosen" {
		query = query.Where("trans_jadwal_kuliah.id_dosen = ?", &user.ActorID)
	} else if user.UserType == "mahasiswa" {
		query = query.Where("mst_kelas.id = ?", kelasMhs).Where("mst_semester.id = ?", &user.SemesterId)
	}
	return query
}
