package main

import (
	"bufio"
	"fmt"
	"os"
	// "github.com/urfave/cli"
)

func getInput() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	return reader.ReadString('\n')
}
func transPostExpress(express string) (postExpress []string, err error) {
	var opStack Stack
	var i int
LABEL:
	for i < len(express) {
		switch {
		case express[i] >= '0' && express[i] <= '9':
			var number []byte
			for ; i < len(express); i++ {
				if express[i] < '0' || express[i] > '9' {
					break
				}
				number = append(number, express[i])

			}
			postExpress = append(postExpress, string(number))
		case express[i] == '+' || express[i] == '-' || express[i] == '*' || express[i] == '/':
			if opStack.Empty() {
				opStack.Push(fmt.Sprintf("%c", express[i]))
				i++
				continue LABEL
			}
			data, _ := opStack.Top()
			if data[0] == '(' || data[0] == ')' {
				opStack.Push(fmt.Sprintf("%c", express[i]))
				i++
				continue LABEL
			}
			if (express[i] == '+' || express[i] == '-') ||
				((express[i] == '*' || express[i] == '/') && (data[0] == '/')) {
				postExpress = append(postExpress, data)
				opStack.Pop()
				opStack.Push(fmt.Sprintf("%c", express[i]))
				i++
				continue LABEL
			}
			opStack.Push(fmt.Sprintf("%c", express[i]))
			i++
		case express[i] == '(':
			opStack.Push(fmt.Sprintf("%c", express[i]))
			i++
		case express[i] == ')':
			for !opStack.Empty() {
				data, _ := opStack.Pop()
				if data[0] == '(' {
					break
				}
				postExpress = append(postExpress, data)
			}
			i++
		default:

		}
	}
}
func main() {

}
