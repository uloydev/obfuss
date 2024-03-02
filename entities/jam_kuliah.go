package entities

import "time"

type JamKuliah struct {
	BaseModel
	ID         uint      `json:"id"`
	NoUrut     int       `json:"no_urut"`
	JamMulai   time.Time `json:"jam_mulai"`
	JamSelesai time.Time `json:"jam_selesai"`
	Keterangan string    `json:"keterangan"`
}

func (e JamKuliah) TableName() string {
	return "mst_jam_kuliah"
}
