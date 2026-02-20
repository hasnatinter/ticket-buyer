package ticket

import (
	"time"

	"gorm.io/gorm"
)

type Ticket struct {
	gorm.Model
	ID        int64
	Seat      string
	UserID    int64
	Status    string
	EventId   int64
	CreatedAt *time.Time `gorm:"autoCreateTime"`
	UpdatedAt *time.Time `gorm:"autoCreateTime"`
	DeletedAt *time.Time
}
