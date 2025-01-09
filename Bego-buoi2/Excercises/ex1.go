package main

import "fmt"

type Person struct {
	name      string
	job       string
	birthYear int
}

func (p *Person) calAge() int {
	return 2024 - p.birthYear
}

func (p *Person) checkSuitable() bool {
	if len(p.name) == 0 {
		return false
	}
	fmt.Println(len(p.name))
	temp := 2000 % len(p.name)

	return temp == 0
}

// func main() {

// 	p := Person{
// 		"Ha", "AI engineer", 200,
// 	}

// 	fmt.Println(p.calAge())

// 	fmt.Println("Does this person suitable for this job: ", p.checkSuitable())

// 	fmt.Println(stringProcess("abx"))

// }
