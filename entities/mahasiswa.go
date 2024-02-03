package entities

import (
	"time"
)

// TODO: Add belongs to kelas, plotkelas, BelongsToMany rewardPunish, and has Many laporanKhm

type Mahasiswa struct {
	BaseModel
	ID                  uint
	UpdatedAt           time.Time
	UserName            string
	Password            string
	Email               string
	Email2              string
	RememberToken       string
	ProfilePicture      string
	Validated           int
	NamaLengkap         string
	NPM                 string `gorm:"unique_index"`
	TahunMasuk          int
	Kelamin             string
	Agama               string
	NIK                 string
	TempatLahir         string
	TanggalLahir        time.Time
	TlpHP               string
	NISN                string
	NPWP                string
	GolDarah            string
	PenerimaKps         string
	NomorKps            string
	AsalSma             string
	NamaDusun           string
	Jalan               string
	Rt                  int
	Rw                  int
	DesaKelurahan       string
	Alamat              string
	Kecamatan           string
	KodePos             string    `gorm:"type:char(5);"`
	KotaKab             string    `gorm:"type:varchar(30);"`
	Provinsi            string    `gorm:"type:varchar(30);"`
	NamaAyah            string    `gorm:"type:varchar(50);"`
	NIKAyah             string    `gorm:"type:varchar(100);"`
	TempatLahirAyah     string    `gorm:"type:varchar(30);"`
	TanggalLahirAyah    time.Time `gorm:"type:date;"`
	PendidikanAyah      string    `gorm:"type:varchar(10);"`
	PekerjaanAyah       string    `gorm:"type:varchar(30);"`
	PenghasilanAyah     float64   `gorm:"type:decimal(18,2);"`
	KebutuhanKhususAyah string    `gorm:"type:varchar(100);"`
	NamaIbu             string    `gorm:"type:varchar(50);"`
	NIKIbu              string    `gorm:"type:varchar(100);"`
	TempatLahirIbu      string    `gorm:"type:varchar(30);"`
	TanggalLahirIbu     time.Time `gorm:"type:date;"`
	PendidikanIbu       string    `gorm:"type:varchar(10);"`
	PekerjaanIbu        string    `gorm:"type:varchar(30);"`
	PenghasilanIbu      float64   `gorm:"type:decimal(18,2);"`
	KebutuhanKhususIbu  string    `gorm:"type:varchar(100);"`
	AlamatOrtu          string    `gorm:"type:varchar(191);"`
	KontakOrtu          string    `gorm:"type:varchar(12);"`
	NamaWali            string    `gorm:"type:varchar(50);"`
	NIKWali             string    `gorm:"type:varchar(100);"`
	AlamatWali          string    `gorm:"type:varchar(191);"`
	KontakWali          string    `gorm:"type:varchar(12);"`
	TempatLahirWali     string    `gorm:"type:varchar(30);"`
	TanggalLahirWali    time.Time `gorm:"type:date;"`
	PendidikanWali      int       `gorm:"type:tinyint(4);"`
	PekerjaanWali       string    `gorm:"type:varchar(30);"`
	PenghasilanWali     float64   `gorm:"type:decimal(18,2);"`
	IDKelas             uint      `gorm:"index"`
	StatPD              int
	StatDO              int
	AlasanDO            string
	KepdirDO            string
	SemesterDO          string
	Confirmed           int
}
