package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	x := input("type a number")
	n, err := strconv.Atoi(x)
	t := 0
	c := 1
	if err != nil {
		goto err // gotoは変数宣言の後に書かないとコンパイルエラーが出る
	}
	for c <= n {
		t += c
		c++
	}
	fmt.Printf("合計は %dです", t)

err:
	fmt.Println("ERROR!")
}

func input(msg string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(msg + ": ")
	scanner.Scan()
	return scanner.Text()
}
