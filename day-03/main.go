package main

import (
	"html/template"
	"net/http"
)

// Struct untuk menyimpan informasi mahasiswa dan alamatnya
type Alamat struct {
    Kota     string
    Provinsi string
}
type Mahasiswa struct {
    ID      int
    Nama    string
    Umur    int
    Jurusan string
    Alamat  Alamat
}

var tpl *template.Template
var mhs1 Mahasiswa

func main() {
    // Inisialisasi data mahasiswa
    mhs1 = Mahasiswa{
        ID:      101,
        Nama:    "Andi Setiawan",
        Umur:    21,
        Jurusan: "Informatika",
        Alamat: Alamat{
            Kota:     "Semarang",
            Provinsi: "jawa Tengah",
        },
    }
    tpl, _ = template.ParseGlob("templates/*.html")
		
    http.HandleFunc("/mahasiswainfo", mahasiswaInfoHandler)
    http.ListenAndServe(":8080", nil)
}
func mahasiswaInfoHandler(w http.ResponseWriter, r *http.Request) {
    tpl.ExecuteTemplate(w, "mahasiswainfo.html", mhs1)
}
