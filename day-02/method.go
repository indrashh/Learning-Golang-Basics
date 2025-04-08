package main

import (
	"fmt"
)

type Student struct{
	Name string
	Age int
}

func (s Student) Greetings(){
	fmt.Printf(`Hello, my name is %s i am %d years old `,s.Name,s.Age)
}

func main(){
	std := Student{"Indra",21}

	std.Greetings()
}