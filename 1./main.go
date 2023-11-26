package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Printf("The program will test between strings.Join vs concatenationg. The array it will perfrom is os.Args which is below")
	fmt.Println(os.Args[0])
	for i, v := range os.Args[1:] {
		fmt.Printf("\nIndex: %d, Value: %s", i, v)
	}

	fmt.Printf("\n testing started \n")
	s := ""
	startTime := time.Now()
	for _, v := range os.Args[1:] {
		s += v
	}
	endTime := time.Now()
	fmt.Printf("Elapsed Time: %s for concatenating %s \n", endTime.Sub(startTime), s)

	startTime2 := time.Now()
	s2 := strings.Join(os.Args[1:], "")
	endTime2 := time.Now()
	fmt.Printf("Elapsed Time: %s for %s \n", endTime2.Sub(startTime2), s2)
}