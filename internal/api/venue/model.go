package venue

import "gorm.io/gorm"

type Venue struct {
	gorm.Model
	ID   int64
	Name string
}
