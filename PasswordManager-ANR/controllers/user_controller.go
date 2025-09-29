package controllers

import (
	"net/http"
	"strconv"

	"github.com/anarrod172/passwordmanager/models"
	"github.com/anarrod172/passwordmanager/services"
	"github.com/gin-gonic/gin"
)

/*
Servicio que devuelve un estatus ok si encuentra un usuario con la url indicada
en caso de no coincidir, devuelve un mensaje de usuario no encontrado
*/
func GetUser(c *gin.Context) {
	url := c.Query("url") //obtiene el parámetro url de la query
	user, err := services.GetUserByUrl(url)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, user)
}

/*
Servicio que devuelve un código de respuesta 201 Created si el nuevo usuario se ha creado
si no es así, muestra un mensaje de error que devuelve un código de respuesta 400 bad request
*/
func PostUser(c *gin.Context) {
	var newUser models.User //Declaramos un usuario
	//El usuario se llena con los datos del body Json
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//Llamada al servicio para crear el usuario
	if err := services.CreateUser(newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error to create a new user": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, newUser)
}

/*
Servicio que devuelve un código de respuesta http 200 con estatus ok si el usuario se ha eliminado
si no se elimina muestra el mensaje usuario no eliminado
*/
func DeleteUser(c *gin.Context) {
	username := c.Param("username")
	err := services.DeleteUser(username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "user not deleted"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "user deleted successfully"})
}

/*
Servicio que devuelve un estatus de respuesta ok si se puede actualizar el campo notas del usuario
si no se puede, muestra un mensaje de error
*/
func PutNotes(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var user models.User
	// Validamos JSON de entrada
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Llamamos al servicio para actualizar
	if err := services.UpdateUser(uint(id), user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "user successfully updated"})

}
