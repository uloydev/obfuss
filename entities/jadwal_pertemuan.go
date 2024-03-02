package entities

import (
	"time"
)

type JadwalPertemuan struct {
	BaseModel
	ID                 int          `json:"id" gorm:"type:int(11);auto_increment;"`
	IDJadwal           int          `json:"id_jadwal" gorm:"type:int(11);"`
	JadwalKuliah       JadwalKuliah `json:"jadwal_kuliah" gorm:"foreignKey:IDJadwal"`
	PertemuanKe        int          `json:"pertemuan_ke" gorm:"type:int(3);"`
	TanggalPertemuan   time.Time    `json:"tanggal_pertemuan" gorm:"type:date;"`
	MulaiJam           int          `json:"mulai_jam" gorm:"type:int(11);"`
	SampaiJam          int          `json:"sampai_jam" gorm:"type:int(11);"`
	TanggalUsulanGanti time.Time    `json:"tanggal_usulan_ganti" gorm:"type:date;"`
	UsulanMulaiJam     int          `json:"usulan_mulai_jam" gorm:"type:int(11);"`
	UsulanSampaiJam    int          `json:"usulan_sampai_jam" gorm:"type:int(11);"`
	TanggalOld         time.Time    `json:"tanggal_old" gorm:"type:date;"`
	MulaiJamOld        int          `json:"mulai_jam_old" gorm:"type:int(11);"`
	SampaiJamOld       int          `json:"sampai_jam_old" gorm:"type:int(11);"`
	Keterangan         string       `json:"keterangan" gorm:"type:varchar(255);"`
	AlasanPerubahan    string       `json:"alasan_perubahan" gorm:"type:text;"`
	StatusUsulan       string       `json:"status_usulan" gorm:"type:varchar(30);"`
	IsTerlaksana       string       `json:"is_terlaksana" gorm:"type:enum('yes','no');default:'no';"`
}

func (e JadwalPertemuan) TableName() string {
	return "trans_jadwal_pertemuan"
}
