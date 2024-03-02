package entities

type Konsentrasi struct {
	BaseModel
	ID              uint   `json:"id"`
	KodeProdi       int    `json:"kode_prodi"`
	NamaProdi       string `json:"nama_prodi"`
	NamaKonsentrasi string `json:"nama_konsentrasi"`
	PejabatTte      uint64 `json:"pejabat_tte"`
	Jurusan         string `json:"jurusan"`
	Gelar           string `json:"gelar"`
	InisialGelar    string `json:"inisial_gelar"`
	TranslateGelar  string `json:"translate_gelar"`
	Keterangan      string `json:"keterangan"`
}

func (e Konsentrasi) TableName() string {
	return "mst_konsentrasi"
}
