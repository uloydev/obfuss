package entities

import "time"

type Dosen struct {
	BaseModel
	ID                 uint
	NomorDosen         string
	NIP                string
	NIDN               string
	NUPN               string
	Instansi           string
	Kelamin            string
	Agama              string
	TempatLahir        string
	TanggalLahir       time.Time
	NIK                string
	PendidikanTerakhir string
	DusunKampung       string
	Jalan              string
	RT                 string
	RW                 string
	DesaKelurahan      string
	Kecamatan          string
	KotaKab            string
	Provinsi           string
	KodePos            string
	TelpHp             string
	Email              string
	RiwayatPendidikan  string
	DaftarKaryaTulis   string
	DaftarPenelitian   string
	DaftarSeminar      string
	DaftarOrganisasi   string
	DaftarPPM          string
	DaftarSertifikasi  string
	DaftarPrestasi     string
	Keterangan         string
	Status             string
	Uname              string
}

func (d Dosen) TableName() string {
	return "mst_dosen"
}
