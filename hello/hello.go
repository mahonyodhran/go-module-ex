package main

import (
	"fmt"
	"log"

	"testy.com/greetings"
)

func main(){
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	
	message, err := greetings.Hello("Maggie")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
