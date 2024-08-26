package main

import (
	"bufio"
	"calc/data"
	"errors"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	data.PrintCalculator()

	for {
		text, err := reader.ReadString('\n')
		processError(err)
		a, b, c := validate(text)
		calculate(a, b, c)
	}
}

func validate(text string) (first string, operator string, second string) {
	text = strings.TrimSpace(text)
	var inputs = strings.Split(text, " ")

	// проверка на кол-во аргументов
	if len(inputs) != 3 {
		processError(errors.New("неверное кол-во аргументов"))
	}

	// оба валидных арабских или  римских числа
	if !isUnified(inputs) {
		processError(errors.New("первое и второе число не унифицированы. Или введенные значения > 10"))
	}

	// проверка оператора
	if !isCorrectOpertor(inputs[1]) {
		processError(errors.New("указан неверный оператор"))
	}

	return inputs[0], inputs[1], inputs[2]
}

func calculate(first string, operator string, second string) {

	processOpertor := func(a int, operator string, b int) int {

		switch operator {
		case "+":
			return a + b
		case "-":
			return a - b
		case "/":
			{
				if b == 0 {
					processError(errors.New("нельзя делить на ноль"))
				}
				return a / b
			}

		case "*":
			return a * b
		default:
			processError(errors.New("указан неверный оператор"))
			return math.MinInt
		}
	}

	var a int
	var b int
	if isRoman(first) {
		a = data.RomanToArabic(first)
		b = data.RomanToArabic(second)
		res := processOpertor(a, operator, b)
		if processOpertor(a, operator, b) < 1 {
			processError(errors.New("результат работы с римскими цифрами может быть только положитлеьнео число"))
		}
		fmt.Println(data.ArabicToRoman(res))
	} else {
		first := strings.TrimSpace(first)
		second := strings.TrimSpace(second)
		a, _ := strconv.Atoi(first)
		b, _ := strconv.Atoi(second)
		fmt.Println(processOpertor(a, operator, b))
	}
}

func processError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}

func isUnified(input []string) bool {
	if (isRoman(input[0]) && isRoman(input[2])) || isArabic(input[0]) && isArabic(input[2]) {
		return true
	}
	return false
}

func isCorrectOpertor(input string) bool {
	switch input {
	case "/", "+", "-", "*":
		return true
	default:
		return false
	}

}

func isRoman(digit string) bool {
	// romanRegex := `^M{0,3}(CM|CD|D?C{0,3})(XC|XL|L?X{0,3})(IX|IV|V?I{0,3})$`
	digit = strings.TrimSpace(digit)
	romanRegex := `^(X|IX|IV|V?I{0,3})$`
	re := regexp.MustCompile(romanRegex)
	var result = re.MatchString(digit)
	return result
}

func isArabic(digit string) bool {
	digit = strings.TrimSpace(digit)
	number, err := strconv.ParseInt(digit, 10, 0)
	if err != nil {
		return false
	}
	if number > 10 {
		return false
	}
	return true
}

func DoNothig() {

}

// func Split(s, sep string) []string
// import "calc/data"
