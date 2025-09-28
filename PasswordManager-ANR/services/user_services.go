package services

import (
	"github.com/anarrod172/passwordmanager/database"
	"github.com/anarrod172/passwordmanager/models"
)

/*func getUsers() ([]models.User, error) {
	var users []models.User
	result := database.DB.Find(&users)
	return users, result.Error
}*/

func GetUserByUrl(url string) (models.User, error) {
	var user models.User
	//Filtro por la url y si no encuentra nada devolvemos un error
	if err := database.DB.Select("username", "url", "notes").Where("url = ?", url).Find(&user).Error; err != nil {
		return models.User{}, err
	}

	//Devolvemos usuario y nil como error
	return user, nil
}

func CreateUser(user models.User) error {
	return database.DB.Create(&user).Error
}

// Eliminamos el usuario en función de su nombre de usuario
func DeleteUser(username string) error {
	return database.DB.Where("username =?", username).Delete(&models.User{}).Error
}

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
