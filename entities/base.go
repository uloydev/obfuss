package entities

import "time"

type BaseModel struct {
	AddDate      time.Time `json:"add_date"`
	ModifiedDate time.Time `json:"modified_date"`
	AddUser      int       `json:"add_user"`
	ModifiedUser int       `json:"modified_user"`
}
