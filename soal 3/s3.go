// Nama Ajr Inda Robby
// Internship - Problem Solving Test SE (BE)
package main

import (
	"bufio"
	"fmt"
	"os"
)

// Fungsi untuk memeriksa apakah string memiliki bracket yang seimbang
func isBalancedBracket(s string) string {
	stack := []rune{}
	bracketPairs := map[rune]rune{')': '(', '}': '{', ']': '['}

	for _, char := range s {
		switch char {
		case '(', '{', '[':
			stack = append(stack, char) // Push bracket buka
		case ')', '}', ']':
			if len(stack) == 0 || stack[len(stack)-1] != bracketPairs[char] {
				return "NO" // Tidak cocok atau stack kosong
			}
			stack = stack[:len(stack)-1] // Pop bracket cocok dari stack
		}
	}

	if len(stack) == 0 {
		return "YES"
	}
	return "NO"
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Masukkan bracket: ")
	scanner.Scan()
	input := scanner.Text()

	fmt.Println(isBalancedBracket(input))
}
