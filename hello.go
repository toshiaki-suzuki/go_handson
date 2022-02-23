package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	m := map[string]int{ //
		"a": 100,
		"b": 200,
		"c": 300,
	}

	m["total"] = m["a"] + m["b"] + m["c"]
	delete(m, "a")
	fmt.Println(m)
}

func input(msg string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(msg + ": ")
	scanner.Scan()
	return scanner.Text()
}
