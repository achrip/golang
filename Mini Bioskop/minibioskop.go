package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	input = bufio.NewScanner(os.Stdin)
)

func main() {
	var (
		film1 string = "Doraemon"
		film2 string = "Pokemon"
		film3 string = "Gundam"
		sc    int
	)
	nama := make([]string, 0)     // nama pembeli
	array_kode := make([]int, 0)  // kode film
	array_tiket := make([]int, 0) // jumlah tiket
	total := make([]int, 0)       // total pembayaran

	for {

		fmt.Print("\nMini Bioskop GOLANG")
		fmt.Print(" Ver. 0.1\n")
		fmt.Println("==========================")
		fmt.Println("| Kode Film | Nama Film  |")
		fmt.Println("| 1         |", film1, "  |")
		fmt.Println("| 2         |", film2, "   |")
		fmt.Println("| 3         |", film3, "    |")
		fmt.Println()

		fmt.Println("1. Beli Tiket")
		fmt.Println("2. Lihat Data")
		fmt.Println("3. Update Data")
		fmt.Println("4. Hapus Data")
		fmt.Println("5. Exit")
		fmt.Print("choose >> : ")

		fmt.Scan(&sc)

		switch sc {
		case 1:
			var (
				recheck string
				id      string
				mail    string
				kode    int
				tiket   int
				harga   int
				bayar   int
			)

			for {
				//ini tempat user input data diri
				fmt.Println()
				fmt.Print("Masukkan nama anda : ")
				input.Scan()
				id = input.Text()
				fmt.Print("Masukkan kode film : ")
				fmt.Scan(&kode)
				fmt.Print("Masukkan email anda : ")
				input.Scan()
				mail = input.Text()
				fmt.Print("Masukkan jumlah pembelian : ")
				fmt.Scan(&tiket)
				fmt.Println()
				fmt.Println()

				//validasi pembelian ke user. belom ditambahkan case jika summary salah
				fmt.Println("\t Summary \t")
				fmt.Println("==========================")
				fmt.Println("Nama : ", id, "\nEmail : ", mail, "\nKode Film : ", kode, "\nJumlah Tiket : ", tiket)

				fmt.Print("Apakah inormasi di atas sudah benar? <y/n> ")
				fmt.Scan(&recheck)

				if recheck == "y" {
					break
				}
			}
			nama = append(nama, id)
			array_tiket = append(array_tiket, tiket)
			array_kode = append(array_kode, kode)

			if kode == 1 {
				harga = 25000
			} else if kode == 2 {
				harga = 36000
			} else {
				harga = 30000
			}

			jumlah := harga * tiket
			fmt.Println("Jumlah pembayaran : ", jumlah)

			fmt.Print("Masukan nominal pembayaran : ")
			fmt.Scan(&bayar)
			change := bayar - jumlah

			if bayar >= jumlah {
				nama = append(nama, id)
				tiekt = append(tiekt, ticket)
				code = append(code, kode)
				total = append(total, jumlah)

				fmt.Println("Kembali: ", change)
				fmt.Println("Selamat menikmati tayangannya!")
			} else {
				fmt.Println("Pembayaran anda tidak dapat di proses.")
			}

		case 2: //viewing data
			if cap(nama) == 0 {
				fmt.Println("Data TIDAK Ditemukan!")
			} else {
				fmt.Println("========================================================")
				fmt.Println("| No. | Nama        | Kode Film | Jumlah    | Price    |")
				fmt.Println("========================================================")
			}

			for i := 0; i < len(nama); i++ {

				fmt.Printf("| %-3d| %-10s| %-10d| %-7d| %-5d|\n",
					(i + 1), nama[i], array_kode[i], array_tiket[i], total[i])
			}

		case 3: //editing data
			if cap(nama) == 0 {
				fmt.Println("Data TIDAK Ditemukan!")
			} else {
				fmt.Println("========================================================")
				fmt.Println("| No. | Nama        | Kode Film | Jumlah    | Price    |")
				fmt.Println("========================================================")
			}
			for i := 0; i < len(nama); i++ {
				fmt.Printf("| %-3d| %-10s| %-10d| %-7d| %-5d|\n",
					(i + 1), nama[i], array_kode[i], array_tiket[i], total[i])

			}

		case 4: //deleting data
			fmt.Println("========================================================")
			fmt.Println("| No. | Nama        | Kode Film | Jumlah    | Price    |")
			fmt.Println("========================================================")

			for i := 0; i < len(nama); i++ {
				fmt.Printf("| %-4d| %-12s| %-10d| %-10d| %-9d|\n",
					(i + 1), nama[i], array_kode[i], array_tiket[i], total[i])
			}
		}

		if sc == 5 {
			break
		}
	}
}
