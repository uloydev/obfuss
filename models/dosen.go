package models

type GetAllDosen struct {
	ID          uint    `json:"id"`
	NamaLengkap *string `json:"nama_lengkap"`
	NIP         *string `json:"nip"`
	NIDN        *string `json:"nidn"`
	NUPN        *string `json:"nupn"`
	Kelamin     *string `json:"kelamin"`
	NIK         *string `json:"nik"`
	Keterangan  *string `json:"keterangan"`
	Status      *string `json:"status"`
	Uname       *string `json:"uname"`
}
