// Nama Ajr Inda Robby
// Internship - Problem Solving Test SE (BE)

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func A000124(n int) []int {
	sequence := []int{1} // Inisialisasi dengan elemen pertama

	for i := 2; i <= n; i++ {
		sequence = append(sequence, sequence[i-2]+i-1)
	}

	return sequence
}

// Fungsi untuk mengubah slice integer menjadi format string dengan tanda "-"
func formatSequence(sequence []int) string {
	strSequence := []string{}
	for _, num := range sequence {
		strSequence = append(strSequence, strconv.Itoa(num))
	}
	return strings.Join(strSequence, "-")
}

func main() {
	var n int
	fmt.Print("Masukkan angka: ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Input harus lebih dari 0")
		return
	}

	sequence := A000124(n)
	fmt.Println("Output:", formatSequence(sequence))
}
