package internal

import (
	"log"
)

//go:generate mockery --name Service
type Service interface {
	GetByUserNameAndPasword(username, password string)(User, error)
	GetAllUsers() []User
	GetUserById(id uint64) (User, error)
	CreateUser(user User) (User, error)
	UpdateUserById(id uint64, updatedUser User) (User, error)
	DeleteUserById(id uint64) error
}

type userService struct {
	repository Repository
}

func (us userService) GetByUserNameAndPasword(username, password string)(User, error) {
	user, err := us.repository.FindUserNameAndPassword(username, password)
	if err != nil {
		log.Printf("Error while fetching user by username and password: %s\n", err)
	}
	log.Printf("Fetched user by username and password: %s\n", username)
	return user, err
}

func (us userService) GetAllUsers() []User {
	return us.repository.GetAllUsers()
}

func (us userService) GetUserById(id uint64) (User, error) {
	user, err := us.repository.GetUserById(id)
	if err != nil {
		log.Printf("Error %d id doesn't exist. %s", id, err)
	}
	return user, err
}

func (us userService) CreateUser(user User) (User, error) {
	user, err := us.repository.CreateUser(user)
	if err != nil {
		log.Printf("An error occur while creating the user: %s\n", err)
	}
	return user, err
}

func (us userService) UpdateUserById(id uint64, updatedUser User) (User, error) {
	user, err := us.repository.UpdateUserById(id, updatedUser)
	if err != nil {
		log.Printf("Error while updating user: %s\n", err)
	}
	return user, err
}

func (us userService) DeleteUserById(id uint64) error {
	err := us.repository.DeleteUserById(id)
	if err != nil {
		log.Println("Something went wrong could not delete task", err)
	}
	return err
}

func NewUserService(repository Repository) Service {
	return &userService{repository}
}
