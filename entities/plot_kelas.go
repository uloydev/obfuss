package entities

type PlotKelas struct {
	BaseModel
	ID          uint      `json:"id"`
	IDSemester  int       `json:"id_semester"`
	Semester    Semester  `json:"semester" gorm:"foreignKey:IDSemester"`
	IDKelas     int       `json:"id_kelas"`
	Kelas       Kelas     `json:"kelas" gorm:"foreignKey:IDKelas"`
	IDMahasiswa int       `json:"id_mahasiswa"`
	Mahasiswa   Mahasiswa `json:"mahasiswa" gorm:"foreignKey:IDMahasiswa"`
	TahunMasuk  string    `json:"tahun_masuk"`
}

func (e PlotKelas) TableName() string {
	return "trans_plot_kelas"
}
