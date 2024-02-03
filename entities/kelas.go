package entities

type Kelas struct {
	BaseModel
	ID             uint
	IDSemester     int
	Semester       Semester `gorm:"foreignKey:IDJadwal"`
	IDKurikulum    int
	Kurikulum      Kurikulum `gorm:"foreignKey:IDKurikulum"`
	IDKonsentrasi  int
	Konsentrasi    Konsentrasi `gorm:"foreignKey:IDKurikulum"`
	NamaKelas      string
	InisialKelas   string
	SemesterKuliah string
	Keterangan     string
}

func (k Kelas) TableName() string {
	return "mst_kelas"
}
