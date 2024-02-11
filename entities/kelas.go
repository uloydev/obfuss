package entities

type Kelas struct {
	BaseModel
	ID             uint
	IDSemester     int
	Semester       Semester `gorm:"foreignKey:IDSemester"`
	IDKurikulum    int
	Kurikulum      Kurikulum `gorm:"foreignKey:IDKurikulum"`
	IDKonsentrasi  int
	Konsentrasi    Konsentrasi `gorm:"foreignKey:IDKonsentrasi"`
	NamaKelas      string
	InisialKelas   string
	SemesterKuliah string
	Keterangan     string
}

func (k Kelas) TableName() string {
	return "mst_kelas"
}
