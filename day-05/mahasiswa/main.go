package main

import (
	"html/template"
	"log"
	"net/http"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Model untuk data tabel
type User struct {
    ID        uint
    Nama_user string
    Last      string
    Handle    string
}

// Koneksi ke database
func connectDB() (*gorm.DB, error) {
    dsn := "root:@tcp(localhost:3306)/golang?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, err
    }
    return db, nil
}

// Handler untuk halaman utama
func homeHandler(w http.ResponseWriter, r *http.Request) {
    db, err := connectDB()
    if err != nil {
        http.Error(w, "Gagal koneksi ke database", http.StatusInternalServerError)
        return
    }

    // Ambil data dari database
    var users []User
    db.Find(&users)

    // Parse template
    tmpl, err := template.ParseFiles("template/index.html")
    if err != nil {
        http.Error(w, "Gagal memuat template", http.StatusInternalServerError)
        return
    }

    // Kirim data ke template
    tmpl.Execute(w, struct {
        Name  string
        Users []User
    }{
        Name:  "Golang Developer",
        Users: users,
    })
}

func main() {
    http.HandleFunc("/", homeHandler)
    log.Println("Server berjalan di :8080")
    http.ListenAndServe(":8080", nil)
}