package service

import (
	"errors"
	"log"
	"machine-coding-food-delivery-golang/internal/food_delivery/model"
	"sync"
)

type UserService interface {
	RegisterUser(user *model.User) (bool, error)
	LoginUser(phoneNumber string) (*model.User, error)
}

type UserServiceMemoryImpl struct {
	Users map[string]*model.User
}

var userServiceInstance UserService
var userServiceOnce sync.Once

func GetUserService() UserService {
	log.Println("initialising UserService")
	userServiceOnce.Do(func() {
		userServiceInstance = &UserServiceMemoryImpl{
			Users: map[string]*model.User{},
		}
	})

	return userServiceInstance
}

func (u *UserServiceMemoryImpl) RegisterUser(user *model.User) (bool, error) {
	_, ok := u.Users[user.PhoneNumber]
	if ok {
		return false, errors.New("user already registered with the same phone number")
	}
	u.Users[user.PhoneNumber] = user
	return true, nil
}

func (u *UserServiceMemoryImpl) LoginUser(phoneNumber string) (*model.User, error) {
	val, ok := u.Users[phoneNumber]
	if !ok {
		return nil, errors.New("user not registered")
	}
	return val, nil
}
