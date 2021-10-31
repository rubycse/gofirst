package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var name string = "Ruby"

func main() {
	writeToFile("hello.txt", "Hello World!")
}

func writeToFile(fileName string, content string) {
	file, er := os.Create("myFile")
	checkError(er)
	length, err := file.WriteString("My name is " + name)
	checkError(err)
	fmt.Println(length)
	defer file.Close()

	//file, err := os.Create("./" + fileName)
	//checkError(err)
	//length, err := io.WriteString(file, content)
	//checkError(err)
	//fmt.Println("length: ", length)
	//defer file.Close()
}

func checkError(err error) {
	if (err != nil) {
		panic(err)
	}
}

func structPrint() {
	b := book{
		name:   "Refactoring",
		author: "Martin Fawler",
	}

	fmt.Println(b)
	b.print(true)
	fmt.Println(b)
}

type book struct {
	name   string
	author string
}

func (b book) print(withAuthor bool) {
	if withAuthor {
		b.name = b.name + " by " + b.author
	}

	fmt.Println(b.name)
}
func mapCreate() {
	m := make(map[string]int)
	m["hi"] = 1
	m["hello"] = 2
	fmt.Println(m)

	for k, v := range m {
		fmt.Println(k, v)
	}
}

func read() {
	reader := bufio.NewReader(os.Stdin)
	s1, _ := reader.ReadString('\n')
	s2, _ := reader.ReadString('\n')

	num1, e1 := strconv.ParseInt(strings.TrimSpace(s1), 10, 64)
	num2, e2 := strconv.ParseInt(strings.TrimSpace(s2), 10, 64)

	if e1 == nil && e2 == nil {
		fmt.Printf("Sum: %v", num1+num2)
	}
}

func add(num1 int, num2 int) int {
	return num1 + num2
}

func find1to3(num int) (lower, upper string) {
	fmt.Println(name)
	switch num {
	case 1:
		return "one", "One"
	case 2:
		return "two", "Two"
	case 3:
		return "three", "Three"
	default:
		return "none", "None"
	}
}
