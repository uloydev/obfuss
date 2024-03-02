package entities

type JadwalKuliahDosen struct {
	BaseModel
	ID           int       `gorm:"type:int(11);auto_increment;" json:"id"`
	IDSemester   *int      `gorm:"type:int(11);" json:"id_semester"`
	Semester     *Semester `gorm:"foreignKey:IDSemester" json:"semester"`
	IDKelas      *int      `gorm:"type:int(11);" json:"id_kelas"`
	Kelas        *Kelas    `gorm:"foreignKey:IDKelas" json:"kelas"`
	IDMk         *int      `gorm:"type:int(11);" json:"id_mk"`
	IDDosen      *int      `gorm:"type:int(11);" json:"id_dosen"`
	Hari         *string   `gorm:"type:enum('Senin','Selasa','Rabu','Kamis','Jumat','Sabtu','Minggu');" json:"hari"`
	IDJamMulai   *int      `gorm:"type:int(11);" json:"id_jam_mulai"`
	IDJamSelesai *int      `gorm:"type:int(11);" json:"id_jam_selesai"`
	Status       *string   `gorm:"type:enum('active','non_active');"`
}

func (e JadwalKuliahDosen) TableName() string {
	return "trans_jadwal_kuliah"
}
