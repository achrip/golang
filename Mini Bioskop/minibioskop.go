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
	nama := make([]string, 0) // nama pembeli
	code := make([]int, 0)    // kode film
	tiekt := make([]int, 0)   // jumlah tiket
	total := make([]int, 0)   // total pembayaran

	for {

		fmt.Println("BIOSKOP GOLANG")
		fmt.Println("==============")
		fmt.Println("| Kode Film | Nama Film  |")
		fmt.Println("| 1         |", film1, "  |")
		fmt.Println("| 2         |", film2, "   |")
		fmt.Println("| 3         |", film3, "    |")
		fmt.Println()

		fmt.Println("1. Buy")
		fmt.Println("2. View")
		fmt.Println("3. Update")
		fmt.Println("4. Delete")
		fmt.Println("5. Exit")
		fmt.Println("choose >> ")

		fmt.Scan(&sc)

		switch sc {
		case 1:
			var (
				recheck string
				id      string
				mail    string
				kode    int
				ticket  int
				price   int
				bayar   int
			)

			for {
				//ini tempat user input data diri serta m
				fmt.Println()
				fmt.Println("Masukkan nama anda : ")
				input.Scan()
				id = input.Text()
				fmt.Println("Masukkan kode film : ")
				fmt.Scan(&kode)
				fmt.Println("Masukkan email anda : ")
				input.Scan()
				mail = input.Text()
				fmt.Println("Masukkan jumlah pembelian : ")
				fmt.Scan(&ticket)
				fmt.Println()
				fmt.Println()

				//validasi pembelian ke user. belom ditambahkan case jika summary salah
				fmt.Println("\t Summary \t")
				fmt.Println("==========================")
				fmt.Println("Nama : ", id, "\nEmail : ", mail, "\nKode Film : ", kode, "\nJumlah Tiket : ", ticket)

				fmt.Print("Apakah inormasi di atas sudah benar? <y/n> ")
				fmt.Scan(&recheck)

				if recheck == "y" {
					break
				}
			}

			if kode == 1 {
				price = 25000
			} else if kode == 2 {
				price = 36000
			} else {
				price = 30000
			}

			jumlah := price * ticket
			fmt.Println("Jumlah pembayaran : ", jumlah)

			fmt.Print("Masukkn nominal pembayaran : ")
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
				fmt.Println("There is no data!")
			} else {
				fmt.Println("========================================================")
				fmt.Println("| No. | Nama        | Kode Film | Jumlah    | Price    |")
				fmt.Println("========================================================")
			}

			for i := 0; i < len(nama); i++ {
				fmt.Printf("| %-4d| %-12s| %-10d| %-10d| %-9d|\n",
					(i + 1), nama[i], code[i], tiekt[i], total[i])
			}

		case 3: //editing data
			if cap(nama) == 0 {
				fmt.Println("There is no data!")
			} else {
				fmt.Println("========================================================")
				fmt.Println("| No. | Nama        | Kode Film | Jumlah    | Price    |")
				fmt.Println("========================================================")
			}
			for i := 0; i < len(nama); i++ {
				fmt.Printf("| %-4d| %-12s| %-10d| %-10d| %-9d|\n",
					(i + 1), nama[i], code[i], tiekt[i], total[i])
			}

		case 4: //deleting data
			fmt.Println("========================================================")
			fmt.Println("| No. | Nama        | Kode Film | Jumlah    | Price    |")
			fmt.Println("========================================================")

			for i := 0; i < len(nama); i++ {
				fmt.Printf("| %-4d| %-12s| %-10d| %-10d| %-9d|\n",
					(i + 1), nama[i], code[i], tiekt[i], total[i])
			}
		}

		if sc == 5 {
			break
		}
	}
}
