package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	modify := func(a []string, f func([]string) []string) []string {
		return f(a)
	}
	m := []string{
		"1st", "2nd", "3rd",
	}
	fmt.Println(m)
	m1 := modify(m, func([]string) []string {
		return append(m, m...)
	})
	fmt.Println(m1)
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
