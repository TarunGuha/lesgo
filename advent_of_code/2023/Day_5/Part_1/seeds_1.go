package main

import (
	"bufio"
	"fmt"
	// "math"
	"os"
	"strconv"
	"strings"
)

func row_creator(line string) []int {
	var temp_row []int
	entries := strings.Split(strings.TrimSpace(line), " ")
	for _, element := range entries {
		if element != "" {
			integer_value, _ := strconv.Atoi(element)
			temp_row = append(temp_row, integer_value)
		}
	}
	return temp_row
}

func mapper(source int, mappings [][]int) int {
	for _, record := range mappings {
		if source >= record[1] && source < record[1]+record[2] {
			return (record[0] + (source - record[1]))
		}
	}
	return source
}

func find_minimum(list []int) int {

	min := list[0]

	for _, num := range list {
		if num < min {
			min = num
		}
	}

	return min
}

func main() {

	readFile, err := os.Open("almanac.txt")

	if err != nil {
		fmt.Println(err)
	}

	var min_location int

	var seed_locations []int
	var seed_to_soil [][]int
	var soil_to_fertilizer [][]int
	var fertilizer_to_water [][]int
	var water_to_light [][]int
	var light_to_temperature [][]int
	var temperature_to_humidity [][]int
	var humidity_to_location [][]int
	var locations []int

	seed_to_soil_entry := false
	soil_to_fertilizer_entry := false
	fertilizer_to_water_entry := false
	water_to_light_entry := false
	light_to_temperature_entry := false
	temperature_to_humidity_entry := false
	humidity_to_location_entry := false

	fileScanner := bufio.NewScanner(readFile)
	for fileScanner.Scan() {
		str := fileScanner.Text()

		if len(str) == 0 {
			continue
		}

		if len(str) >= 5 && str[0:5] == "seeds" {
			seed_locations_string := strings.TrimSpace(strings.Split(str, ":")[1])
			seed_locations_strings := strings.Split(seed_locations_string, " ")
			for _, element := range seed_locations_strings {
				if element != "" {
					integer_value, _ := strconv.Atoi(element)
					seed_locations = append(seed_locations, integer_value)
				}
			}
			continue
		}

		if len(str) >= 12 && str[0:12] == "seed-to-soil" {
			seed_to_soil_entry = true
			continue
		}

		if len(str) >= 18 && str[0:18] == "soil-to-fertilizer" {
			seed_to_soil_entry = false
			soil_to_fertilizer_entry = true
			continue
		}

		if len(str) >= 19 && str[0:19] == "fertilizer-to-water" {
			soil_to_fertilizer_entry = false
			fertilizer_to_water_entry = true
			continue
		}

		if len(str) >= 14 && str[0:14] == "water-to-light" {
			fertilizer_to_water_entry = false
			water_to_light_entry = true
			continue
		}

		if len(str) >= 20 && str[0:20] == "light-to-temperature" {
			water_to_light_entry = false
			light_to_temperature_entry = true
			continue
		}

		if len(str) >= 23 && str[0:23] == "temperature-to-humidity" {
			light_to_temperature_entry = false
			temperature_to_humidity_entry = true
			continue
		}

		if len(str) >= 20 && str[0:20] == "humidity-to-location" {
			temperature_to_humidity_entry = false
			humidity_to_location_entry = true
			continue
		}

		entry := row_creator(str)

		if seed_to_soil_entry {
			seed_to_soil = append(seed_to_soil, entry)
		}

		if soil_to_fertilizer_entry {
			soil_to_fertilizer = append(soil_to_fertilizer, entry)
		}

		if fertilizer_to_water_entry {
			fertilizer_to_water = append(fertilizer_to_water, entry)
		}

		if water_to_light_entry {
			water_to_light = append(water_to_light, entry)
		}

		if light_to_temperature_entry {
			light_to_temperature = append(light_to_temperature, entry)
		}

		if temperature_to_humidity_entry {
			temperature_to_humidity = append(temperature_to_humidity, entry)
		}

		if humidity_to_location_entry {
			humidity_to_location = append(humidity_to_location, entry)
		}

	}

	for _, seed := range seed_locations {
		soil := mapper(seed, seed_to_soil)
		fertilizer := mapper(soil, soil_to_fertilizer)
		water := mapper(fertilizer, fertilizer_to_water)
		light := mapper(water, water_to_light)
		temperature := mapper(light, light_to_temperature)
		humidity := mapper(temperature, temperature_to_humidity)
		location := mapper(humidity, humidity_to_location)
		locations = append(locations, location)
	}

	min_location = find_minimum(locations)

	fmt.Println(min_location)
}
