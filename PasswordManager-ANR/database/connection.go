package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

/*Database connection*/
//var DSN = "host=localhost user=admin password=pss123 dbname=passworddb port=5432"
//Variables globales
var DSN = "host=localhost user=postgres password=pss123 dbname=postgres port=5433"
var DB *gorm.DB //Objeto que representa la conexión a la bbdd

// Funcion que conecta con la base de datos usando la librería gorm
func DBConnection() {
	var error error
	DB, error = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if error != nil {
		log.Fatal(error)
	} else {
		log.Println("DB connected")
	}
}
