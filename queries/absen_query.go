package queries

import (
	"database/sql"

	"gorm.io/gorm"
)

func FindAllAbsenQuery(
	db *gorm.DB,
	userType string,
	smtId, kelasId, DosenId int,
) *gorm.DB {
	var (
		filterConds string
		filterArgs  []sql.NamedArg
	)
	if userType == "dosen" {
		filterConds += `and trans_jadwal_kuliah.id_dosen= @userId`
		filterArgs = append(filterArgs, sql.Named("userId", DosenId))
	} else if userType == "mahasiswa" {
		filterConds += `and mst_kelas.id= @kelasId`
		filterArgs = append(filterArgs, sql.Named("kelasId", kelasId))
	}

	query := `SELECT  trans_jadwal_pertemuan.id AS id_pertemuan, trans_jadwal_kuliah.id AS id, SUBSTRING(jam_kuliah_mulai.jam_mulai,1,5) AS jam_mulai, SUBSTRING(jam_kuliah_mulai.jam_selesai,1,5) AS jam_selesai, trans_jadwal_kuliah.hari, 
        mst_dosen.nama_lengkap AS nama_dosen, trans_jadwal_kuliah.id_kelas, mst_kurikulum.kode_kurikulum, trans_jadwal_kuliah.id_mk, 
        mst_mata_kuliah.nama_mk, trans_jadwal_kuliah.id_semester, mst_kelas.inisial_kelas, mst_kelas.nama_kelas, mst_semester.semester,
        mst_tahun_ajaran.tahun, mst_konsentrasi.nama_prodi, mst_konsentrasi.nama_konsentrasi,
        COUNT(trans_jadwal_pertemuan.id) AS jumlah_pertemuan, terlaksana.count_terlaksana AS jumlah_hadir,
        case when trans_jadwal_kuliah.hari = 'Senin' then 1 
                when trans_jadwal_kuliah.hari = 'Selasa' then 2 
                when trans_jadwal_kuliah.hari = 'Rabu' then 3
                when trans_jadwal_kuliah.hari = 'Kamis' then 4 
                when trans_jadwal_kuliah.hari = 'Jumat' then 5 
                when trans_jadwal_kuliah.hari = 'Sabtu' then 6 else 7 end AS "day_number"
        FROM trans_jadwal_pertemuan
        JOIN trans_jadwal_kuliah ON trans_jadwal_pertemuan.id_jadwal = trans_jadwal_kuliah.id
        LEFT JOIN (
            SELECT trans_jadwal_pertemuan.id AS id_pertemuan, trans_jadwal_kuliah.id AS id_jadwal, COUNT(trans_jadwal_pertemuan.id) AS count_terlaksana
            FROM trans_jadwal_pertemuan
            JOIN trans_jadwal_kuliah ON trans_jadwal_pertemuan.id_jadwal = trans_jadwal_kuliah.id
            WHERE trans_jadwal_pertemuan.is_terlaksana = 'yes'
            GROUP BY trans_jadwal_pertemuan.id_jadwal
        ) terlaksana ON terlaksana.id_jadwal = trans_jadwal_pertemuan.id_jadwal
        JOIN mst_jam_kuliah jam_kuliah_mulai ON trans_jadwal_kuliah.id_jam_mulai = jam_kuliah_mulai.id
        JOIN mst_jam_kuliah jam_kuliah_selesai ON trans_jadwal_kuliah.id_jam_selesai = jam_kuliah_selesai.id
        JOIN mst_dosen ON trans_jadwal_kuliah.id_dosen = mst_dosen.id
        JOIN mst_kelas ON trans_jadwal_kuliah.id_kelas = mst_kelas.id
        JOIN mst_kurikulum ON mst_kelas.id_kurikulum = mst_kurikulum.id
        JOIN mst_konsentrasi ON mst_kelas.id_kurikulum = mst_konsentrasi.id
        JOIN mst_mata_kuliah ON trans_jadwal_kuliah.id_mk = mst_mata_kuliah.id
        JOIN mst_semester ON trans_jadwal_kuliah.id_semester = mst_semester.id
        JOIN mst_tahun_ajaran ON mst_semester.id_tahun = mst_tahun_ajaran.id 
        WHERE  trans_jadwal_kuliah.id_semester = @smtId` +
		filterConds +
		`GROUP BY trans_jadwal_kuliah.id
        ORDER BY mst_kelas.id , day_number;`

	filterArgs = append(filterArgs, sql.Named("smtId", smtId))

	return db.Raw(query, filterArgs)
}

func CountAbsenQuery(db *gorm.DB, idPertemuan int) *gorm.DB {
	query := `(SELECT COUNT(1) FROM absen_mahasiswa WHERE id_pertemuan = ? AND is_hadir = 'yes') as jumlah_hadir, (SELECT COUNT(1) FROM absen_mahasiswa WHERE id_pertemuan = ? AND is_hadir = 'no') as jumlah_tidak_hadir, as (SELECT COUNT(1) FROM absen_mahasiswa WHERE id_pertemuan = ? AND keterangan = 'luring') jumlah_luring, (SELECT COUNT(1) FROM absen_mahasiswa WHERE id_pertemuan = ? AND keterangan = 'daring') jumlah_daring`
	return db.Select(query, idPertemuan, idPertemuan, idPertemuan, idPertemuan)
}
