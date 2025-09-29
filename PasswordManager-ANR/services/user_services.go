package services

import (
	"errors"

	"github.com/anarrod172/passwordmanager/database"
	"github.com/anarrod172/passwordmanager/models"
)

/*func getUsers() ([]models.User, error) {
	var users []models.User
	result := database.DB.Find(&users)
	return users, result.Error
}*/

/*
Función que busca en la bbdd un usuario en función de que su campo url
coincida con el valor proporcionado
*/
func GetUserByUrl(url string) (models.User, error) {
	var user models.User
	//Filtro por la url y si no encuentra nada devolvemos un error
	result := database.DB.Select("username", "url", "notes").Where("url = ?", url).Find(&user)
	if result.Error != nil {
		return models.User{}, result.Error
	}

	if result.RowsAffected == 0 {
		return models.User{}, errors.New("user not found")
	}
	//Devolvemos usuario y nil como error
	return user, nil
}

/*Función que registra un usuario en la bbdd*/
func CreateUser(user models.User) error {
	return database.DB.Create(&user).Error
}

/*
Función en la que eliminamos el usuario en función
de su nombre de usuario
*/
func DeleteUser(username string) error {
	result := database.DB.Where("username =?", username).Delete(&models.User{})
	if result.Error != nil {
		// Error de la base de datos
		return result.Error
	}

	if result.RowsAffected == 0 {
		// No se encontró ningún usuario con ese username
		return errors.New("user not found")
	}

	return nil

}

/*
Función que actualiza el campo notas de un usuario dependiendo del id
del usuario
*/
func UpdateUser(id uint, user models.User) error {
	var existingUser models.User
	//Buscamos el usuario por id
	if err := database.DB.First(&existingUser, id).Error; err != nil {
		return err
	}
	// Actualizamos campos
	existingUser.Notes = user.Notes
	// Guardamos el usuario
	if err := database.DB.Save(&existingUser).Error; err != nil {
		return err
	}
	return nil

}
