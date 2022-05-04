package main

import (
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) >= 1 {
		for i := 0; i <= len(args)-1; i++ {
			file, err := os.ReadFile(args[i])
			if err != nil {
				os.Stderr.WriteString("ERROR: open " + string(args[i]) + ":no such file or directory\n")
				os.Exit(1)
			}
			os.Stderr.WriteString(string(file))
		}
	} else {
		str, _ := ioutil.ReadAll(os.Stdin)
		os.Stderr.WriteString(string(str))
	}
}
