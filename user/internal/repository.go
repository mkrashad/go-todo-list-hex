package internal

import (
	"gorm.io/gorm"
)

//go:generate mockery --name Repository
type Repository interface {
	FindUserNameAndPassword(username, password string)(User, error)
	GetAllUsers() []User
	GetUserById(id uint64) (User, error)
	CreateUser(user User) (User, error)
	UpdateUserById(id uint64, updatedUser User) (User, error)
	DeleteUserById(id uint64) error
}

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) Repository {
	return &userRepository{DB: db}
}

func (tr userRepository) FindUserNameAndPassword(username, password string) (User, error){
	var user User 
	result := tr.DB.Where("username =? AND password =?", username, password).First(&user)
	return user, result.Error
}

func (tr userRepository) GetAllUsers() []User {
	var users []User
	tr.DB.Preload("Tasks").Find(&users)
	return users
}

func (tr userRepository) GetUserById(id uint64) (User, error) {
	var user User
	result := tr.DB.Preload("Tasks").Find(&user, id)
	return user, result.Error
}

func (tr userRepository) CreateUser(user User) (User, error) {
	result := tr.DB.Create(&user)
	tr.DB.Save(&user)
	return user, result.Error
}

func (tr userRepository) UpdateUserById(id uint64, updatedUser User) (User, error) {
	var user User
	tr.DB.Find(&user, id)
	result := tr.DB.Model(&user).Updates(updatedUser)
	return user, result.Error
}

func (tr userRepository) DeleteUserById(id uint64) error {
	result := tr.DB.Delete(&User{}, id)
	return result.Error
}


