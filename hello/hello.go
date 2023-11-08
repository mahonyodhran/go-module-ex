package main

import (
	"fmt"
	"log"

	"testy.com/greetings"
)

func main(){
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	
	names := []string{"Maggie", "Lisa", "Bart"}

	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
}
