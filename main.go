package main

import (
	"os"
	"bufio"
	"log"
	"fmt"
	"io"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		log.Fatal("Arguments must have a filename!")
		return
	}
	file, err := os.Open(args[1])
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewReader(file)
	var count_row int
	for {
		_, err := scanner.ReadString('\n')
		if err == io.EOF {
			break
		}
		count_row++
	}
	fmt.Printf("%d", count_row)
	return
}
