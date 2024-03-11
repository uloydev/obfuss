package models

import "mime/multipart"

type SaveTransLaporanPerkuliahan struct {
	IDPertemuan     int                   `form:"id_pertemuan" binding:"required"`
	IDDosen         int                   `form:"id_dosen" binding:"required"`
	PertemuanKe     int                   `form:"pertemuan_ke" binding:"required"`
	JamKuliah       string                `form:"jam_kuliah" binding:"required"`
	RingkasanMateri string                `form:"materi" binding:"required"`
	Keterangan      string                `form:"keterangan" binding:"required"`
	Foto            *multipart.FileHeader `form:"foto" swaggerignore:"true"`
	FileName        string
}
