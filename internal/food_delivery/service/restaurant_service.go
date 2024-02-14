package service

import (
	"errors"
	"log"
	"machine-coding-food-delivery-golang/internal/food_delivery/model"
	"sync"
)

type RestaurantService interface {
	RegisterRestaurant(restaurant *model.Restaurant) (bool, error)
	UpdateQuantity(restaurantName string, quantity int) (bool, error)
	ShowOrders(sortBy model.SortBy, pinCode string) []*model.Restaurant
}

type RestaurantServiceMemoryImpl struct {
	Restaurants     map[string]*model.Restaurant
	PinToRestaurant map[string][]*model.Restaurant
}

var restaurantServiceInstance RestaurantService
var restaurantServiceOnce sync.Once

func GetRestaurantService() RestaurantService {
	log.Println("initialising RestaurantService")
	restaurantServiceOnce.Do(func() {
		restaurantServiceInstance = &RestaurantServiceMemoryImpl{
			Restaurants:     map[string]*model.Restaurant{},
			PinToRestaurant: map[string][]*model.Restaurant{},
		}
	})

	return restaurantServiceInstance
}

func (r *RestaurantServiceMemoryImpl) RegisterRestaurant(restaurant *model.Restaurant) (bool, error) {
	_, ok := r.Restaurants[restaurant.Name]
	if ok {
		return false, errors.New("restaurant already registered with the same name")
	}
	r.Restaurants[restaurant.Name] = restaurant
	for _, pinCode := range restaurant.PinCodes {
		r.PinToRestaurant[pinCode] = append(r.PinToRestaurant[pinCode], restaurant)
	}
	return true, nil
}

func (r *RestaurantServiceMemoryImpl) UpdateQuantity(restaurantName string, quantity int) (bool, error) {
	val, ok := r.Restaurants[restaurantName]
	if !ok {
		return false, errors.New("restaurant is not registered")
	}
	val.Quantity = quantity
	return false, nil
}

func (r *RestaurantServiceMemoryImpl) ShowOrders(sortBy model.SortBy, pinCode string) []*model.Restaurant {
	restaurants := r.PinToRestaurant[pinCode]

	sortProvider := GetSortProvider(sortBy)
	return sortProvider.Sort(restaurants)

}
