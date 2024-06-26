package main

import (
	"fmt"
	"github.com/hconn7/wcChallenge/sources"

	"log"
)

func commandBytes() error {
	data, err := sources.Readfile("./sources/test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The total bytes are:", len(data))
	return nil
}
