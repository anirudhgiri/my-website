package schema

import (
	"gorm.io/gorm"
)

type Blog struct {
	gorm.Model
	Title    string
	SubTitle string
	Date     string
	Content  string
	Slug     string
}
