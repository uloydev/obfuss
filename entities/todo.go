package entities

import (
	"time"
)

type Todo struct {
	ID        uint      // Standard field for the primary key
	Action    string    // Action is the todo action
	CreatedAt time.Time // Automatically managed by GORM for creation time
	UpdatedAt time.Time // Automatically managed by GORM for update time
}
