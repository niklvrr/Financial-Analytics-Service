package utils

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func AskString(in *bufio.Reader, prompt string) (string, error) {
	fmt.Print(prompt)
	line, err := in.ReadString('\n')
	if err != nil {
		return "", err
	}

	line = strings.TrimSpace(line)
	return line, nil
}

func AskInt(in *bufio.Reader, prompt string) (int, error) {
	fmt.Println(prompt)
	line, err := in.ReadString('\n')
	if err != nil {
		return 0, err
	}

	line = strings.TrimSpace(line)
	num, err := strconv.Atoi(line)
	if err != nil {
		return 0, err
	}

	return num, nil
}

func AskFloat(in *bufio.Reader, prompt string) (float64, error) {
	fmt.Println(prompt)
	line, err := in.ReadString('\n')
	if err != nil {
		return 0, err
	}

	line = strings.TrimSpace(line)
	num, err := strconv.ParseFloat(line, 64)
	if err != nil {
		return 0, err
	}

	return num, nil
}
