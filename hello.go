package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	x := input("type a number")
	ar := strings.Split(x, " ")
	t := 0

	for _, v := range ar {
		n, er := strconv.Atoi(v)
		if er != nil {
			goto err
		}
		t += n
	}
	fmt.Printf("合計は %dです", t)
	return

err:
	fmt.Println("ERROR!")
}

func input(msg string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(msg + ": ")
	scanner.Scan()
	return scanner.Text()
}
