package main

import (
	"log"
	"os"
)

func ReadFile() ([]byte, error) {
	fileName := "./test.txt"
	data, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	return data, nil
}
