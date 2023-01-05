package main

import "fmt"

func banyakDigit(angka int, banyak int) int {
	var tem int = banyak
	if angka > -1 && angka < 10 {
		tem++
		return tem
	}
	var ber int = angka
	ber = (ber - (ber % 10)) / 10
	tem++
	return banyakDigit(ber, tem)
}

func main() {
	var angka int = 12345
	hasil := banyakDigit(angka, 0)
	fmt.Println(hasil)
}
