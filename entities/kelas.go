package entities

type Kelas struct {
	BaseModel
	ID             uint
	IDSemester     int
	IDKurikulum    int
	Kurikulum      Kurikulum `gorm:"foreignKey:IDKurikulum"`
	IDKonsentrasi  int
	NamaKelas      string
	InisialKelas   string
	SemesterKuliah string
	Keterangan     string
}

func (k *Kelas) TableName() string {
	return "mst_kelas"
}
