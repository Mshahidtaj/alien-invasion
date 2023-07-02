package main

import (
	"os"
	"testing"
)

func TestReadWorldMap(t *testing.T) {
	// Create a temporary map file
	file, err := os.CreateTemp("", "map_test.txt")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(file.Name())

	// Write the sample map data to the file
	_, err = file.WriteString("City1 north=City2 west=City3\nCity2 south=City1\nCity3 east=City1\n")
	if err != nil {
		t.Fatalf("Failed to write to temporary file: %v", err)
	}

	worldMap, err := readWorldMap(file.Name())
	if err != nil {
		t.Errorf("Failed to read world map: %v", err)
	}
	file.Close()

	if len(worldMap) != 3 {
		t.Errorf("Expected 3 cities in the world map, but got %d", len(worldMap))
	}

	if _, exists := worldMap["City1"]; !exists {
		t.Error("City1 is missing from the world map")
	}

	if _, exists := worldMap["City2"]; !exists {
		t.Error("City2 is missing from the world map")
	}

	if _, exists := worldMap["City3"]; !exists {
		t.Error("City3 is missing from the world map")
	}
}
