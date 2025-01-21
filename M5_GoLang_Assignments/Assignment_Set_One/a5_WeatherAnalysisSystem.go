package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type City struct {
	Name        string
	Temperature float64
	Rainfall    float64
}

func findHighestTemp(cities []City) *City {
	highestTemp := cities[0]
	for _, city := range cities {
		if city.Temperature > highestTemp.Temperature {
			highestTemp = city
		}
	}
	return &highestTemp
}

func findLowestTemp(cities []City) *City {
	lowestTemp := cities[0]
	for _, city := range cities {
		if city.Temperature < lowestTemp.Temperature {
			lowestTemp = city
		}
	}
	return &lowestTemp
}

func calculateAvgRainfall(cities []City) float64 {
	totalRainfall := 0.0
	for _, city := range cities {
		totalRainfall += city.Rainfall
	}
	return totalRainfall / float64(len(cities))
}

func filterByRainfall(cities []City, threshold float64) []City {
	var filteredCities []City
	for _, city := range cities {
		if city.Rainfall > threshold {
			filteredCities = append(filteredCities, city)
		}
	}
	return filteredCities
}

func searchCityByName(cities []City, name string) *City {
	for _, city := range cities {
		if strings.EqualFold(city.Name, name) {
			return &city
		}
	}
	return nil
}

func main() {
	cities := []City{
		{Name: "Mumbai", Temperature: 32.5, Rainfall: 100.5},
		{Name: "Delhi", Temperature: 30.5, Rainfall: 50.5},
		{Name: "Chennai", Temperature: 35.5, Rainfall: 150.5},
		{Name: "Kolkata", Temperature: 28.5, Rainfall: 80.5},
		{Name: "Bangalore", Temperature: 29.5, Rainfall: 120.5},
	}

	reader := bufio.NewReader(os.Stdin)

	for {

		fmt.Println("Menu: ")
		fmt.Println("1. Highest Temperature City")
		fmt.Println("2. Lowest Temperature City")
		fmt.Println("3. Average Rainfall")
		fmt.Println("4. Filter by Rainfall")
		fmt.Println("5. Search City by Name")
		fmt.Println("6. Exit")
		fmt.Print("Enter your choice: ")
		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			highestTempCity := findHighestTemp(cities)
			fmt.Printf("City with highest temperature: %s %.2f\n", highestTempCity.Name, highestTempCity.Temperature)
		case 2:
			lowestTempCity := findLowestTemp(cities)
			fmt.Printf("City with lowest temperature: %s %.2f\n", lowestTempCity.Name, lowestTempCity.Temperature)
		case 3:
			avgRainfall := calculateAvgRainfall(cities)
			fmt.Printf("Average rainfall: %.2f\n", avgRainfall)
		case 4:
			var threshold float64
			fmt.Print("Enter rainfall threshold: ")
			thresholdInput, _ := reader.ReadString('\n')
			thresholdInput = strings.TrimSpace(thresholdInput)
			threshold, err := strconv.ParseFloat(thresholdInput, 64)
			if err != nil {
				fmt.Println("Invalid threshold")
				return
			}
			filteredCities := filterByRainfall(cities, threshold)
			fmt.Println("Cities with rainfall greater than threshold: ")
			for _, city := range filteredCities {
				fmt.Printf("%s %.2f\n", city.Name, city.Rainfall)
			}
		case 5:
			var cityName string
			fmt.Print("Enter city name: ")
			cityNameInput, _ := reader.ReadString('\n')
			cityName = strings.TrimSpace(cityNameInput)
			city := searchCityByName(cities, cityName)
			if city != nil {
				fmt.Printf("City found: %s %.2f %.2f\n", city.Name, city.Temperature, city.Rainfall)
			} else {
				fmt.Println("City not found")
			}
		case 6:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invlaid Choice")
		}
		if choice == 6 {
			break
		}
	}
}
