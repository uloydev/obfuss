package entities

import "time"

type BaseModel struct {
	AddDate      time.Time
	ModifiedDate time.Time
	AddUser      int
	ModifiedUser int
}
