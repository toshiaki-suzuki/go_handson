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
	// n, err := strconv.Atoi(x); ショートステイトメント
	// err == nil　条件
	n, err := strconv.Atoi(x)
	if fmt.Println("hoge"); err == nil {
		if n%2 == 0 {
			fmt.Println("偶数です")
		} else {
			fmt.Println("奇数です")
		}
	} else {
		fmt.Println("整数ではありません")
	}

}

func input(msg string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(msg + ": ")
	scanner.Scan()
	return scanner.Text()
}
