package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Mshahidtaj/alien-invasion/game"
)

func main() {
	var (
		numAliens    int
		worldMapFile string
	)

	// parsing the mandatory input
	flag.StringVar(&worldMapFile, "worldmap", "map.txt", "Provide file name to load World Map")
	flag.IntVar(&numAliens, "n", 5, "Provide no of Aliens")
	flag.Parse()

	fmt.Println("input: ", worldMapFile, numAliens)
	// validation

	// Simulate alien movement
	world, err := readWorldMap(worldMapFile)
	if err != nil {
		fmt.Printf("Failed to read world map: %v\n", err)
		return
	}

	// it has been assumed that there can be no more than twice the number of aliens as there are cities
	if numAliens > world.TotalCitiesInWorld()*2 {
		log.Fatal("Aliens cannot be greater then twice (2X) of number of cities", world.TotalCitiesInWorld())
	}

	fmt.Println("World Map before Alien Invasion: ", world.TotalCitiesInWorld())
	world.PrintWorldMap()
	simulateAlienInvasion(world, numAliens)
}

func simulateAlienInvasion(world game.WorldMap, numAliens int) {
	fmt.Println("\nHang on Alien Inavsion about to start:")
	aliens := world.CreateAliens(numAliens)

	fmt.Println("\nInitial Alien Mapping:")
	for _, alien := range aliens {
		fmt.Printf("Alien %d: City %s\n", alien.ID, alien.City.Name)
	}

	// Simulate alien movement, lets run 2 simulation
	world.SimulateAlienMovement(aliens)

	fmt.Println("\nRemaining World Map:")
	world.PrintWorldMap()
	fmt.Println("\n********************Simulation Completed:********************")
}

// ReadWorldMap reads the world map file and populates the WorldMap data structure
func readWorldMap(filename string) (game.WorldMap, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	worldMap := make(game.WorldMap)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		// The​ ​city​ ​and​ ​each​ ​of​ ​the​ ​pairs​ ​are​ ​separated​ ​by​ ​a​ ​single​ ​space,​ ​
		parts := strings.Split(line, " ")
		cityName := parts[0]
		directions := parts[1:]

		city := &game.City{
			Name:   cityName,
			Neighs: make(map[string]*game.City),
		}

		// directions​ ​are​ ​separated​ ​from​ ​their​ ​respective​ ​cities​ ​with​ ​an​ ​equals​ ​(=)​ ​sign.
		for _, dir := range directions {
			pair := strings.Split(dir, "=")
			dirName := pair[0]
			neighName := pair[1]

			neighCity := worldMap[neighName]
			if neighCity == nil {
				neighCity = &game.City{
					Name:   neighName,
					Neighs: make(map[string]*game.City),
				}
				worldMap[neighName] = neighCity
			}

			city.Neighs[dirName] = neighCity
			neighCity.Neighs[game.OppositeDirection(dirName)] = city
		}

		worldMap[cityName] = city
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return worldMap, nil
}
