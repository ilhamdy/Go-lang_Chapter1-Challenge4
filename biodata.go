package main

import (
	"fmt"
	"os"
)

type teman struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func main() {
	// Inisialisasi data teman-teman di kelas
	temanKelas := []teman{
		{"Andi", "Jakarta", "Software Engineer", "Ingin mempelajari bahasa pemrograman Go untuk meningkatkan skill programming."},
		{"Budi", "Bandung", "Data Analyst", "Mendapatkan pengalaman baru dalam menggunakan bahasa pemrograman Go."},
		{"Cindy", "Surabaya", "UI/UX Designer", "Tertarik dengan performa tinggi dan efisiensi Go dalam pengembangan aplikasi."},
		{"Deni", "Yogyakarta", "Backend Developer", "Ingin memperdalam skill di bidang backend development dengan mempelajari bahasa Go."},
		{"Eva", "Semarang", "Full-Stack Developer", "Ingin mempelajari bahasa pemrograman Go untuk mengembangkan aplikasi web yang lebih cepat dan efisien."},
	}

	// Mendapatkan argumen nomor absen dari command line
	args := os.Args[1:]

	// Konversi argumen menjadi integer
	absen := 0
	if len(args) > 0 {
		fmt.Sscanf(args[0], "%d", &absen)
	}

	// Menampilkan data teman dengan nomor absen yang sesuai
	if absen > 0 && absen <= len(temanKelas) {
		t := temanKelas[absen-1]
		fmt.Printf("Data teman dengan nomor absen %d:\n", absen)
		fmt.Printf("Nama      : %s\n", t.Nama)
		fmt.Printf("Alamat    : %s\n", t.Alamat)
		fmt.Printf("Pekerjaan : %s\n", t.Pekerjaan)
		fmt.Printf("Alasan    : %s\n", t.Alasan)
	} else {
		fmt.Println("Nomor absen tidak valid")
	}
}
