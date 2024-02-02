package entities

import "time"

type Kurikulum struct {
	BaseModel
	ID             uint
	KodeKurikulum  string
	MulaiKurikulum time.Time
	SksWajib       uint
	SksPilihan     uint
	ActiveStatus   string `gorm:"type:enum('yes','no');default:'no';"`
}

func (k *Kurikulum) TableName() string {
	return "mst_kurikulum"
}
