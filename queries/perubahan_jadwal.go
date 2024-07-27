package queries

import (
	"gorm.io/gorm"
)

func FindAllPerubahanData(
	db *gorm.DB,
	smtId uint,
) *gorm.DB {
	return db.Table("trans_jadwal_pertemuan").
		Select("trans_jadwal_pertemuan.id, mst_dosen.nama_lengkap AS nama_dosen, trans_jadwal_kuliah.id_kelas, mst_kurikulum.kode_kurikulum, trans_jadwal_kuliah.id_mk, mst_mata_kuliah.nama_mk, mst_kelas.inisial_kelas, mst_kelas.nama_kelas, mst_semester.semester, mst_tahun_ajaran.tahun, mst_konsentrasi.nama_prodi, mst_konsentrasi.nama_konsentrasi, trans_jadwal_pertemuan.tanggal_pertemuan, CONCAT(SUBSTR(jam_masuk.jam_mulai,1,5), ' - ', SUBSTR(jam_keluar.jam_selesai,1,5)) AS jam_kuliah, trans_jadwal_pertemuan.tanggal_usulan_ganti, trans_jadwal_pertemuan.keterangan, CONCAT(SUBSTR(jam_masuk_old.jam_mulai,1,5), ' - ', SUBSTR(jam_selesai_old.jam_selesai,1,5)) AS jam_kuliah_old, trans_jadwal_pertemuan.alasan_perubahan, trans_jadwal_pertemuan.status_usulan, DATE_FORMAT(trans_jadwal_pertemuan.tanggal_pertemuan, '%d/%m/%Y') AS tanggal_kuliah, DATE_FORMAT(trans_jadwal_pertemuan.tanggal_usulan_ganti, '%d/%m/%Y') AS tanggal_perubahan, DATE_FORMAT(trans_jadwal_pertemuan.tanggal_old, '%d/%m/%Y') AS tanggal_old, CASE WHEN DAYNAME(trans_jadwal_pertemuan.tanggal_pertemuan) = 'Monday' THEN 'Senin' WHEN DAYNAME(trans_jadwal_pertemuan.tanggal_pertemuan) = 'Tuesday' THEN 'Selasa' WHEN DAYNAME(trans_jadwal_pertemuan.tanggal_pertemuan) = 'Wednesday' THEN 'Rabu' WHEN DAYNAME(trans_jadwal_pertemuan.tanggal_pertemuan) = 'Thursday' THEN 'Kamis' WHEN DAYNAME(trans_jadwal_pertemuan.tanggal_pertemuan) = 'Friday' THEN 'Jumat' WHEN DAYNAME(trans_jadwal_pertemuan.tanggal_pertemuan) = 'Saturday' THEN 'Sabtu' ELSE 'Minggu' END AS hari_jadwal, CASE WHEN DAYNAME(trans_jadwal_pertemuan.tanggal_old) = 'Monday' THEN 'Senin' WHEN DAYNAME(trans_jadwal_pertemuan.tanggal_old) = 'Tuesday' THEN 'Selasa' WHEN DAYNAME(trans_jadwal_pertemuan.tanggal_old) = 'Wednesday' THEN 'Rabu' WHEN DAYNAME(trans_jadwal_pertemuan.tanggal_old) = 'Thursday' THEN 'Kamis' WHEN DAYNAME(trans_jadwal_pertemuan.tanggal_old) = 'Friday' THEN 'Jumat' WHEN DAYNAME(trans_jadwal_pertemuan.tanggal_old) = 'Saturday' THEN 'Sabtu' ELSE 'Minggu' END AS hari_old").
		Joins("JOIN trans_jadwal_kuliah ON trans_jadwal_kuliah.id = trans_jadwal_pertemuan.id_jadwal").
		Joins("JOIN mst_kelas ON mst_kelas.id = trans_jadwal_kuliah.id_kelas").
		Joins("JOIN mst_kurikulum ON mst_kurikulum.id = mst_kelas.id_kurikulum").
		Joins("JOIN mst_konsentrasi ON mst_konsentrasi.id = mst_kelas.id_konsentrasi").
		Joins("JOIN mst_semester ON mst_semester.id = mst_kelas.id_semester").
		Joins("JOIN mst_tahun_ajaran ON mst_tahun_ajaran.id = mst_semester.id_tahun").
		Joins("JOIN mst_mata_kuliah ON mst_mata_kuliah.id = trans_jadwal_kuliah.id_mk").
		Joins("JOIN mst_dosen ON mst_dosen.id = trans_jadwal_kuliah.id_dosen").
		Joins("JOIN mst_jam_kuliah AS jam_masuk ON jam_masuk.id = trans_jadwal_pertemuan.mulai_jam").
		Joins("JOIN mst_jam_kuliah AS jam_keluar ON jam_keluar.id = trans_jadwal_pertemuan.sampai_jam").
		Joins("JOIN mst_jam_kuliah AS jam_masuk_old ON jam_masuk_old.id = trans_jadwal_pertemuan.mulai_jam_old").
		Joins("JOIN mst_jam_kuliah AS jam_selesai_old ON jam_selesai_old.id = trans_jadwal_pertemuan.sampai_jam_old").
		Joins("JOIN mst_jam_kuliah AS jam_keluar_old ON jam_keluar_old.id = trans_jadwal_pertemuan.sampai_jam_old").
		Where("trans_jadwal_pertemuan.status_usulan = ? AND mst_kelas.id_semester = ?", "disetujui", smtId).
		Order("id ASC")
}
