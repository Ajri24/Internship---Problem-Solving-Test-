// Nama Ajr Inda Robby
// Internship - Problem Solving Test SE (BE)
package main

import (
	"fmt"
)

// Fungsi untuk membuat daftar peringkat unik
func getRankings(scores []int) []int {
	rankings := []int{}
	seen := map[int]bool{} // Set untuk menyimpan skor unik

	for _, score := range scores {
		if !seen[score] {
			rankings = append(rankings, score)
			seen[score] = true
		}
	}
	return rankings
}

// Fungsi untuk mencari peringkat GITS berdasarkan skor yang diperoleh
func getPlayerRanks(rankings, playerScores []int) []int {
	result := []int{}
	for _, score := range playerScores {
		rank := binarySearch(rankings, score)
		result = append(result, rank)
	}
	return result
}

// Fungsi Binary Search untuk mencari posisi peringkat
func binarySearch(rankings []int, score int) int {
	left, right := 0, len(rankings)-1

	for left <= right {
		mid := (left + right) / 2
		if rankings[mid] == score {
			return mid + 1
		} else if rankings[mid] < score {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left + 1
}

func main() {
	// Input jumlah pemain dan skor mereka
	var n int
	fmt.Print("Masukkan jumlah pemain: ")
	fmt.Scan(&n)

	scores := make([]int, n)
	fmt.Println("Masukkan skor pemain:")
	for i := 0; i < n; i++ {
		fmt.Scan(&scores[i])
	}

	// Input jumlah pertandingan GITS
	var m int
	fmt.Print("Masukkan jumlah pertandingan GITS: ")
	fmt.Scan(&m)

	playerScores := make([]int, m)
	fmt.Println("Masukkan skor GITS:")
	for i := 0; i < m; i++ {
		fmt.Scan(&playerScores[i])
	}

	// Mendapatkan daftar peringkat unik yang terurut
	rankings := getRankings(scores)

	// Mencari peringkat GITS
	results := getPlayerRanks(rankings, playerScores)

	// Menampilkan hasil
	fmt.Println("Peringkat GITS:", results)
}
