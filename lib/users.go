package lib
import (
	"simple-users-mvc/configs"
	"simple-users-mvc/models"
)

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := configs.DB.Find(&users).Error
	if err != nil {
		return []models.User{}, err
	}
	return users, nil
}

func PostUsers(user models.User) (*models.User, error) {
	if err := configs.DB.Save(&user).Error; err != nil {
		return &models.User{}, err
	}
	return &user, nil
}

func GetUserByID(ID int) (*models.User, error) {
	var user models.User
	err := configs.DB.Where("id = ?", ID).First(&user).Error
	if err != nil {
		return &models.User{}, err
	}
	return &user, nil
}

func UpdateUser(user models.User, ID int) (*models.User, error) {
	err := configs.DB.Where("id = ?", ID).Updates(user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func DeleteUser(ID int) (*models.User, error) {
	var user models.User
	err := configs.DB.Where("id = ?", ID).Delete(&user).Error
	if err != nil {
		return &models.User{}, err
	}
	return &user, nil
}