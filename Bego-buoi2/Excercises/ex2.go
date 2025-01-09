package main

import "fmt"

// Kiểu kí tự trong golang
func stringProcess(s string) map[byte]int {

	m := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}

	return m
}

func unicodeStringProcess(s string) map[rune]int {
	m := make(map[rune]int)

	for _, char := range s {
		m[char]++
	}
	return m
}

// func main() {

// 	fmt.Println(stringProcess("abx"))
// 	fmt.Println(unicodeStringProcess("abx"))

// 	m := unicodeStringProcess("abx")

// 	for char, count := range m {
// 		fmt.Printf("%c:%v,", char, count)
// 	}
// }
