package models

type SaveTransJadwalKuliahDosen struct {
	ID           int    `json:"id"`
	IDKelas      int    `json:"id_kelas"`
	IDMK         int    `json:"id_mk"`
	IDDosen      int    `json:"id_dosen"`
	Hari         string `json:"hari"`
	IDJamMulai   int    `json:"id_jam_mulai"`
	IDJamSelesai int    `json:"id_jam_selesai"`
}
