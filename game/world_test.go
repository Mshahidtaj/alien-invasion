package game

import (
	"testing"
)

func TestCreateAliens(t *testing.T) {
	city1 := &City{Name: "City1"}
	city2 := &City{Name: "City2"}
	city3 := &City{Name: "City3"}

	cities := WorldMap{
		"City1": city1,
		"City2": city2,
		"City3": city3,
	}

	numAliens := 5
	aliens := cities.createAliens(numAliens)

	if len(aliens) != numAliens {
		t.Errorf("Expected %d aliens, but got %d", numAliens, len(aliens))
	}

	for _, alien := range aliens {
		if alien.ID <= 0 {
			t.Errorf("Alien ID should be greater than 0, but got %d", alien.ID)
		}
		if alien.City == nil {
			t.Error("Alien should have a valid city assigned")
		}
	}
}

func TestSimulateAlienMovement(t *testing.T) {
	city1 := &City{Name: "City1"}
	city2 := &City{Name: "City2"}
	city3 := &City{Name: "City3"}

	cities := WorldMap{
		"City1": city1,
		"City2": city2,
		"City3": city3,
	}

	alien1 := &Alien{ID: 1, Active: true, City: city1}
	alien2 := &Alien{ID: 2, Active: true, City: city2}

	aliens := []*Alien{alien1, alien2}

	cities.simulateAlienMovement(aliens)

	if (alien1.City != city1 && alien1.City != city2 && alien1.City != city3) ||
		(alien2.City != city1 && alien2.City != city2 && alien2.City != city3) {
		t.Error("Aliens should move to a valid city")
	}
}

func TestDestroyCity(t *testing.T) {
	city1 := &City{Name: "City1"}
	city2 := &City{Name: "City2"}
	city3 := &City{Name: "City3"}

	city1.Neighs = map[string]*City{
		"north": city2,
		"west":  city3,
	}

	city2.Neighs = map[string]*City{
		"south": city1,
	}

	city3.Neighs = map[string]*City{
		"east": city1,
	}

	worldMap := WorldMap{
		"City1": city1,
		"City2": city2,
		"City3": city3,
	}

	destroyCity(city1, worldMap)

	if _, exists := worldMap["City1"]; exists {
		t.Error("City1 should be destroyed and removed from the world map")
	}

	if len(city2.Neighs) != 0 || len(city3.Neighs) != 0 {
		t.Error("The neighboring cities should have their connections to City1 removed")
	}
}

func TestGetDirectionToCity(t *testing.T) {
	city1 := &City{Name: "City1"}
	city2 := &City{Name: "City2"}

	direction := getDirectionToCity(city1, city2)

	if direction != "" {
		t.Errorf("Expected empty direction, but got %s", direction)
	}

	city1.Neighs = map[string]*City{
		"north": city2,
	}

	direction = getDirectionToCity(city1, city2)

	if direction != "north" {
		t.Errorf("Expected direction 'north', but got %s", direction)
	}
}

func TestRemoveAlienFromSlice(t *testing.T) {
	alien1 := &Alien{ID: 1}
	alien2 := &Alien{ID: 2}
	alien3 := &Alien{ID: 3}

	slice := []*Alien{alien1, alien2, alien3}

	newSlice := removeAlienFromSlice(slice, alien2)

	if len(newSlice) != 2 {
		t.Errorf("Expected slice length 2, but got %d", len(newSlice))
	}

	if newSlice[0] != alien1 || newSlice[1] != alien3 {
		t.Error("Unexpected aliens in the new slice")
	}

	newSlice = removeAlienFromSlice(slice, alien2)

	if len(newSlice) != 3 {
		t.Errorf("Expected slice length 3, but got %d", len(newSlice))
	}
}

func TestOppositeDirection(t *testing.T) {
	direction := oppositeDirection("north")
	if direction != "south" {
		t.Errorf("Expected direction 'south', but got %s", direction)
	}

	direction = oppositeDirection("south")
	if direction != "north" {
		t.Errorf("Expected direction 'north', but got %s", direction)
	}

	direction = oppositeDirection("east")
	if direction != "west" {
		t.Errorf("Expected direction 'west', but got %s", direction)
	}

	direction = oppositeDirection("west")
	if direction != "east" {
		t.Errorf("Expected direction 'east', but got %s", direction)
	}

	direction = oppositeDirection("invalid")
	if direction != "" {
		t.Errorf("Expected empty direction, but got %s", direction)
	}
}

func TestTotalCitiesInWorld(t *testing.T) {
	city1 := &City{Name: "City1"}
	city2 := &City{Name: "City2"}
	city3 := &City{Name: "City3"}

	worldMap := WorldMap{
		"City1": city1,
		"City2": city2,
		"City3": city3,
	}

	totalCities := worldMap.totalCitiesInWorld()

	if totalCities != 3 {
		t.Errorf("Expected total cities 3, but got %d", totalCities)
	}
}
