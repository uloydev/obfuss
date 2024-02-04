package queries

import (
	"database/sql"

	"gorm.io/gorm"
)

func FindAllJadwalKuliah(
	db *gorm.DB,
	smtId uint,
) *gorm.DB {
	var (
		filterArgs []sql.NamedArg
	)
	query := `
	SELECT mst_kurikulum.kode_kurikulum, trans_jadwal_kuliah.id_kelas, COUNT(trans_jadwal_kuliah.id_mk) AS jumlah_mk, mst_kelas.inisial_kelas, mst_kelas.nama_kelas, mst_semester.semester, mst_tahun_ajaran.tahun, mst_konsentrasi.nama_prodi, mst_konsentrasi.nama_konsentrasi FROM trans_jadwal_kuliah 
		JOIN mst_kelas ON mst_kelas.id=trans_jadwal_kuliah.id_kelas
		JOIN mst_kurikulum ON mst_kurikulum.id=mst_kelas.id_kurikulum
		JOIN mst_konsentrasi ON mst_konsentrasi.id=mst_kelas.id_konsentrasi
		JOIN mst_semester ON mst_semester.id=mst_kelas.id_semester
		JOIN mst_tahun_ajaran ON mst_tahun_ajaran.id=mst_semester.id_tahun
	WHERE mst_semester.id = @id_semester
	GROUP BY mst_kurikulum.kode_kurikulum, trans_jadwal_kuliah.id_kelas,mst_kelas.inisial_kelas, mst_kelas.nama_kelas, mst_semester.semester, mst_tahun_ajaran.tahun,
	mst_konsentrasi.nama_prodi, mst_konsentrasi.nama_konsentrasi`

	filterArgs = append(filterArgs, sql.Named("id_semester", smtId))

	return db.Raw(query, filterArgs)
}
