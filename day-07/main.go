package main

import (
	"html/template" // Buat parsing dan rendering file HTML template
	"log"           // Buat log server jalan atau debug pesan error
	"net/http"      // Untuk handle HTTP request dan response

	"gorm.io/driver/mysql" // MySQL driver-nya GORM
	"gorm.io/gorm"         // ORM-nya buat query ke database
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
        Name:  "Indra Agustyawan",
        Users: users,
    })
}

func createHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        http.Redirect(w, r, "/", http.StatusSeeOther)
        return
    }

    db, err := connectDB()
    if err != nil {
        http.Error(w, "Gagal koneksi ke database", http.StatusInternalServerError)
        return
    }

    nama := r.FormValue("nama_user")
    last := r.FormValue("last")
    handle := r.FormValue("handle")

    user := User{
        Nama_user: nama,
        Last:      last,
        Handle:    handle,
    }

    db.Create(&user)

    http.Redirect(w, r, "/", http.StatusSeeOther)
}


func main() {
    http.HandleFunc("/", homeHandler)
    http.HandleFunc("/add", createHandler)
    log.Println("Server berjalan di :8080")
    http.ListenAndServe(":8080", nil)
}
