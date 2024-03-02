package entities

import (
	"time"
)

// TODO: Add belongs to kelas, plotkelas, BelongsToMany rewardPunish, and has Many laporanKhm

type Mahasiswa struct {
	BaseModel
	ID                  uint      `json:"id"`
	UpdatedAt           time.Time `json:"updated_at"`
	UserName            string    `json:"user_name"`
	Password            string    `json:"password"`
	Email               string    `json:"email"`
	Email2              string    `json:"email2"`
	RememberToken       string    `json:"remember_token"`
	ProfilePicture      string    `json:"profile_picture"`
	Validated           int       `json:"validated"`
	NamaLengkap         string    `json:"nama_lengkap"`
	NPM                 string    `json:"npm" gorm:"unique_index"`
	TahunMasuk          int       `json:"tahun_masuk"`
	Kelamin             string    `json:"kelamin"`
	Agama               string    `json:"agama"`
	NIK                 string    `json:"nik"`
	TempatLahir         string    `json:"tempat_lahir"`
	TanggalLahir        time.Time `json:"tanggal_lahir"`
	TlpHP               string    `json:"tlp_hp"`
	NISN                string    `json:"nisn"`
	NPWP                string    `json:"npwp"`
	GolDarah            string    `json:"gol_darah"`
	PenerimaKps         string    `json:"penerima_kps"`
	NomorKps            string    `json:"nomor_kps"`
	AsalSma             string    `json:"asal_sma"`
	NamaDusun           string    `json:"nama_dusun"`
	Jalan               string    `json:"jalan"`
	Rt                  int       `json:"rt"`
	Rw                  int       `json:"rw"`
	DesaKelurahan       string    `json:"desa_kelurahan"`
	Alamat              string    `json:"alamat"`
	Kecamatan           string    `json:"kecamatan"`
	KodePos             string    `json:"kode_pos" gorm:"type:char(5);"`
	KotaKab             string    `json:"kota_kab" gorm:"type:varchar(30);"`
	Provinsi            string    `json:"provinsi" gorm:"type:varchar(30);"`
	NamaAyah            string    `json:"nama_ayah" gorm:"type:varchar(50);"`
	NIKAyah             string    `json:"nik_ayah" gorm:"type:varchar(100);"`
	TempatLahirAyah     string    `json:"tempat_lahir_ayah" gorm:"type:varchar(30);"`
	TanggalLahirAyah    time.Time `json:"tanggal_lahir_ayah" gorm:"type:date;"`
	PendidikanAyah      string    `json:"pendidikan_ayah" gorm:"type:varchar(10);"`
	PekerjaanAyah       string    `json:"pekerjaan_ayah" gorm:"type:varchar(30);"`
	PenghasilanAyah     float64   `json:"penghasilan_ayah" gorm:"type:decimal(18,2);"`
	KebutuhanKhususAyah string    `json:"kebutuhan_khusus_ayah" gorm:"type:varchar(100);"`
	NamaIbu             string    `json:"nama_ibu" gorm:"type:varchar(50);"`
	NIKIbu              string    `json:"nik_ibu" gorm:"type:varchar(100);"`
	TempatLahirIbu      string    `json:"tempat_lahir_ibu" gorm:"type:varchar(30);"`
	TanggalLahirIbu     time.Time `json:"tanggal_lahir_ibu" gorm:"type:date;"`
	PendidikanIbu       string    `json:"pendidikan_ibu" gorm:"type:varchar(10);"`
	PekerjaanIbu        string    `json:"pekerjaan_ibu" gorm:"type:varchar(30);"`
	PenghasilanIbu      float64   `json:"penghasilan_ibu" gorm:"type:decimal(18,2);"`
	KebutuhanKhususIbu  string    `json:"kebutuhan_khusus_ibu" gorm:"type:varchar(100);"`
	AlamatOrtu          string    `json:"alamat_ortu" gorm:"type:varchar(191);"`
	KontakOrtu          string    `json:"kontak_ortu" gorm:"type:varchar(12);"`
	NamaWali            string    `json:"nama_wali" gorm:"type:varchar(50);"`
	NIKWali             string    `json:"nik_wali" gorm:"type:varchar(100);"`
	AlamatWali          string    `json:"alamat_wali" gorm:"type:varchar(191);"`
	KontakWali          string    `json:"kontak_wali" gorm:"type:varchar(12);"`
	TempatLahirWali     string    `json:"tempat_lahir_wali" gorm:"type:varchar(30);"`
	TanggalLahirWali    time.Time `json:"tanggal_lahir_wali" gorm:"type:date;"`
	PendidikanWali      int       `json:"pendidikan_wali" gorm:"type:tinyint(4);"`
	PekerjaanWali       string    `json:"pekerjaan_wali" gorm:"type:varchar(30);"`
	PenghasilanWali     float64   `json:"penghasilan_wali" gorm:"type:decimal(18,2);"`
	IDKelas             uint      `json:"id_kelas" gorm:"index"`
	StatPD              int       `json:"stat_pd"`
	StatDO              int       `json:"stat_do"`
	AlasanDO            string    `json:"alasan_do"`
	KepdirDO            string    `json:"kepdir_do"`
	SemesterDO          string    `json:"semester_do"`
	Confirmed           int       `json:"confirmed"`
}

func (e Mahasiswa) TableName() string {
	return "mst_mahasiswa"
}
