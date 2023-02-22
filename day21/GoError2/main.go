package main

import (
	"fmt"
	"os"
	"time"
)

type PathError struct {
	path      string
	op        string
	creatTime string
	message   string
}

func (p *PathError) Error() string {
	return fmt.Sprintf("path=%s \nop=%s \ncreatrTime=%s \nmessage=%s\n",
		p.path, p.op, p.creatTime, p.message)
}

func Open(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return &PathError{
			path:      filename,
			op:        "read",
			message:   err.Error(),
			creatTime: fmt.Sprint("%v", time.Now()),
		}
	}
	defer file.Close()
	return nil
}

func main() {
	err := Open("D:\\Desktop\\文件夹\\test.txt")
	switch v := err.(type) {
	case *PathError:
		fmt.Println("get path err:", v)
	default:

	}
}
