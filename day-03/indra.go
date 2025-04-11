package main

import (
	"html/template"
	"net/http"
)

type Alamat struct{
	City,Provincy  string
	
}

type Mahasiswa struct{
	Name string
	Age int
	Major string
	Alamat

}

var tpl *template.Template
var mhs1 Mahasiswa

func main(){
	mhs1 = Mahasiswa{
		"Indra",21,"Informatics",Alamat{"Semarang","kendal"},
	}
	tpl,_ = template.ParseGlob("templates/*.html")
	
	http.HandleFunc("/",mahasiswaHandler)
	http.ListenAndServe(":8080",nil)

}

func mahasiswaHandler(w http.ResponseWriter,r *http.Request){
	tpl.ExecuteTemplate(w,"mahasiswainfo.html",mhs1)
}