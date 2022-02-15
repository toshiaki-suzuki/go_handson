package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	x := input("type 1~5")
	t := 0
	n, err := strconv.Atoi(x)
	if err != nil {
		return
	}
	switch n {
	case 5:
		t += 5
		fallthrough //一つ下のcaseも実行する
	case 4:
		t += 4
		fallthrough
	case 3:
		t += 3
		fallthrough
	case 2:
		t += 2
		fallthrough
	case 1:
		t += 1
	default:
		fmt.Println("範囲外です")
		return
	}
	fmt.Printf("合計は %dです", t)
}

func input(msg string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(msg + ": ")
	scanner.Scan()
	return scanner.Text()
}
