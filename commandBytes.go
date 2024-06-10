package main

import (
	"fmt"
	"github.com/hconn7/wcChallenge/sources/readFile"
	"log"
)

func commandBytes() error {
	data, err := readFile.ReadFile()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The total bytes are:", len(data))
	return nil
}
