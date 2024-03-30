package entities

import "time"

type Dosen struct {
	BaseModel
	ID                 uint      `json:"id"`
	NomorDosen         string    `json:"nomor_dosen"`
	NamaLengkap        string    `json:"nama_lengkap"`
	NIP                string    `json:"nip"`
	NIDN               string    `json:"nidn"`
	NUPN               string    `json:"nupn"`
	Instansi           string    `json:"instansi"`
	Kelamin            string    `json:"kelamin"`
	Agama              string    `json:"agama"`
	TempatLahir        string    `json:"tempat_lahir"`
	TanggalLahir       time.Time `json:"tanggal_lahir"`
	NIK                string    `json:"nik"`
	PendidikanTerakhir string    `json:"pendidikan_terakhir"`
	DusunKampung       string    `json:"dusun_kampung"`
	Jalan              string    `json:"jalan"`
	RT                 string    `json:"rt"`
	RW                 string    `json:"rw"`
	DesaKelurahan      string    `json:"desa_kelurahan"`
	Kecamatan          string    `json:"kecamatan"`
	KotaKab            string    `json:"kota_kab"`
	Provinsi           string    `json:"provinsi"`
	KodePos            string    `json:"kode_pos"`
	TelpHp             string    `json:"telp_hp"`
	Email              string    `json:"email"`
	RiwayatPendidikan  string    `json:"riwayat_pendidikan"`
	DaftarKaryaTulis   string    `json:"daftar_karya_tulis"`
	DaftarPenelitian   string    `json:"daftar_penelitian"`
	DaftarSeminar      string    `json:"daftar_seminar"`
	DaftarOrganisasi   string    `json:"daftar_organisasi"`
	DaftarPPM          string    `json:"daftar_ppm"`
	DaftarSertifikasi  string    `json:"daftar_sertifikasi"`
	DaftarPrestasi     string    `json:"daftar_prestasi"`
	Keterangan         string    `json:"keterangan"`
	Status             string    `json:"status"`
	Uname              string    `json:"uname"`
}

func (e Dosen) TableName() string {
	return "mst_dosen"
}
