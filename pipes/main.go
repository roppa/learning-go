package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Start typing. Ctrl+C to quit")
	for scanner.Scan() {
		fmt.Println(strings.ToUpper(scanner.Text()))
	}
}
