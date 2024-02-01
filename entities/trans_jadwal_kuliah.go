package entities

import (
	"time"

	"github.com/jinzhu/gorm"
)

type TransJadwalKuliah struct {
	gorm.Model
	ID           int       `gorm:"type:int(11);auto_increment;"`
	IDSemester   int       `gorm:"type:int(11);"`
	IDKelas      int       `gorm:"type:int(11);"`
	IDMk         int       `gorm:"type:int(11);"`
	IDDosen      int       `gorm:"type:int(11);"`
	Hari         string    `gorm:"type:enum('Senin','Selasa','Rabu','Kamis','Jumat','Sabtu','Minggu');"`
	IDJamMulai   int       `gorm:"type:int(11);"`
	IDJamSelesai int       `gorm:"type:int(11);"`
	Status       string    `gorm:"type:enum('active','non_active');"`
	AddDate      time.Time `gorm:"type:timestamp;default:current_timestamp();"`
	ModifiedDate time.Time `gorm:"type:datetime;"`
	AddUser      int       `gorm:"type:int(11);"`
	ModifiedUser int       `gorm:"type:int(11);"`
}

func (s *TransJadwalKuliah) TableName() string {
	return "trans_jadwal_kuliah"
}
