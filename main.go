package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

/*
golang лексический парсинг строки логического выражения
*/

func main() {
	// построчное считываение данных из консоли
	var input string
	vars := make(map[string]string)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.Contains(line, ";") {
			input = strings.ReplaceAll(line, " ", "")
			break
		}
		line = line[:len(line)-1]
		v := strings.Split(line, "=")
		if v[1] == "" {
			v[1] = "False"
		}
		val := strings.ToLower(v[1])
		vars[v[0]] = val
	}

	fmt.Println(input)

	re := regexp.MustCompile(`[\w]+`) // \w+ находит одну или более букв, цифр или нижнее подчеркивание

	tokens := re.FindAllString(input, -1)

	fmt.Println("Токены:")
	for _, token := range tokens {
		fmt.Println(token)
	}

	fmt.Println(vars)

	fmt.Println(operations(input))

	if err := scanner.Err(); err != nil {
		fmt.Println("[error]")
	}
}

func operations(line string) []string {
	op := make([]string, 0)
	if strings.Contains(line, ")not(") {
		op = append(op, "not")
	}
	if strings.Contains(line, ")and(") {
		op = append(op, "and")
	}
	if strings.Contains(line, ")or(") {
		op = append(op, "or")
	}
	if strings.Contains(line, ")xor(") {
		op = append(op, "xor")
	}
	return op
}
