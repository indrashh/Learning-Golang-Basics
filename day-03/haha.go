package main

import (
	"html/template"
	"net/http"
)

type Alamat struct {
    Kota     string
    Provinsi string
}
type Mahasiswa struct { 
    ID      int
    Nama    string
    Umur    int
    Jurusan string
    Alamat
}

var tpl *template.Template
var mhs1 Mahasiswa

func main(){
	mhs1 = Mahasiswa{
		1,"indra",21,"informatics",Alamat{"Semarang","Kendal"},
	}
	tpl, _ = template.ParseGlob("templates/*.html")
	http.HandleFunc("/", mahasiswaHandler)
    http.ListenAndServe(":8080", nil)
}

func mahasiswaHandler(w http.ResponseWriter,r *http.Request){
	tpl.ExecuteTemplate(w, "mahasiswainfo.html", mhs1)
}