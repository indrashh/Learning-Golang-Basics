package main

import "fmt"

type Address struct {
	Country, City string
}

type Student struct {
	Name string
	Age  int
	Address
}

func main() {
	std := Student{"Indra", 21, Address{"Indonesia", "Semarang"}}
	fmt.Printf(`Name: %s
Age: %d
Country: %s
City: %s
	
	`,std.Name,std.Age,std.Country,std.City)
}