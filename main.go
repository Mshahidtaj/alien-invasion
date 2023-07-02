package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Simulate alien movement
	world, err := readWorldMap("map.txt")
	if err != nil {
		fmt.Printf("Failed to read world map: %v\n", err)
		return
	}

	fmt.Println("World Map before Alien Invasion: ", world.totalCitiesInWorld())
	world.printWorldMap()

	numAliens := 10
	simulateAlienInvasion(world, numAliens)
}

func simulateAlienInvasion(world WorldMap, numAliens int) {
	fmt.Println("\nHang on Alien Inavsion about to start:")
	aliens := world.createAliens(numAliens)

	fmt.Println("\nInitial Alien Mapping:")
	for _, alien := range aliens {
		fmt.Printf("Alien %d: City %s\n", alien.ID, alien.City.Name)
	}

	// Simulate alien movement, lets run 2 simulation
	world.simulateAlienMovement(aliens)

	fmt.Println("\nRemaining World Map1:")
	world.printWorldMap()
	fmt.Println("\n********************Simulation Completed:********************")
}

// readWorldMap reads the world map file and populates the WorldMap data structure
func readWorldMap(filename string) (WorldMap, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	worldMap := make(WorldMap)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		// The​ ​city​ ​and​ ​each​ ​of​ ​the​ ​pairs​ ​are​ ​separated​ ​by​ ​a​ ​single​ ​space,​ ​
		parts := strings.Split(line, " ")
		cityName := parts[0]
		directions := parts[1:]

		city := &City{
			Name:   cityName,
			Neighs: make(map[string]*City),
		}

		// directions​ ​are​ ​separated​ ​from​ ​their​ ​respective​ ​cities​ ​with​ ​an​ ​equals​ ​(=)​ ​sign.
		for _, dir := range directions {
			pair := strings.Split(dir, "=")
			dirName := pair[0]
			neighName := pair[1]

			neighCity := worldMap[neighName]
			if neighCity == nil {
				neighCity = &City{
					Name:   neighName,
					Neighs: make(map[string]*City),
				}
				worldMap[neighName] = neighCity
			}

			city.Neighs[dirName] = neighCity
			neighCity.Neighs[oppositeDirection(dirName)] = city
		}

		worldMap[cityName] = city
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return worldMap, nil
}
