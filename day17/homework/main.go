package main

import "fmt"

var (
	coinsNum           = 50
	usersList []string = []string{"Matthew", "Sarah", "Augustus", "Heidi", "Emilie",
		"Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(usersList))
)

func calcCoin(username string) int {
	var sum = 0
	for _, char := range username {
		switch char {
		case 'a', 'A':
			sum = sum + 1
		case 'e', 'E':
			sum = sum + 1
		case 'i', 'I':
			sum = sum + 2
		case 'o', 'O':
			sum = sum + 3
		case 'u', 'U':
			sum = sum + 5
		}

	}
	return sum
}
func UserCoin() {
	for _, username := range usersList {
		fmt.Println(username)
		result := calcCoin(username)
		value, ok := distribution[username]
		if !ok {
			distribution[username] = result
		} else {
			fmt.Println(value)
		}
	}
	fmt.Println(distribution)
}

func main() {
	UserCoin()
}
