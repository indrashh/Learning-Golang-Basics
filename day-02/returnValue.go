package main

import "fmt"

type Student struct{
	Name string
	Age int
	Major string
}

func (s Student) getIntro() string{
	return fmt.Sprintf("Hi, my name is %s I'm %d, majoring in %s",s.Name,s.Age,s.Major)
}

func main(){
	std := Student{"Indra",21,"Informatics"}
	intro := std.getIntro()
	fmt.Println(intro)
}