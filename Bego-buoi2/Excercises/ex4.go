package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type PERSON struct {
	name string
	job  string
	yob  int
}

func main() {

	f, err := os.Open("a.txt")

	if err != nil {
		log.Fatal(err)
	}

	team := make([]PERSON, 0)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		line := scanner.Text()

		temp := strings.Split(line, "|")
		//fmt.Println(temp)
		strings.ToUpper(temp[0])
		strings.ToLower(temp[1])
		year, _ := strconv.Atoi(temp[2])

		p := PERSON{
			strings.ToUpper(temp[0]),
			strings.ToLower(temp[1]),
			year,
		}
		team = append(team, p)
		//fmt.Println(p)
		//fmt.Println("\n")

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(team)
}
