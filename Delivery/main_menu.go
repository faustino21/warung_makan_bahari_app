package Delivery

import (
	"fmt"
	"os"
	"strings"
)

func MainMenu() {
	fmt.Println("WARUNG MAKAN BAHARI")
	fmt.Println(strings.Repeat("*", 50))
	fmt.Println("1. Pesan Makanan")
	fmt.Println("2. Update Meja")
	fmt.Println("3. Update Payment")
	fmt.Println("Pilih menu (1-3) : ")
}

func Cancel() {
	os.Exit(1)
}
