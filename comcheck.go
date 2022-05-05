package piscine

import (
	"fmt"
	"os"
)

func comcheck() {
	args := os.Args[1:]
	str := []string(args)
	for i := 0; i < len(str); i++ {
		if str[i] == "01" || str[i] == "galaxy" || str[i] == "galaxy01" {
			fmt.Println("Alert!!!")
			return
		}
	}
}
