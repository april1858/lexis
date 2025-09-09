package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// построчное считываение данных из консоли
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		input += line
		if !strings.Contains(line, ";") {
			break
		}
	}

	fmt.Println(input)

	if err := scanner.Err(); err != nil {
		fmt.Println("[error]")
	}
}
