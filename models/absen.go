package models

type GetAllAbsenResponse struct {
	IdPertemuan     int    `json:"id_pertemuan" gorm:"column:id_pertemuan"`
	Id              int    `json:"id" gorm:"column:id"`
	JamMulai        string `json:"jam_mulai" gorm:"column:jam_mulai"`
	JamSelesai      string `json:"jam_selesai" gorm:"column:jam_selesai"`
	Hari            string `json:"hari" gorm:"column:hari"`
	NamaLengkap     string `json:"nama_lengkap" gorm:"column:nama_lengkap"`
	NamaDosen       string `json:"nama_dosen" gorm:"column:nama_dosen"`
	KodeKurikulum   string `json:"kode_kurikulum" gorm:"column:kode_kurikulum"`
	IdMk            int    `json:"id_mk" gorm:"column:id_mk"`
	NamaMk          string `json:"nama_mk" gorm:"column:nama_mk"`
	IdSemester      int    `json:"id_semester" gorm:"column:id_semester"`
	InisialKelas    string `json:"inisial_kelas" gorm:"column:inisial_kelas"`
	NamaKelas       string `json:"nama_kelas" gorm:"column:nama_kelas"`
	Semester        string `json:"semester" gorm:"column:semester"`
	Tahun           string `json:"tahun" gorm:"column:tahun"`
	NamaProdi       string `json:"nama_prodi" gorm:"column:nama_prodi"`
	NamaKonsentrasi string `json:"nama_konsentrasi" gorm:"column:nama_konsentrasi"`
	JumlahPertemuan int    `json:"jumlah_pertemuan" gorm:"column:jumlah_pertemuan"`
	JumlahHadir     int    `json:"jumlah_hadir" gorm:"column:jumlah_hadir"`
	DayNumber       int    `json:"day_number" gorm:"column:day_number"`
	IdKelas         int    `json:"id_kelas" gorm:"column:id_kelas"`
}

type SaveAbsenTransRequest struct {
	IdJadwal    int              `json:"id_jadwal" example:"1" binding:"required"`
	IdPertemuan int              `json:"id_pertemuan" example:"1" binding:"required"`
	IdKelas     int              `json:"id_kelas" example:"1" binding:"required"`
	Absen       []SaveTransAbsen `json:"absen" binding:"required"`
}

type SaveTransAbsen struct {
	IdMahasiswa int    `json:"id_mahasiswa" example:"1" binding:"required"`
	IsHadir     string `json:"is_hadir" example:"yes"`      // yes or no
	Keterangan  string `json:"keterangan" example:"daring"` // daring or luring
}

type AbsenCountResult struct {
	JumlahHadir      int `json:"jumlah_hadir" gorm:"column:jumlah_hadir"`
	JumlahTidakHadir int `json:"jumlah_tidak_hadir" gorm:"column:jumlah_tidak_hadir"`
	JumlahLuring     int `json:"jumlah_luring" gorm:"column:jumlah_luring"`
	JumlahDaring     int `json:"jumlah_daring" gorm:"column:jumlah_daring"`
}

type Result struct {
	Kehadiran   *string `gorm:"column:kehadiran"`
	Keterangan  *string `gorm:"column:keterangan"`
	UserName    string  `gorm:"column:user_name"`
	NamaLengkap string  `gorm:"column:nama_lengkap"`
}
