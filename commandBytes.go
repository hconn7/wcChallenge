package main

import (
	"fmt"
	"github.com/wcChallenge/sources"
	"log"
)

func commandBytes() error {
	data, err := ReadFile()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The total bytes are:", len(data))
	return nil
}
