package main

import (
	"fmt"
	"github.com/hconn7/wcChallenge/sources"
	"log"
	"strings"
)

func commandChars() error {
	totalChars := []string{}

	data, err := sources.Readfile("./sources/test.txt")
	if err != nil {
		log.Fatal(err)
	}
	myStrings := string(data)
	completeType := strings.Fields(myStrings)

	for _, word := range completeType {
		for _, char := range word {
			totalChars = append(totalChars, string(char))
		}
	}
	fmt.Println(len(totalChars))
	return nil
}
