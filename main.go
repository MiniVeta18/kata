package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func romanToInt(roman string) (int, error) {
	if roman == "" {
		return 0, fmt.Errorf("invalid roman numeral: %s", roman)
	}

	mapping := map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10,
	}

	if val, ok := mapping[roman]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("invalid roman numeral: %s", roman)
}

func intToRoman(n int) string {
	if n == 0 {
		return "N"
	}

	if n < 0 {
		return strings.Trim("-"+intToRoman(-n), "N")
	}

	if n >= 1000 {
		return strings.Trim("M"+intToRoman(n-1000), "N")
	}

	if n >= 900 {
		return strings.Trim("CM"+intToRoman(n-900), "N")
	}

	if n >= 500 {
		return strings.Trim("D"+intToRoman(n-500), "N")
	}

	if n >= 400 {
		return strings.Trim("CD"+intToRoman(n-400), "N")
	}

	if n >= 100 {
		return strings.Trim("C"+intToRoman(n-100), "N")
	}

	if n >= 90 {
		return strings.Trim("XC"+intToRoman(n-90), "N")
	}

	if n >= 50 {
		return strings.Trim("L"+intToRoman(n-50), "N")
	}

	if n >= 40 {
		return strings.Trim("XL"+intToRoman(n-40), "N")
	}

	if n >= 10 {
		return strings.Trim("X"+intToRoman(n-10), "N")
	}

	if n >= 9 {
		return strings.Trim("IX"+intToRoman(n-9), "N")
	}

	if n >= 5 {
		return strings.Trim("V"+intToRoman(n-5), "N")
	}

	if n >= 4 {
		return strings.Trim("IV"+intToRoman(n-4), "N")
	}

	return strings.Trim("I"+intToRoman(n-1), "N")
}

func calculate(a, b int, operator string) (int, error) {
	if a > 10 || b > 10 {
		fmt.Println("Число больше 10, разрешено вводить значение не превышающее 10")
		os.Exit(1)
	}
	switch operator {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, fmt.Errorf("division by zero")
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("invalid operator: %s", operator)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var a, b int
	var operator string
	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	input = strings.Trim(input, " \n\r")

	splited := strings.Split(input, " ")
	if len(splited) != 3 {
		panic("invalid input")
	}
	operator = splited[1]
	if !strings.ContainsAny(operator, "+-/*") {
		panic("invalid operator " + operator)
	}
	isRomanic := strings.ContainsAny(splited[0], "IVX") && strings.ContainsAny(splited[2], "IVX")
	isArabic := strings.ContainsAny(splited[0], "0123456789") && strings.ContainsAny(splited[2], "0123456789")
	if !isRomanic && !isArabic {
		panic("invalid input")
	}
	var result int
	if isRomanic {
		// Roman numerals
		a, err = romanToInt(splited[0])
		if err != nil {
			panic(err)
		}
		b, err = romanToInt(splited[2])
		if err != nil {
			panic(err)
		}
		result, err = calculate(a, b, operator)
		if err != nil {
			panic(err)
		}
		romanResult := intToRoman(result)
		fmt.Println(romanResult)
	} else {
		// Arabic numerals
		a, err = strconv.Atoi(splited[0])
		if err != nil {
			panic(err)
		}
		b, err = strconv.Atoi(splited[2])
		if err != nil {
			panic(err)
		}
		result, err = calculate(a, b, operator)
		if err != nil {
			panic(err)
		}
		fmt.Println(result)
	}
}
