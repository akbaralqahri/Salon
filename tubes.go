package main

import (
	"fmt"
	"strings"
)

const NMAX = 100

type transaksi struct {
	namaPelanggan, menu, member string
	harga                       int
	total                       float64
}

type tabTransaksi [NMAX]transaksi

func main() {
	var tab tabTransaksi
	var n, trTinggi, trRendah int

	inputData(&tab, &n)

	fmt.Println("\nRata-rata transaksi:", rataRata(tab, n))

	trTinggi = transaksiTertinggi(tab, n)
	fmt.Println("\nTransaksi terbesar:", "Nama:", tab[trTinggi].namaPelanggan, "|", "Menu:", tab[trTinggi].menu, "|", "Harga:", tab[trTinggi].harga, "|", "Member:", tab[trTinggi].member, "|", "Total:", tab[trTinggi].total)

	trRendah = transaksiTerendah(tab, n)
	fmt.Println("\nTransaksi terkecil:", "Nama:", tab[trRendah].namaPelanggan, "|", "Menu:", tab[trRendah].menu, "|", "Harga:", tab[trRendah].harga, "|", "Member:", tab[trRendah].member, "|", "Total:", tab[trRendah].total)

	urutkan(&tab, n)

	fmt.Println("\nTransaksi terurut:")
	for i := 0; i < n; i++ {
		fmt.Println("Nama:", tab[i].namaPelanggan, "|", "Menu:", tab[i].menu, "|", "Harga:", tab[i].harga, "|", "Member:", tab[i].member, "|", "Total:", tab[i].total)
	}
}

// prosedur
func daftarJasa() {
	fmt.Println(
		"Bangs Cut : 20k\n",
		"Haircut : 40k\n",
		"Hair Wash : 20k\n",
		"Set or Style : 60k\n",
		"Creambath : 100k\n",
		"Mullet : 100k",
	)
}

// prosedur
func inputData(t *tabTransaksi, n *int) {
	var (
		namaPelanggan, menu, member string
		harga                       float64
	)

	fmt.Print("Masukkan banyak transaksi: ")
	fmt.Scanln(n)

	for i := 0; i < *n; i++ {
		daftarJasa()

		fmt.Print("Masukkan Nama Anda: ")
		fmt.Scanln(&namaPelanggan)

		pilihModel(&menu, &harga)

		fmt.Printf("%s, apakah kamu member? (no/epic/legend/mythic) ", namaPelanggan)
		fmt.Scanln(&member)

		fmt.Printf("%s, kamu memilih %s, dengan harga %.0f\n", namaPelanggan, menu, harga)

		t[i].namaPelanggan = namaPelanggan
		t[i].menu = menu
		t[i].harga = int(harga)
		t[i].member = member
		t[i].total = harga - (harga * hitungHargaMember(member))

		fmt.Println("Harga awal: Rp", t[i].harga)
		fmt.Println("Diskon member: Rp", float64(t[i].harga)*hitungHargaMember(member))
		fmt.Println("Total harga: Rp", t[i].total)
		fmt.Println()
	}
}

// prosedur
func pilihModel(menu *string, harga *float64) {
	fmt.Print("Masukkan model yang kamu pilih: ")
	fmt.Scanln(menu)

	for {
		if strings.ToUpper(*menu) == "BANGSCUT" {
			*harga = 20000
			break
		} else if strings.ToUpper(*menu) == "HAIRCUT" {
			*harga = 40000
			break
		} else if strings.ToUpper(*menu) == "HAIRWASH" {
			*harga = 20000
			break
		} else if strings.ToUpper(*menu) == "SETORSTYLE" {
			*harga = 60000
			break
		} else if strings.ToUpper(*menu) == "CREAMBATH" {
			*harga = 100000
			break
		} else if strings.ToUpper(*menu) == "MULLET" {
			*harga = 100000
			break
		} else {
			fmt.Println("Menu tidak tersedia, silahkan masukkan menu yang tersedia")
			fmt.Print("Masukkan model yang kamu pilih: ")
			fmt.Scanln(menu)
		}
	}
}

// prosedur
func urutkan(t *tabTransaksi, n int) {
	var max int
	var temp transaksi

	for i := 0; i < n; i++ {
		max = i

		for j := i + 1; j < n; j++ {
			if t[j].total > t[max].total {
				max = j
			}
		}

		temp = t[i]
		t[i] = t[max]
		t[max] = temp
	}
}

// function
func hitungHargaMember(member string) float64 {
	var total float64

	if member == "epic" {
		total = 0.1
	} else if member == "legend" {
		total = 0.2
	} else if member == "mythic" {
		total = 0.3
	} else {
		total = 0
	}

	return total
}

// function
func rataRata(t tabTransaksi, n int) float64 {
	var total float64

	for i := 0; i < n; i++ {
		total += t[i].total
	}

	return total / float64(n)
}

// function
func transaksiTertinggi(t tabTransaksi, n int) int {
	var max int

	max = 0

	for i := 0; i < n; i++ {
		if t[i].total > t[max].total {
			max = i
		}
	}

	return max
}

func transaksiTerendah(t tabTransaksi, n int) int {
	var min int

	min = 0

	for i := 0; i < n; i++ {
		if t[i].total < t[min].total {
			min = i
		}
	}

	return min
}
