package models

import (
	"gorm.io/gorm"
)

type Membro struct {
	gorm.Model
	Name       string
	Nascimento string
	Foto       string
	Meme       string
	Desc       string
}
