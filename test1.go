package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"regexp"
	"time"
)

type MadlibData struct {
	Nouns      []string  `json:"Nouns"`
	Numbers    []float64 `json:"Numbers"`
	Adjectives []string  `json:"Adjectives"`
}

func randomIndex(list []interface{}) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(len(list))
}

func loadJSON(filename string) (MadlibData, error) {
	var data MadlibData
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return data, err
	}

	err = json.Unmarshal(file, &data)
	if err != nil {
		return data, err
	}

	return data, nil
}

func replaceMadlibPlaceholders(madlib string, nouns []string, adjectives []string, numbers []float64) string {
	replacements := map[string]string{
		"(noun)":      nouns[randomIndexToInterface(nouns)],
		"(adjective)": adjectives[randomIndexToInterface(adjectives)],
		"(number)":    fmt.Sprintf("%v", numbers[randomIndexToInterface(numbers)]),
	}

	for placeholder, replacement := range replacements {
		re := regexp.MustCompile(placeholder)
		madlib = re.ReplaceAllString(madlib, replacement)
	}

	return madlib
}

func randomIndexToInterface(list []string) int {
	return rand.Intn(len(list))
}


func main() {
	data, err := loadJSON("data.json")
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return
	}

	// Read Madlib file
	madlibFile, err := ioutil.ReadFile("madlib.txt")
	if err != nil {
		fmt.Println("Error reading Madlib file:", err)
		return
	}

	madlib := string(madlibFile)
	// Replace placeholders with random selections
	madlib = replaceMadlibPlaceholders(madlib, data.Nouns, data.Adjectives, data.Numbers)

	// Output the completed Madlib
	fmt.Println("\nHere's your Madlib:")
	fmt.Println(madlib)
}
