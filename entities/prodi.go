package entities

type Prodi struct {
	BaseModel
	ID         uint   `json:"id"`
	Kode       string `json:"kode"`
	NamaProdi  string `json:"nama_prodi"`
	Keterangan string `json:"keterangan"`
}

func (e Prodi) TableName() string {
	return "mst_program_studi"
}
