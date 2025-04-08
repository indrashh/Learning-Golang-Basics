package main

import "fmt"

type Student struct{
	Name string
	Age int
	Major string
}

func (s *Student) updateProfile(newName string,newAge int,newMajor string){
	s.Name = newName
	s.Age = newAge
	s.Major = newMajor
}

func main(){
	std := Student{"Indra",21,"Informatics"}
	std.updateProfile("Bulan",19,"Indra's girlfriend")
	fmt.Println("My name: ", std.Name)
	fmt.Println("My age: ", std.Age)
	fmt.Println("My status: ", std.Major)
}