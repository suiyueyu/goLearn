package exercise

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("1. print the command")
	fmt.Println("the command is: " + os.Args[0])

	fmt.Println("2. output the index and value")
	for index, val := range os.Args[1:] {
		fmt.Println(string(index) + ": " + val)
	}

	fmt.Println("3. test the execute effiecient between strings.Join and your code")
	// emmm... todo
}
