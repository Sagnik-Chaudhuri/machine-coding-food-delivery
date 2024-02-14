package service

import (
	"log"
	"machine-coding-food-delivery-golang/internal/food_delivery/model"
	"sort"
)

type SortProviderService interface {
	Sort(restaurant []*model.Restaurant) []*model.Restaurant
}

type PricingSortProvider struct {
}

type RatingSortProvider struct {
}

func GetSortProvider(sortBy model.SortBy) SortProviderService {
	log.Println("initialising GetSortProvider")
	switch sortBy {
	case model.SORTBYPRICING:
		return &PricingSortProvider{}
	case model.SORTBYRATING:
		return &RatingSortProvider{}
	}
	return nil
}

func (sp *PricingSortProvider) Sort(restaurants []*model.Restaurant) []*model.Restaurant {
	sortedRestaurants := restaurants[:]

	sort.Slice(sortedRestaurants, func(i, j int) bool {
		return restaurants[i].Price < restaurants[j].Price
	})
	return sortedRestaurants
}

func (rp *RatingSortProvider) Sort(restaurants []*model.Restaurant) []*model.Restaurant {
	sortedRestaurants := restaurants[:]

	sort.Slice(sortedRestaurants[:], func(i, j int) bool {
		return restaurants[i].Rating > restaurants[j].Rating
	})
	return sortedRestaurants
}
