package entities

import "time"

type TahunAjaran struct {
	BaseModel
	ID            int `gorm:"type:int(11);auto_increment;"`
	Tahun         string
	MulaiBerjalan time.Time
	ActiveStatus  string
}

func (t TahunAjaran) TableName() string {
	return "mst_tahun_ajaran"
}
