package models

import "gorm.io/gorm"

//Creamos el tipo de dato icono
type Icon struct {
	Path   string `gorm:"not null"`
	Width  int
	Height int
}

//Creamos el modelo usuario
type User struct {
	//
	gorm.Model
	Username string `gorm:"not null;unique index"`
	Password string `gorm:"unique;not null;unique_index"`
	Url      string
	Notes    string `gorm:"size:250"`
	Icon     Icon   `gorm:"embedded"` //Los campos de icon se guardan en la tabla user
}
