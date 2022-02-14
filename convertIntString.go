package convertIntString

import (
	"fmt"
	"strconv"
)

func main() {
	i, err := strconv.Atoi("3")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(i)
	a := strconv.Itoa(3)
	fmt.Println(a)
}
