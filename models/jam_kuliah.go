package models

type GetAllJamKuliah struct {
	ID         uint    `json:"id"`
	NoUrut     *int    `json:"no_urut"`
	JamMulai   *string `json:"jam_mulai"`
	JamSelesai *string `json:"jam_selesai"`
	Keterangan *string `json:"keterangan"`
}
