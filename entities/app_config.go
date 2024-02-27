package entities

type AppConfig struct {
	ID             int `json:"id"`
	IDTahun        int `json:"id_tahun"`
	IDSemester     int `json:"id_semester"`
	TotalMk        int `json:"total_mk"`
	TotalPertemuan int `json:"total_pertemuan"`
}

func (e AppConfig) TableName() string {
	return "app_config"
}
