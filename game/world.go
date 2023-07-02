package game

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// City represents a city in the world map
type City struct {
	Name   string           // Name of the city
	Neighs map[string]*City // Neighboring cities and their directions
	Aliens []*Alien         // Aliens currently in the city
}

// Alien represents an alien in the simulation
type Alien struct {
	ID     int   // ID of the alien
	Active bool  // Indicates if the alien is still active
	City   *City // Current city of the alien
}

// WorldMap represents the world map containing cities & then each city have link to its Neighbours
type WorldMap map[string]*City

// getRandomCityWithLimit selects a random city from the world map that doesn't exceed the alien limit
func (cities WorldMap) getRandomCityWithLimit(alienCount map[*City]int) *City {
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(cities))
	i := 0

	for _, city := range cities {
		// Skip the city if the alien limit has been reached
		if alienCount[city] >= 2 {
			continue
		}

		if i == index {
			return city
		}
		i++
	}

	// If no city can accommodate more aliens, return nil
	return nil
}

// createAliens creates the specified number of aliens and assigns them random starting cities
func (cities WorldMap) CreateAliens(numAliens int) []*Alien {
	aliens := make([]*Alien, numAliens)
	alienCount := make(map[*City]int) // Track the number of aliens in each city

	for i := 0; i < numAliens; i++ {
		alien := &Alien{
			ID:     i + 1,
			Active: true,
		}

		// Find a valid random city for the alien
		for {
			alien.City = cities.getRandomCityWithLimit(alienCount)
			if alien.City != nil {
				break
			}
		}

		aliens[i] = alien

		// Increment the alien count for the city
		alienCount[alien.City]++
	}

	return aliens
}

// simulateAlienMovement simulates aliens moving randomly around the cities
func (world WorldMap) SimulateAlienMovement(aliens []*Alien) {
	for _, alien := range aliens {
		if !alien.Active {
			// Skip inactive aliens
			continue
		}

		currentCity := alien.City

		// Get the neighboring cities of the current city
		neighboringCities := currentCity.Neighs

		if len(neighboringCities) == 0 {
			// Skip alien movement if there are no neighboring cities
			continue
		}

		// Select a random direction to move
		directions := make([]string, 0, len(neighboringCities))
		for dir := range neighboringCities {
			directions = append(directions, dir)
		}
		randomDirection := directions[rand.Intn(len(directions))]

		// Move the alien to the new city in the selected direction
		newCity := neighboringCities[randomDirection]

		if len(newCity.Aliens) > 0 {
			// Alien encounters another alien in the new city
			fmt.Printf("Alien %d encountered Alien %d in %s! They fought and destroyed the city.\n", alien.ID, newCity.Aliens[0].ID, newCity.Name)
			alien.Active = false
			newCity.Aliens[0].Active = false

			// Remove the aliens from the city's alien slice
			newCity.Aliens = removeAlienFromSlice(newCity.Aliens, alien)
			newCity.Aliens = removeAlienFromSlice(newCity.Aliens, newCity.Aliens[0])

			destroyCity(newCity, world)
		} else {
			// Move the alien to the new city
			currentCity.Aliens = removeAlienFromSlice(currentCity.Aliens, alien)
			alien.City = newCity
			newCity.Aliens = append(newCity.Aliens, alien)
			// Print the movement of the alien
			fmt.Printf("Alien %d moved from %s to %s\n", alien.ID, currentCity.Name, newCity.Name)
		}

	}
}

// destroyCity destroys the specified city and removes it from the world map
func destroyCity(city *City, worldMap WorldMap) {
	// Remove the city from its neighboring cities' connections
	for _, neighCity := range city.Neighs {
		delete(neighCity.Neighs, getDirectionToCity(neighCity, city))
	}

	// Remove the city from the world map
	delete(worldMap, city.Name)
}

// getDirectionToCity returns the direction from one city to another
func getDirectionToCity(fromCity, toCity *City) string {
	for dir, neighCity := range fromCity.Neighs {
		if neighCity == toCity {
			return dir
		}
	}
	return ""
}

// removeAlienFromSlice removes the specified alien from the slice
func removeAlienFromSlice(slice []*Alien, alien *Alien) []*Alien {
	index := -1
	for i, a := range slice {
		if a == alien {
			index = i
			break
		}
	}
	if index >= 0 {
		slice = append(slice[:index], slice[index+1:]...)
	}
	return slice
}

// oppositeDirection returns the opposite direction given a direction
func oppositeDirection(dir string) string {
	switch dir {
	case "north":
		return "south"
	case "south":
		return "north"
	case "east":
		return "west"
	case "west":
		return "east"
	default:
		return ""
	}
}

// printWorldMap prints the current state of the world map
func (world WorldMap) PrintWorldMap() {
	for _, city := range world {
		directions := make([]string, 0, len(city.Neighs))
		for dir := range city.Neighs {
			directions = append(directions, dir)
		}
		fmt.Printf("%s %s\n", city.Name, strings.Join(directions, " "))
	}
}

func (cities WorldMap) totalCitiesInWorld() int {
	return (len(cities))
}
