package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type branch struct {
	name       string
	Parentname string
	left       string
	right      string
}

var stuct []branch

func main() {
	fmt.Printf("kurwa")
	for i := 0; i < 3; i++ {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("branch defenition:Parentname;name;left;right:\n")
		text, _ := reader.ReadString('\n')
		fmt.Println(text)
		fmt.Print(stuct)
		stuct = append(stuct, branch{name: "A", left: "B"})
	}
	parse("nya")

}
func parse(kurwa string) {
	s := strings.Split("127.0.0.1:5432", ":")
    ip, port := s[0], s[1]
    fmt.Println(ip, port)
}
