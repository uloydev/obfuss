package entities

// i dont know where come from IDJamMulai and IDJamSelesai
type JadwalKuliah struct {
	BaseModel
	ID           int      `gorm:"type:int(11);auto_increment;"`
	IDSemester   int      `gorm:"type:int(11);"`
	Semester     Semester `gorm:"foreignKey:IDSemester"`
	IDKelas      int      `gorm:"type:int(11);"`
	Kelas        Kelas    `gorm:"foreignKey:IDKelas"`
	IDMk         int      `gorm:"type:int(11);"`
	IDDosen      int      `gorm:"type:int(11);"`
	Hari         string   `gorm:"type:enum('Senin','Selasa','Rabu','Kamis','Jumat','Sabtu','Minggu');"`
	IDJamMulai   int      `gorm:"type:int(11);"`
	IDJamSelesai int      `gorm:"type:int(11);"`
	Status       string   `gorm:"type:enum('active','non_active');"`
}

func (s *JadwalKuliah) TableName() string {
	return "trans_jadwal_kuliah"
}
