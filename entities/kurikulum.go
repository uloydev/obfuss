package entities

import "time"

type Kurikulum struct {
	BaseModel
	ID             uint      `json:"id"`
	KodeKurikulum  string    `json:"kode_kurikulum"`
	MulaiKurikulum time.Time `json:"mulai_kurikulum"`
	SksWajib       uint      `json:"sks_wajib"`
	SksPilihan     uint      `json:"sks_pilihan"`
	ActiveStatus   string    `json:"active_status" gorm:"type:enum('yes','no');default:'no';"`
}

func (e Kurikulum) TableName() string {
	return "mst_kurikulum"
}
