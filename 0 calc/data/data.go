package data

import "fmt"

var RomanDigits = [5]string{"X", "IX", "V", "IV", "I"}

func ArabicToRoman(num int) string {

	keyValues := []int{10, 9, 5, 4, 1}

	romanResult := ""
	for i := 0; i < len(RomanDigits); i++ {
		for num >= keyValues[i] {
			num -= keyValues[i]
			romanResult += RomanDigits[i]
		}
	}
	return romanResult
}

func RomanToArabic(s string) int {
	// Map to hold Roman numeral values
	romanMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
	}

	// Variable to store the final result
	result := 0

	// Iterate through the Roman numeral string
	for i := 0; i < len(s); i++ {
		// Get the value of the current Roman numeral character
		value := romanMap[s[i]]

		// If we're not at the last character and the next Roman numeral is greater,
		// subtract the current value (handling subtractive combinations)
		if i+1 < len(s) && romanMap[s[i]] < romanMap[s[i+1]] {
			result -= value
		} else {
			// Otherwise, add the value to the result
			result += value
		}
	}

	return result
}

var calculator = [][]string{
	{" ", "_", "_", "_", "_", "_", "_", "_", "_", "_", "_", "_", "_", "_", "_", "_", "_", "_", "_", "_", "_", "_", " "},
	{"|", " ", " ", "_", "_", "_", "_", "_", "_", "_", "_", "_", "_", "_", "_", "_", "_", "_", "_", "_", " ", " ", "|"},
	{"|", " ", "|", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", "0", ".", " ", "|", " ", "|"},
	{"|", " ", "|", "_", "_", "_", "_", "_", "_", "_", "_", "_", "_", "_", "_", "_", "_", "_", "_", "_", "|", " ", "|"},
	{"|", " ", " ", "_", "_", "_", " ", "_", "_", "_", " ", "_", "_", "_", " ", " ", " ", "_", "_", "_", " ", " ", "|"},
	{"|", " ", "|", " ", "7", " ", "|", " ", "8", " ", "|", " ", "9", " ", "|", " ", "|", " ", "+", " ", "|", " ", "|"},
	{"|", " ", "|", "_", "_", "_", "|", "_", "_", "_", "|", "_", "_", "_", "|", " ", "|", "_", "_", "_", "|", " ", "|"},
	{"|", " ", "|", " ", "4", " ", "|", " ", "5", " ", "|", " ", "6", " ", "|", " ", "|", " ", "-", " ", "|", " ", "|"},
	{"|", " ", "|", "_", "_", "_", "|", "_", "_", "_", "|", "_", "_", "_", "|", " ", "|", "_", "_", "_", "|", " ", "|"},
	{"|", " ", "|", " ", "1", " ", "|", " ", "2", " ", "|", " ", "3", " ", "|", " ", "|", " ", "x", " ", "|", " ", "|"},
	{"|", " ", "|", "_", "_", "_", "|", "_", "_", "_", "|", "_", "_", "_", "|", " ", "|", "_", "_", "_", "|", " ", "|"},
	{"|", " ", "|", " ", ".", " ", "|", " ", "0", " ", "|", " ", "=", " ", "|", " ", "|", " ", "/", " ", "|", " ", "|"},
	{"|", " ", "|", "_", "_", "_", "|", "_", "_", "_", "|", "_", "_", "_", "|", " ", "|", "_", "_", "_", "|", " ", "|"},
	{"|", "_", "_", "_", "_", "_", "_", "_", "_", "_", "_", "_", "_", "_", "_", "_", "_", "_", "_", "_", "_", "_", "|"},
}

var calcAscii = []string{
	" _____       _        _           ",
	"/  __ \\     | |      | |          ",
	"| /  \\/ __ _| | ___  | |__  _   _ ",
	"| |    / _` | |/ __| | '_ \\| | | |",
	"| \\__/\\ (_| | | (__  | |_) | |_| |",
	" \\____/\\__,_|_|\\___| |_.__/ \\__, |",
	"                             __/ |",
	"                            |___/ ",
}
var dmiseDev = [6]string{
	"     _           _               _            ",
	"    | |         (_)             | |           ",
	"  __| |_ __ ___  _ ___  ___   __| | _____   __",
	" / _` | '_ ` _ \\| / __|/ _ \\ / _` |/ _ \\ \\ / /",
	"| (_| | | | | | | \\__ \\  __/| (_| |  __/\\ V /  ",
	" \\__,_|_| |_| |_|_|___/\\___(_)__,_|\\___| \\_/  ",
}

func PrintCalculator() {

	//fmt.Println("         Calcultor designed by")
	for _, line := range calcAscii {
		fmt.Println(line)
	}
	for _, line := range dmiseDev {
		fmt.Println(line)
	}

	for _, line := range calculator {
		for _, char := range line {
			fmt.Print(char)
		}
		fmt.Println()
	}
}
