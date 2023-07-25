package main

import (
	"fmt"
	"log"
	"greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Gladys", "Samantha", "Darring"}

	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}
	
	for _, msg := range messages {
		fmt.Println(msg)
	}
}
