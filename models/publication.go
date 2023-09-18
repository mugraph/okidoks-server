// models/publication.go

package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Publication struct {
	gorm.Model
	Name				string
	Attributes	datatypes.JSON
}