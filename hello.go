package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	data := "*新しい値*"
	m1 := modify(data)
	data = "+new data+"
	m2 := modify(data)

	fmt.Println(m1())
	fmt.Println(m2())
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

func modify(d string) func() []string {
	m := []string{
		"1st", "2nd",
	}
	return func() []string {
		return append(m, d)
	}
}
