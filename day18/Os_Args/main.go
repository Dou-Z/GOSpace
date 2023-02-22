package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("args[0]=", os.Args)
	if len(os.Args) > 0 {
		for ind, val := range os.Args {
			if ind == 0 {
				continue
			}
			fmt.Printf("os.args[%d]=%s\n", ind, val)
		}
	}
}
