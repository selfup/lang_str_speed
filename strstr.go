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
	check(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "OK db=") {
			split := strings.Split(line, "OK ")
			lines = append(lines, split[1])
		}
	}

	check(scanner.Err())
	fmt.Println(lines[0:40])
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
