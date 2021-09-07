package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/twpayne/go-polyline"
)

func main() {
	var path string
	fmt.Println("enter file path")
	fmt.Scanln(&path)
	file, err := os.Open(path)

	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}

	fileScanner := bufio.NewScanner(file)
	var coords [][]float64

	for fileScanner.Scan() {
		textLine := fileScanner.Text()
		if existsCoordinates(textLine) {
			coord := findCoordinates(textLine)
			coords = append(coords, coord)
		}
	}

	if err := fileScanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}
	file.Close()

	encodeCoords(coords)
}

func existsCoordinates(line string) bool {
	return strings.Contains(line, "latitude") && strings.Contains(line, "longitude")
}

func findCoordinates(line string) (coord []float64) {
	latitudeIndex := strings.Index(line, "latitude")
	longitudeIndex := len(line) - 1

	latLongText := strings.ReplaceAll(line[latitudeIndex:longitudeIndex], ", ", " ")

	latLongParts := strings.Split(latLongText, " ")
	latitudeString := latLongParts[1]
	longitudeString := latLongParts[3]

	if latitude, err := strconv.ParseFloat(latitudeString, 64); err == nil {
		if longitude, err := strconv.ParseFloat(longitudeString, 64); err == nil {
			coord = append(coord, latitude)
			coord = append(coord, longitude)
		}
	}
	return
}

func encodeCoords(coords [][]float64) {
	fmt.Println(string(polyline.EncodeCoords(coords)))
}
