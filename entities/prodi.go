package entities

type Prodi struct {
	BaseModel
	ID         uint
	Kode       string
	NamaProdi  string
	Keterangan string
}

func (p Prodi) TableName() string {
	return "mst_program_studi"
}
