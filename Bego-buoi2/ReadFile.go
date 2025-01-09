// Read whole file --> string

package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
)

// // Read whole file --> string
// func main() {
// 	//content, err := ioutil.ReadFile("themopylae.txt")
// 	content, err := os.ReadFile("themopylae.txt")
// 		// Tại sao lại không cần close hay open?
// 		// Tại sao lại không cần kiểm tra EOF?

// 	// Nếu có lỗi
// 	if err != nil {
// 		log.Fatal(err) // Fatal = Print + os.exit()
// 	}

// 	fmt.Println(string(content))
// }

// Read lines or words of file
func main() {
	f, err := os.Open("themopylae.txt") // Để public cho package khác thì dùng chữ viết hoa đầu function

	// Kiểm tra xem có lỗi khi open không
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords) // Commant hoặc uncomment dòng này

	for scanner.Scan() { // Làm biến đổi s.token() // Default split function of scanner is ScanLines
		fmt.Println(scanner.Text()) // Hàm Text() trả về s.token() // là từng dòng
		fmt.Println("\n")
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

// // Read file in chunks
// func main() {

// 	// Block of code for open files
// 	f, err := os.Open("themopylae.txt")

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defer f.Close()

// 	// Nó sử dụng reader thay cho scan
// 	reader := bufio.NewReader(f)
// 	buf := make([]byte, 16)

// 	// Vòng lặp đọc khoảng 16 byte mỗi lần cho đến khi gặp lỗi hoặc cuối file
// 	for {
// 		n, err := reader.Read(buf)  // Đọc và lưu vào biến buf

// 		if err != nil {

// 			if err != io.EOF {
// 				log.Fatal(err)
// 			}
// 			break // error là EOF thì break vòng for
// 		}
// 		fmt.Print(string(buf[0:n])) // Đọc ra từ biến buf, sử dụng slice
// 	}
// 	fmt.Println()
// }

// // // Read binary file
// func main() {
// 	f, err := os.Open("28427991_job204-wan-07.jpg")

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defer f.Close()

// 	reader := bufio.NewReader(f)
// 	buf := make([]byte, 256)

// 	// Cấu trúc check lỗi tương tự vòng lặp đọc từng chunk of data, mỗi chunk có 256 bytes thay vì 16 bytes
// 	// như hàm read chunk of data
// 	for {
// 		_, err := reader.Read(buf)
// 		if err != nil {
// 			if err != io.EOF {
// 				fmt.Println(err)
// 			}
// 			break
// 		}

// 		fmt.Printf("%s", hex.Dump(buf))
// 	}

// }
