package main

import (
	"fmt"
	"os"
)

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

	// buffer, err := os.ReadFile(filename)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(buffer)
	s := "namm"
	// create a map of runes to ints and increase values as we iterate over a given string
	counter := make(map[rune]int)

	for _, c := range s {
		if c != ' ' {
			counter[c] += 1
		}
	}
	fmt.Println(counter)
}
