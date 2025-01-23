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
	// snippet that receives a string and outputs the frequency of each character
	// off to optimize for files probably using buffer instead of string?
	s := "namm"
	frequency := make(map[rune]int)
	for _, char := range s {
		frequency[char] = frequency[char] + 1
	}

	for char, count := range frequency {
		fmt.Printf("%c: %d\n", char, count)
	}

}
