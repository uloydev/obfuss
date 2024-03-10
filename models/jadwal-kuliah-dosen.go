package models

import "skripsi.id/obfuss/entities"

type SaveTransJadwalKuliahDosen struct {
	ID           int    `json:"id"`
	IDSemester   int    `json:"id_semester"`
	IDKelas      int    `json:"id_kelas"`
	IDMK         int    `json:"id_mk"`
	IDDosen      int    `json:"id_dosen"`
	Hari         string `json:"hari"`
	IDJamMulai   int    `json:"id_jam_mulai"`
	IDJamSelesai int    `json:"id_jam_selesai"`
}

type GetMahasiswaRaw struct {
	IDKelas     int `json:"id_kelas"`
	IDPertemuan int `json:"id_pertemuan"`
}

type GetGridJadwalKuliahDosenResult struct {
	entities.Kurikulum
}
