package main

import (
	"fmt"
	"log"
	"machine-coding-food-delivery-golang/internal/food_delivery/handler"
	"machine-coding-food-delivery-golang/internal/food_delivery/model"
)

func main() {
	fmt.Print("inside main")

	restaurantHandler := handler.GetRestaurantHandler()
	userHandler := handler.GetUserHandler()
	var err error

	rx1 := &model.Restaurant{
		Name:     "rx 1",
		ItemName: "item 1",
		PinCodes: []string{"10", "20", "30"},
		Quantity: 0,
		Price:    300,
	}
	rx2 := &model.Restaurant{
		Name:     "rx 2",
		ItemName: "item 2",
		PinCodes: []string{"10", "20"},
		Quantity: 0,
		Price:    200,
	}

	err = restaurantHandler.RegisterRestaurant(rx1)
	if err != nil {
		log.Println("handler RegisterRestaurant error", err)
	}
	err = restaurantHandler.RegisterRestaurant(rx2)
	if err != nil {
		log.Println("handler RegisterRestaurant error", err)
	}

	//user3 := userHandler.LoginUser("123")
	//if user3 == nil {
	//	log.Println("login error", err)
	//}

	user1 := &model.User{
		Name:        "u1",
		PhoneNumber: "123",
		Gender:      "M",
		PinCode:     "20",
	}
	user2 := &model.User{
		Name:        "u2",
		PhoneNumber: "234",
		Gender:      "M",
		PinCode:     "30",
	}

	err = userHandler.RegisterUser(user1)
	if err != nil {
		log.Println("handler creation error", err)
	}
	err = userHandler.RegisterUser(user2)
	if err != nil {
		log.Println("handler creation error", err)
	}

	loggedInUser1 := userHandler.LoginUser("123")
	if loggedInUser1 == nil {
		log.Println("handler login error")
	}
	loggedInUser2 := userHandler.LoginUser("234")
	if loggedInUser2 == nil {
		log.Println("handler login error")
	} else {
		log.Println("user details: ", loggedInUser2)
	}

	sortedRestaurants := userHandler.ShowRestaurants(model.SORTBYPRICING, "10")
	for _, restaurant := range sortedRestaurants {
		log.Println("restaurant details: ", restaurant)
	}

}
