package main

import (
	"fmt"      // Paket untuk mencetak ke layar
	"net/http" // Paket untuk melakukan permintaan HTTP
	"time"     // Paket untuk mengukur waktu
)

func main() {
	// Daftar URL yang akan diuji
	urls := []string{
		"https://www.google.com",
		"https://www.cloudflare.com",
	}

	// Loop utama yang berjalan terus menerus
	for {
		// Loop melalui setiap URL dalam daftar
		for _, url := range urls {
			// Mencatat waktu mulai sebelum melakukan permintaan
			start := time.Now()

			// Melakukan permintaan HTTP GET ke URL
			resp, err := http.Get(url)
			// Jika ada kesalahan dalam permintaan, cetak kesalahan dan lanjutkan ke URL berikutnya
			if err != nil {
				fmt.Printf("Error fetching %s: %v\n", url, err)
				continue
			}

			// Menutup body respons untuk mencegah kebocoran sumber daya
			resp.Body.Close()

			// Menghitung waktu yang dibutuhkan untuk mendapatkan respons
			elapsed := time.Since(start)
			// Mencetak waktu respons ke layar
			fmt.Printf("Response time for %s: %v\n", url, elapsed)
		}

		// Menunggu selama 10 detik sebelum mengulangi pengecekan
		time.Sleep(10 * time.Second)
	}
}
