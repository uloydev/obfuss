package entities

type Konsentrasi struct {
	BaseModel
	ID              uint
	KodeProdi       int
	NamaProdi       string
	NamaKonsentrasi string
	PejabatTte      uint64
	Jurusan         string
	Gelar           string
	InisialGelar    string
	TranslateGelar  string
	Keterangan      string
}

func (k Konsentrasi) TableName() string {
	return "mst_konsentrasi"
}
