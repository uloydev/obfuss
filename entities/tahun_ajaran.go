package entities

import "time"

type TahunAjaran struct {
	BaseModel
	ID            int       `json:"id"`
	Tahun         string    `json:"tahun"`
	MulaiBerjalan time.Time `json:"mulai_berjalan"`
	ActiveStatus  string    `json:"active_status"`
}

func (e TahunAjaran) TableName() string {
	return "mst_tahun_ajaran"
}
