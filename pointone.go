package piscine

import (
	"fmt"
	"piscine"
)

func PointOne(n *int) {
	*n = 1
}

func main() {
	n := 0
	piscine.Pointone(&n)
	fmt.Println(n)
}
