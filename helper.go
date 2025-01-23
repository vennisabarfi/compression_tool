package main

import (
	"fmt"
	"log"
	"os"
)

func frequencyCounter(filename string) (frequency map[rune]int) {
	buffer, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	// convert buffer into string for counter
	s := string(buffer)
	// frequency counter that creates a map of runes and integers
	frequency = make(map[rune]int)
	// iterate over string and create count
	for _, char := range s {
		frequency[char] = frequency[char] + 1
	}
	return frequency
}

func FileCompressor(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	// check for error to be returned and exit if file not found
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	// extract frequency using frequency counter function
	frequency := frequencyCounter(filename)
	for char, count := range frequency {
		fmt.Printf("%c: %d\n", char, count)
	}

}
