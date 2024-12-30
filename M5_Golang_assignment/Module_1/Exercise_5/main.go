package main

import (
	"errors"
	"fmt"
)

// City struct to store climate data
type City struct {
	Name          string
	Temperature   float64 
	Rainfall      float64 
}

var climateData []City 

// AddCity adds a new city to the climate data
func AddCity(name string, temperature, rainfall float64) {
	climateData = append(climateData, City{
		Name:        name,
		Temperature: temperature,
		Rainfall:    rainfall,
	})
}

// MaxMinTemp finds the cities with the highest and lowest temperatures
func MaxMinTemp() (City, City, error) {
	if len(climateData) == 0 {
		return City{}, City{}, errors.New("no climate data available")
	}

	highest := climateData[0]
	lowest := climateData[0]

	for _, city := range climateData {
		if city.Temperature > highest.Temperature {
			highest = city
		}
		if city.Temperature < lowest.Temperature {
			lowest = city
		}
	}

	return highest, lowest, nil
}

// CalculateAverageRainfall calculates the average rainfall across all cities
func CalculateAverageRainfall() (float64, error) {
	if len(climateData) == 0 {
		return 0, errors.New("no climate data available")
	}

	totalRainfall := 0.0
	for _, city := range climateData {
		totalRainfall += city.Rainfall
	}

	return totalRainfall / float64(len(climateData)), nil
}

// FilterCitiesByRainfall returns cities with rainfall above a given threshold
func FilterCitiesByRainfall(threshold float64) []City {
	var filtered []City
	for _, city := range climateData {
		if city.Rainfall > threshold {
			filtered = append(filtered, city)
		}
	}
	return filtered
}

// SearchCityByName searches for a city by name
func SearchCityByName(name string) (*City, error) {
	for _, city := range climateData {
		if city.Name == name {
			return &city, nil
		}
	}
	return nil, errors.New("city not found")
}

func main() {
	// Add initial data
	AddCity("New York", 16.5, 1200)
	AddCity("Mumbai", 27.0, 2500)
	AddCity("London", 10.0, 800)
	AddCity("Sydney", 20.5, 900)

	// Find highest and lowest temperature cities
	highest, lowest, err := MaxMinTemp()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("City with the highest temperature:", highest)
		fmt.Println("City with the lowest temperature:", lowest)
	}

	// Calculate average rainfall
	averageRainfall, err := CalculateAverageRainfall()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Average Rainfall: %.2f mm\n", averageRainfall)
	}

	// Filter cities by rainfall threshold
	threshold := 1000.0
	fmt.Printf("\nCities with rainfall above %.2f mm:\n", threshold)
	for _, city := range FilterCitiesByRainfall(threshold) {
		fmt.Println(city)
	}

	// Search for a city by name
	cityName := "Mumbai"
	city, err := SearchCityByName(cityName)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("\nData for city %s: %+v\n", cityName, *city)
	}
}
