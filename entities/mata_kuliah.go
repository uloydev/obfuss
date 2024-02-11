package entities

type MataKuliah struct {
	BaseModel
	ID             uint   `json:"id"`
	IDKonsentrasi  int    `json:"id_konsentrasi"`
	IDKurikulum    int    `json:"id_kurikulum"`
	KodeMk         string `json:"kode_mk"`
	NamaMk         string `json:"nama_mk"`
	Keterangan     string `json:"keterangan"`
	SemesterKuliah int    `json:"semester_kuliah"`
	SKS            int    `json:"sks"`
	Status         string `json:"status"`
	IsAgama        int    `json:"is_agama"`
	IsTaPkl        string `json:"is_ta_pkl"`
	JenisMk        string `json:"jenis_mk"`
	SksPraktek     uint   `json:"sks_praktek"`
	SksLapangan    uint   `json:"sks_lapangan"`
	SksSimulasi    uint   `json:"sks_simulasi"`
}

func (m MataKuliah) TableName() string {
	return "mst_mata_kuliah"
}
