package main

import "fmt"

// func main() {

// 	// Cách dùng 1
// 	m1 := make(map[int]string) // khoi tao map
// 	m1[42] = "Hello"           // write ~ gan gia tri

// 	value, exist := m1[43]    // read ~ doc gia tri tu map: 1. No kiem tra co ton tai va tra ve exist 2. Khi ton tai no tra ve value cua map, neu ko thi tra ve rong thi p
// 	fmt.Println(value, exist) // So sánh với QMap và QHash trong QT

// 	// Dùng vs empty struct
// 	set := make(map[int]struct{})
// 	set[42] = struct{}{}
// 	_, exist1 := set[42]
// 	if exist1 {
// 		fmt.Println("42 is in set")
// 	}

// 	//3. Sự ngẫu nhiên trong map
// 	m := make(map[int]int)

// 	// Muc dich: Test conclusion: Second loop of map provide different order campare with first loop of map
// 	for j := 0; j < 10; j++ {
// 		m[j] = j * 2
// 	}

// 	// Duyệt map lần 1
// 	for k, v := range m {
// 		fmt.Printf("%d*2=%d ", k, v)
// 	}

// 	fmt.Printf("\n ----- \n")

// 	// Duyệt map lần 2
// 	for k, v := range m {
// 		fmt.Printf("%d*2=%d ", k, v)
// 	}

// }
