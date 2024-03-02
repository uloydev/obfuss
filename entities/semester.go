package entities

import (
	"time"
)

type Semester struct {
	BaseModel
	ID              uint        `json:"id"`
	IDTahun         int         `json:"id_tahun"`
	TahunAjaran     TahunAjaran `json:"tahun_ajaran" gorm:"foreignKey:IDTahun"`
	Semester        string      `json:"semester"`
	TanggalMulai    time.Time   `json:"tanggal_mulai"`
	TanggalBerakhir time.Time   `json:"tanggal_berakhir"`
	ActiveStatus    string      `json:"active_status"`
}

func (e Semester) TableName() string {
	return "mst_semester"
}
