package performer

import "gorm.io/gorm"

type Performer struct {
	gorm.Model
	ID   int64
	Name string
}
