package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

func testMkdir() {
	// 创建单极目录，如果已存在则err有错
	err := os.Mkdir("x2", 0755)
	fmt.Println(err)
	// 创建多级目录，已存在则什么都不做
	err1 := os.MkdirAll("x2/src/code/rest", 0755)
	fmt.Println(err1)
}
func testOtherMkdir() {
	// 创建单级目录
	if err := os.Mkdir("t2", 0755); err == nil {
		fmt.Println("文件夹已创建")
	} else {
		fmt.Println(err)
	}
}

func testDelFiles() {
	if err := os.Remove("t2"); err != nil {
		fmt.Println("删除失败", err)
	} else {
		fmt.Println("删除成功")
	}
	if err := os.RemoveAll("x2"); err != nil {
		fmt.Println("删除失败", err)

	} else {
		fmt.Println("删除成功")
	}
}

// 判断文件或文件夹是否存在
func testExistOrNot() {
	_, err := os.Stat("D:\\douz\\AppData\\Go_Pro\\src\\github.com\\study\\day16\\GoDirFiles")
	fmt.Println(err)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("文件或文件夹不存在")
		}
	} else {
		fmt.Println("存在")
	}
	// 是否是文件夹
	fileObject, _ := os.Stat("D:\\douz\\AppData\\Go_Pro\\src\\github.com\\study\\day16\\GoDirFiles")
	result := fileObject.IsDir()
	fmt.Println(result)
}

// 生成绝对路径
func TestAbsPath() {
	absPath, _ := filepath.Abs("t2/src")
	fmt.Println(absPath)
}

// 获取上级目录
func testGetLastDir() {
	absPath, _ := filepath.Abs("t2")
	fmt.Println(absPath)

	basePath := filepath.Dir(filepath.Dir(absPath))
	fmt.Println(basePath)
}

// 遍历目录(一级)
func testTraverseDir() {
	fileList, _ := ioutil.ReadDir("D:\\douz\\AppData\\Go_Pro\\src\\github.com\\study\\day16\\GoDirFiles")
	for _, obj := range fileList {
		fmt.Println(obj.Name())
	}
}

// 遍历目录(多级)
func testDirMore() {
	filepath.Walk("D:\\douz\\AppData\\Go_Pro\\src\\github.com\\study", func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			fmt.Println("文件有：", path)
		} else {
			fmt.Printf("文件夹有：%s\n", path)
		}
		return nil
	})
}

// 路径拼接
func testpathspace() {
	filePath := path.Join("v1", "v2", "v3/v4", "v6.exe")
	fmt.Println(filePath)
}

// 文件扩展名
func testFileExt() {
	ext := path.Ext("/xxx/xx/x.txt")
	fmt.Println(ext)
}

func main() {
	// testMkdir()
	// testOtherMkdir()
	// testDelFiles()
	// testExistOrNot()
	// TestAbsPath()
	// testGetLastDir()
	// testTraverseDir()
	// testDirMore()
	// testpathspace()
	testFileExt()
}
