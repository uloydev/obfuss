package models

type GetAllPerubahanJadwal struct {
	ID                 int    `gorm:"column:id" json:"id"`
	TanggalPertemuan   string `gorm:"column:tanggal_pertemuan" json:"tanggal_pertemuan"`
	TanggalUsulanGanti string `gorm:"column:tanggal_usulan_ganti" json:"tanggal_usulan_ganti"`
	Keterangan         string `gorm:"column:keterangan" json:"keterangan"`
	AlasanPerubahan    string `gorm:"column:alasan_perubahan" json:"alasan_perubahan"`
	StatusUsulan       string `gorm:"column:status_usulan" json:"status_usulan"`
	TanggalOld         string `gorm:"column:tanggal_old" json:"tanggal_old"`
	MulaiJam           int    `gorm:"column:mulai_jam" json:"mulai_jam"`
	SampaiJam          int    `gorm:"column:sampai_jam" json:"sampai_jam"`
	MulaiJamOld        int    `gorm:"column:mulai_jam_old" json:"mulai_jam_old"`
	SampaiJamOld       int    `gorm:"column:sampai_jam_old" json:"sampai_jam_old"`
	InisialKelas       string `gorm:"column:inisial_kelas" json:"inisial_kelas"`
	JamKuliah          string `gorm:"column:jam_kuliah" json:"jam_kuliah"`
	JamKuliahOld       string `gorm:"column:jam_kuliah_old" json:"jam_kuliah_old"`
	KodeKurikulum      string `gorm:"column:kode_kurikulum" json:"kode_kurikulum"`
	NamaDosen          string `gorm:"column:nama_dosen" json:"nama_dosen"`
	NamaKelas          string `gorm:"column:nama_kelas" json:"nama_kelas"`
	NamaKonsentrasi    string `gorm:"column:nama_konsentrasi" json:"nama_konsentrasi"`
	NamaMk             string `gorm:"column:nama_mk" json:"nama_mk"`
	NamaProdi          string `gorm:"column:nama_prodi" json:"nama_prodi"`
	Semester           string `gorm:"column:semester" json:"semester"`
	Tahun              string `gorm:"column:tahun" json:"tahun"`
	HariJadwal         string `gorm:"column:hari_jadwal" json:"hari_jadwal"`
	HariOld            string `gorm:"column:hari_old" json:"hari_old"`
}

type TransJadwalKuliah struct {
	ID      int `gorm:"column:id"`
	IDKelas int `gorm:"column:id_kelas"`
	IDMk    int `gorm:"column:id_mk"`
	IDDosen int `gorm:"column:id_dosen"`
}

type MstKelas struct {
	ID            int `gorm:"column:id"`
	IDKurikulum   int `gorm:"column:id_kurikulum"`
	IDKonsentrasi int `gorm:"column:id_konsentrasi"`
	IDSemester    int `gorm:"column:id_semester"`
}

type MstJamKuliah struct {
	ID         int    `gorm:"column:id"`
	JamMulai   string `gorm:"column:jam_mulai"`
	JamSelesai string `gorm:"column:jam_selesai"`
}

type UpdateJadwalPertemuanRequest struct {
	IDJadwal           int    `json:"id_jadwal" example:"1"`
	TanggalUsulanGanti string `json:"tanggal_usulan_ganti" example:"2006-01-02 15:04:05"`
	UsulanMulaiJam     int    `json:"usulan_mulai_jam" example:"1"`
	UsulanSampaiJam    int    `json:"usulan_sampai_jam" example:"1"`
	AlasanPerubahan    string `json:"alasan_perubahan"`
	StatusUsulan       string `json:"status_usulan" example:"disetujui"`
}
