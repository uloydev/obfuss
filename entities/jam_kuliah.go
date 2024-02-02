package entities

import "time"

type JamKuliah struct {
	BaseModel
	ID         uint
	NoUrut     int
	JamMulai   time.Time
	JamSelesai time.Time
	Keterangan string
}

func (j *JamKuliah) TableName() string {
	return "mst_jam_kuliah"
}
