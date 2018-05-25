package main

import (
	"bufio"
	"fmt"
	"os"
)

// Exercis e 1.4: Modify dup2 to print the names of all files in which each duplicated line occurs
func main() {
	counts := make(map[string]int)

	//todo, 不知道怎么写list和复杂数据结构...
	// Map<String, List<String>> dupString2File = new HashMap<>();
	// dupString2File := make(map[string]string)

	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Println("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
