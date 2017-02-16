package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		if input.Text() != "" {
			counts[input.Text()]++
		} else {
			break
		}
	}

	fmt.Println(counts)

}
