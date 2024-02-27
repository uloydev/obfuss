package models

type GetAllJadwalKuliahResponse struct {
	KodeKurikulum   string `json:"kode_kurikulum" gorm:"column:kode_kurikulum"`
	IdKelas         int    `json:"id_kelas" gorm:"column:id_kelas"`
	JumlahMK        string `json:"jumlah_mk" gorm:"column:jumlah_mk"`
	InisialKelas    string `json:"inisial_kelas" gorm:"column:inisial_kelas"`
	NamaKelas       string `json:"nama_kelas" gorm:"column:nama_kelas"`
	Semester        int    `json:"semester" gorm:"column:semester"`
	Tahun           string `json:"tahun" gorm:"column:tahun"`
	NamaProdi       string `json:"nama_prodi" gorm:"nama_prodi"`
	NamaKonsentrasi string `json:"nama_konsentrasi" gorm:"nama_konsentrasi"`
}

type JadwalKuliahCreateRequest struct {
	IDMk         int    `json:"id_mk" binding:"required"`
	IDDosen      int    `json:"id_dosen" binding:"required"`
	Hari         string `json:"hari" binding:"required"`
	IDJamMulai   int    `json:"id_jam_mulai" binding:"required"`
	IDJamSelesai int    `json:"id_jam_selesai" binding:"required"`
}

type JadwalKuliahRequest struct {
	IdKelas      int                         `json:"id_kelas" binding:"required"`
	JadwalKuliah []JadwalKuliahCreateRequest `json:"jadwal_kuliah" binding:"required"`
}
