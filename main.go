package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Konfigurasi koneksi database
	dsn := "mysql_wsl_username:mysql_wsl_password@tcp(localhost:3306)/dbname"

	// Membuka koneksi ke database
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Gagal membuka koneksi ke database: %v", err)
	}
	defer db.Close()

	// Menguji koneksi ke database
	err = db.Ping()
	if err != nil {
		log.Fatalf("Gagal melakukan ping ke database: %v", err)
	}

	fmt.Println("Koneksi ke database MySQL berhasil!")
}