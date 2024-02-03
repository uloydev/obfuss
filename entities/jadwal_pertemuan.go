package entities

import (
	"time"
)

type JadwalPertemuan struct {
	BaseModel
	ID                 int          `gorm:"type:int(11);auto_increment;"`
	IDJadwal           int          `gorm:"type:int(11);"`
	JadwalKuliah       JadwalKuliah `gorm:"foreignKey:IDJadwal"`
	PertemuanKe        int          `gorm:"type:int(3);"`
	TanggalPertemuan   time.Time    `gorm:"type:date;"`
	MulaiJam           int          `gorm:"type:int(11);"`
	SampaiJam          int          `gorm:"type:int(11);"`
	TanggalUsulanGanti time.Time    `gorm:"type:date;"`
	UsulanMulaiJam     int          `gorm:"type:int(11);"`
	UsulanSampaiJam    int          `gorm:"type:int(11);"`
	TanggalOld         time.Time    `gorm:"type:date;"`
	MulaiJamOld        int          `gorm:"type:int(11);"`
	SampaiJamOld       int          `gorm:"type:int(11);"`
	Keterangan         string       `gorm:"type:varchar(255);"`
	AlasanPerubahan    string       `gorm:"type:text;"`
	StatusUsulan       string       `gorm:"type:varchar(30);"`
	IsTerlaksana       string       `gorm:"type:enum('yes','no');default:'no';"`
}

func (j JadwalPertemuan) TableName() string {
	return "trans_jadwal_pertemuan"
}
