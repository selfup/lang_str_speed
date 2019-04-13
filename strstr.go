package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	lines := []string{}

	file, err := os.Open("tmp/logs.log")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	// doesn't do great with lines with more than 65536 characters
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) > 2 {
			if strings.Contains(line, "OK db=") {
				split := strings.Split(line, "OK ")
				lines = append(lines, split[1])
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(lines[0:40])
}
