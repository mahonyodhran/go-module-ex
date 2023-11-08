package main

import (
	"fmt"
	"testy.com/greetings"
)

func main(){
	message := greetings.Hello("Maggie")
	fmt.Println(message)
}
