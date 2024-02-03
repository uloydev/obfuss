package entities

import (
	"time"
)

type Semester struct {
	BaseModel
	ID              uint
	IDTahun         int
	TahunAjaran     TahunAjaran `gorm:"foreignKey:IDTahun"`
	Semester        string
	TanggalMulai    time.Time
	TanggalBerakhir time.Time
	ActiveStatus    string
}

func (s Semester) TableName() string {
	return "mst_semester"
}
