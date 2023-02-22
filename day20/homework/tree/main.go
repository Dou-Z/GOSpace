package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/urfave/cli"
)

func LisDir(dirPath string, deep int) (err error) {
	dir, err := ioutil.ReadDir(dirPath)
	// fmt.Println(dir, err)

	if err != nil {
		return err
	}
	if deep == 1 {
		// fmt.Println(filepath.Base(dirPath))
		fmt.Printf("|---%s\n", filepath.Base(dirPath))

	}
	sep := string(os.PathSeparator)
	// fmt.Println("sep:", sep)
	for _, fi := range dir {
		// 如果是目录，继续调用Lisdir进行遍历
		// fmt.Println("fi:", fi.IsDir())

		if fi.IsDir() {
			fmt.Printf("|")
			for i := 0; i < deep; i++ {
				fmt.Printf("    |")
			}
			fmt.Printf("----%s\n", fi.Name())
			LisDir(dirPath+sep+fi.Name(), deep+1)
			continue
		}
		fmt.Printf("|")
		for i := 0; i < deep; i++ {
			fmt.Printf("    |")
		}
		fmt.Printf("----%s\n", fi.Name())
	}
	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "tree"

	app.Usage = "list all file"
	app.Action = func(c *cli.Context) error {
		var dir string = "."
		if c.NArg() > 0 {
			dir = c.Args().First()
		}
		LisDir(dir, 1)
		return nil
	}
	app.Run(os.Args)
}
