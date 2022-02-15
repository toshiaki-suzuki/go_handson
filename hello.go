package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	x := input("type a number")
	fmt.Println(x + "は")
	n, err := strconv.Atoi(x)
	switch fmt.Println("ショートステートメント"); err == nil {
	case n%2 == 0:
		fmt.Println("偶数です")
	case n%2 != 0:
		fmt.Println("奇数です")
	default:
		fmt.Println("整数ではありません")
	}
}

func input(msg string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(msg + ": ")
	scanner.Scan()
	return scanner.Text()
}
