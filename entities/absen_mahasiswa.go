package entities

import (
	"time"
)

type AbsenMahasiswa struct {
	BaseModel
	ID              int             `json:"id"`
	IDPertemuan     int             `json:"id_pertemuan"`
	JadwalPertemuan JadwalPertemuan `json:"jadwal_pertemuan" gorm:"foreignKey:IDPertemuan"`
	IDMahasiswa     int             `json:"id_mahasiswa"`
	Mahasiswa       Mahasiswa       `json:"mahasiswa" gorm:"foreignKey:IDMahasiswa"`
	IsHadir         string          `json:"is_hadir" gorm:"type:enum('yes','no');"`
	AbsenDate       time.Time       `json:"absen_date" gorm:"type:datetime;"`
	Keterangan      string          `json:"keterangan" gorm:"type:enum('daring','luring');"`
}

func (a AbsenMahasiswa) TableName() string {
	return "trans_absen_mahasiswa"
}

func (a AbsenMahasiswa) TableIndex() [][]string {
	return [][]string{
		{"ID", "IDPertemuan", "IDMahasiswa"},
	}
}
