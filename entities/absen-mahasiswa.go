package entities

import (
	"time"
)

type AbsenMahasiswa struct {
	BaseModel
	ID              int `gorm:"type:int(11);auto_increment;"`
	IDPertemuan     int
	JadwalPertemuan JadwalPertemuan `gorm:"foreignKey:IDPertemuan"`
	IDMahasiswa     int
	Mahasiswa       Mahasiswa `gorm:"foreignKey:IDMahasiswa"`
	IsHadir         string    `gorm:"type:enum('yes','no');"`
	AbsenDate       time.Time `gorm:"type:datetime;"`
	Keterangan      string    `gorm:"type:enum('daring','luring');"`
}

func (t *AbsenMahasiswa) TableName() string {
	return "trans_absen_mahasiswa"
}

func (t *AbsenMahasiswa) TableIndex() [][]string {
	return [][]string{
		{"ID", "IDPertemuan", "IDMahasiswa"},
	}
}
