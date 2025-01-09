package main

import "fmt"

// Passing a slice, not a array
func twoSum(nums []int, target int) []int {

	for i, v := range nums {
		temp := target - v

		for j, v1 := range nums {
			if temp == v1 && i != j {
				a := []int{i, j}
				return a
			}
		}
	}
	a := []int{-10000, -10000}
	return a
}

// Cach 2: Dng hash map
func twoSumWithHashMap(nums []int, target int) []int {
	numMap := make(map[int]int) // cấu trúc map ngược lại: Từ phần tử đến index, Và sẽ tăng dần

	for i, v := range nums {
		complement := target - v

		// nếu complement nằm trong mảng thì sẽ phải hash được
		if index, found := numMap[complement]; found {
			return []int{i, index}
		}

		// Nếu complement chưa có trong hash map thì thêm vào. Bt mình sẽ tạo map trước (như thế lặp) chứ ko phải xây dần như thế này
		numMap[v] = i
	}

	return []int{-1, -1}

}

// func main() {

// 	a := twoSumWithHashMap([]int{2, 7, 11, 5}, 9)
// 	b := twoSum([]int{3, 2, 4}, 6)
// 	c := twoSum([]int{3, 3}, 6)
// 	fmt.Println(a, b, c)
// }
