package main

import "fmt"

type Mahasiswa struct{
	Name,Major string
	Age int
}

func main(){
	mhs := Mahasiswa{
		Name:"indra",
		Major: "Informatics",
		Age:21,
	}

	fmt.Printf("Name: %s\nMajor: %s\nAge: %d\n", mhs.Name, mhs.Major, mhs.Age)

}