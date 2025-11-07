package utils

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	emptyStringError   = errors.New("Ошибка пустая строка")
	negativeNumber     = errors.New("Ошибка число меньше нуля")
	incorrectKindError = errors.New("Ошибка некорректный тип")
)

const (
	incomeKindRus      = "доход"
	incomeKindEng      = "income"
	expenditureKindRus = "расход"
	expenditureKindEng = "expenditure"
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
	fmt.Print(prompt)
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
	fmt.Print(prompt)
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

func ValidateString(s string) error {
	if len(s) == 0 {
		return emptyStringError
	}
	return nil
}

func ValidateInt64(n int64) error {
	if n < 0 {
		return negativeNumber
	}
	return nil
}

func ValidateFloat(f float64) error {
	if f < 0 {
		return negativeNumber
	}
	return nil
}

func ValidateKind(kind string) error {
	if kind != incomeKindRus && kind != expenditureKindRus && kind != incomeKindEng && kind != expenditureKindEng {
		return incorrectKindError
	}
	return nil
}
