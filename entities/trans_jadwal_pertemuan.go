package entities

import (
	"time"

	"gorm.io/gorm"
)

type TransJadwalPertemuan struct {
	gorm.Model
	ID                 int       `gorm:"type:int(11);auto_increment;"`
	IDJadwal           int       `gorm:"type:int(11);"`
	PertemuanKe        int       `gorm:"type:int(3);"`
	TanggalPertemuan   time.Time `gorm:"type:date;"`
	MulaiJam           int       `gorm:"type:int(11);"`
	SampaiJam          int       `gorm:"type:int(11);"`
	TanggalUsulanGanti time.Time `gorm:"type:date;"`
	UsulanMulaiJam     int       `gorm:"type:int(11);"`
	UsulanSampaiJam    int       `gorm:"type:int(11);"`
	TanggalOld         time.Time `gorm:"type:date;"`
	MulaiJamOld        int       `gorm:"type:int(11);"`
	SampaiJamOld       int       `gorm:"type:int(11);"`
	Keterangan         string    `gorm:"type:varchar(255);"`
	AlasanPerubahan    string    `gorm:"type:text;"`
	StatusUsulan       string    `gorm:"type:varchar(30);"`
	IsTerlaksana       string    `gorm:"type:enum('yes','no');default:'no';"`
	AddDate            time.Time `gorm:"type:timestamp;default:current_timestamp();"`
	ModifiedDate       time.Time `gorm:"type:datetime;"`
	AddUser            int       `gorm:"type:int(11);"`
	ModifiedUser       int       `gorm:"type:int(11);"`
}

func (s *TransJadwalPertemuan) TableName() string {
	return "trans_jadwal_pertemuan"
}
