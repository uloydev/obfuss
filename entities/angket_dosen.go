package entities

import "time"

type AngketDosen struct {
	ID               int        `gorm:"column:id;primaryKey" json:"id"`
	IDDosen          int        `gorm:"column:id_dosen" json:"id_dosen"`
	IDPertemuan      int        `gorm:"column:id_pertemuan" json:"id_pertemuan"`
	PertemuanKe      *int       `gorm:"column:pertemuan_ke" json:"pertemuan_ke"`
	JamKuliah        *string    `gorm:"column:jam_kuliah" json:"jam_kuliah"`
	JumlahHadir      *int       `gorm:"column:jumlah_hadir" json:"jumlah_hadir"`
	JumlahTidakHadir *int       `gorm:"column:jumlah_tidak_hadir" json:"jumlah_tidak_hadir"`
	JumlahLuring     *int       `gorm:"column:jumlah_luring" json:"jumlah_luring"`
	JumlahDaring     *int       `gorm:"column:jumlah_daring" json:"jumlah_daring"`
	RingkasanMateri  *string    `gorm:"column:ringkasan_materi" json:"ringkasan_materi"`
	Keterangan       *string    `gorm:"column:keterangan" json:"keterangan"`
	File             *string    `gorm:"column:file" json:"file"`
	FileGambar       *string    `gorm:"column:file_gambar" json:"file_gambar"`
	Status           string     `gorm:"column:status; type:enum('Draft','Updated','Saved','Signed'); default:'Draft'" json:"status"`
	AddDate          *time.Time `gorm:"column:add_date" json:"add_date"`
	ModifiedDate     *time.Time `gorm:"column:modified_date;" json:"modified_date"`
	AddUser          *int       `gorm:"column:add_user" json:"add_user"`
	ModifiedUser     *int       `gorm:"column:modified_user" json:"modified_user"`
}

func (AngketDosen) TableName() string {
	return "trans_angket_dosen"
}
