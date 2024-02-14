package handler

import (
	"log"
	"machine-coding-food-delivery-golang/internal/food_delivery/model"
	"machine-coding-food-delivery-golang/internal/food_delivery/service"
)

type RestaurantHandler struct {
	restaurantService service.RestaurantService
}

type UserHandler struct {
	userService       service.UserService
	restaurantService service.RestaurantService
}

func GetRestaurantHandler() *RestaurantHandler {
	return &RestaurantHandler{
		restaurantService: service.GetRestaurantService(),
	}
}

func GetUserHandler() *UserHandler {
	return &UserHandler{
		userService:       service.GetUserService(),
		restaurantService: service.GetRestaurantService(),
	}
}

func (r *RestaurantHandler) RegisterRestaurant(restaurant *model.Restaurant) error {
	ok, err := r.restaurantService.RegisterRestaurant(restaurant)
	if !ok || err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (r *RestaurantHandler) UpdateQuantity(restaurantName string, quantity int) error {
	ok, err := r.restaurantService.UpdateQuantity(restaurantName, quantity)
	if !ok || err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (r *UserHandler) ShowRestaurants(sortBy model.SortBy, pinCode string) []*model.Restaurant {
	return r.restaurantService.ShowOrders(sortBy, pinCode)
}

func (u *UserHandler) RegisterUser(user *model.User) error {
	ok, err := u.userService.RegisterUser(user)
	if !ok || err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (u *UserHandler) LoginUser(phoneNumber string) *model.User {
	user, err := u.userService.LoginUser(phoneNumber)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return user
}
