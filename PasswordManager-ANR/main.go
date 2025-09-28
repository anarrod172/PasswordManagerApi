package main

import (
	"github.com/anarrod172/passwordmanager/database"
	"github.com/anarrod172/passwordmanager/models"
	"github.com/anarrod172/passwordmanager/routes"
)

func main() {

	//Ejecutamos la función DBConnection
	database.DBConnection()

	//Migraciones automáticas para crear las tablas de la bbdd
	database.DB.AutoMigrate(&models.User{})

	//Llamada a los servicios
	r := routes.SetupRoutes()
	r.Run(":8080")
}
