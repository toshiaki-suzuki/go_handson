package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	m := []string{}
	m, _ = push(m, "apple")
	m, _ = push(m, "banana")
	m, _ = push(m, "orange")
	fmt.Println(m)
	m, v := pop(m)
	fmt.Println("get "+v+" ->", m)
}

func input(msg string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(msg + ": ")
	scanner.Scan()
	return scanner.Text()
}
func push(a []string, v string) ([]string, int) {
	return append(a, v), len(a)
}

func pop(a []string) ([]string, string) {
	return a[:len(a)-1], a[len(a)-1]
}
