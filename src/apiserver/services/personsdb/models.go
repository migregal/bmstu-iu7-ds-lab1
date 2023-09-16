package personsdb

import "gorm.io/gorm"

type Person struct {
	gorm.Model

	ID      int32 `gorm:"primaryKey"`
	Name    string
	Age     int32
	Address string
	Work    string
}
