package queries

import "gorm.io/gorm"

func GetMahasiswaRaw(db *gorm.DB, idKelas int, idPertemuan int) *gorm.DB {
	return db.Table("mst_mahasiswa").
		Select("DISTINCT trans_absen_mahasiswa.is_hadir, mst_mahasiswa.*").
		Joins("JOIN trans_plot_kelas ON trans_plot_kelas.id_mahasiswa = mst_mahasiswa.id AND trans_plot_kelas.id_kelas = ?", idKelas).
		Joins("LEFT JOIN trans_absen_mahasiswa ON trans_absen_mahasiswa.id_mahasiswa = mst_mahasiswa.id AND trans_absen_mahasiswa.id_pertemuan = ?", idPertemuan)
}
