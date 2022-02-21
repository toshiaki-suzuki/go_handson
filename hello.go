package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	a := [5]int{10, 20, 30, 40, 50} // 配列
	b := a[:3]                      // 配列から取り出したスライス
	c := []int{1, 2, 3, 4}          // スライス
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	c = append(c, 50)
	fmt.Println(c)
	d := append([]int{0}, c...)
	fmt.Println(d)
}

func input(msg string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(msg + ": ")
	scanner.Scan()
	return scanner.Text()
}
