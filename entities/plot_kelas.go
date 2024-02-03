package entities

type PlotKelas struct {
	BaseModel
	ID          uint
	IDSemester  int
	Semester    Semester `gorm:"foreignKey:IDSemester"`
	IDKelas     int
	Kelas       Kelas `gorm:"foreignKey:IDKelas`
	IDMahasiswa int
	Mahasiswa   Mahasiswa `gorm:"foreignKey:IDMahasiswa`
	TahunMasuk  string
}

func (p *PlotKelas) TableName() string {
	return "trans_plot_kelas"
}
