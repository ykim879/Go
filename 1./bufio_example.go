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
		val := input.Text()
		if val == "e" {
			break
		}
		counts[val]++
	}
	//ignore input.Err()
	for line, n := range counts {
		fmt.Printf("count:%d\tinput.text():%s\n", n, line)
	}
}