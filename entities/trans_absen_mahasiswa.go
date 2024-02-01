package entities

import (
	"time"

	"gorm.io/gorm"
)

type TransAbsenMahasiswa struct {
	gorm.Model
	ID              int `gorm:"type:int(11);auto_increment;"`
	IDPertemuan     int
	JadwalPertemuan TransJadwalPertemuan
	IDMahasiswa     int
	Mahasiswa       Mahasiswa
	IsHadir         string    `gorm:"type:enum('yes','no');"`
	AbsenDate       time.Time `gorm:"type:datetime;"`
	Keterangan      string    `gorm:"type:enum('daring','luring');"`
	AddDate         time.Time `gorm:"type:timestamp;default:current_timestamp();"`
	ModifiedDate    time.Time `gorm:"type:datetime;"`
	AddUser         int       `gorm:"type:int(11);"`
	ModifiedUser    int       `gorm:"type:int(11);"`
}

func (t *TransAbsenMahasiswa) TableName() string {
	return "trans_absen_mahasiswa"
}

func (t *TransAbsenMahasiswa) TableIndex() [][]string {
	return [][]string{
		{"ID", "IDPertemuan", "IDMahasiswa"},
	}
}
