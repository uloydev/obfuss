package entities

type MataKuliah struct {
	BaseModel
	ID             uint
	IDKonsentrasi  int
	IDKurikulum    int
	KodeMk         string
	NamaMk         string
	Keterangan     string
	SemesterKuliah int
	SKS            int
	Status         string
	IsAgama        int
	IsTaPkl        string
	JenisMk        string
	SksPraktek     uint
	SksLapangan    uint
	SksSimulasi    uint
}

func (m MataKuliah) TableName() string {
	return "mst_mata_kuliah"
}
