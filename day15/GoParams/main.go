package main

import (
	"flag"
	"fmt"
)

func main() {
	// fmt.Println(os.Args)
	host := flag.String("h", "127.0.0.1", "主机IP")
	port := flag.Int("p", 3306, "端口")

	var email string
	flag.StringVar(&email, "e", "", "邮箱")

	flag.Parse()

	fmt.Println(*port, *host, email)
}
