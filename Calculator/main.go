package main

import (
	"fmt"
	"strings"
)

func toRomanian(num int) string {
	val := []int{100, 90, 50, 40, 30, 20, 10, 5, 1}
	sysm := []string{"C", "XC", "LXXX", "LXX", "LX", "L", "XL", "XXX", "XX", "X", "V", "I"}

	romanNumeral := ""
	i := 0

	for num > 0 {
		for num >= val[i] {
			romanNumeral += sysm[i]
			num -= val[i]
		}
		i++
	}

	return romanNumeral
}

// Блок калькулятора
func calculate(expression string) (int, error) {
	parts := strings.Fields(expression)
	if len(parts) != 3 {
		return 0, fmt.Errorf("Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).") // Проверяет количество оперантов в строке если их больше 3х возвращает ошибку
	}

	a, operator, b := parts[0], parts[1], parts[2] // Определяет где что находится в полученной строке
	aInt := 0
	bInt := 0

	if operator != "+", "-", "/", "*" {
		return 0, fmt.Errorf("Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
	}

	if !strings.ContainsAny(a, "CXLVI") && !strings.ContainsAny(b, "CXLVI") {
		aInt = parseInt(a) //Принимает две строки а и b и проверяет их на условие есть ли там символы CXLVI или нет
		bInt = parseInt(a) //в зависимотсти от полученного результата выдаст ошибку.
	} else if strings.ContainsAny(a, "CXLVI") && strings.ContainsAny(b, "CXLVI") {
		return 0, fmt.Errorf("Вывод ошибки, так как используются одновременно разные системы счисления.")
	} else {
		return 0, fmt.Errorf("Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
	}

	result := 0
	switch operator {
	case "+":
		result = aInt + bInt
	case "*":
		result = aInt * bInt
	case "-":
		result = aInt - bInt
	case "/":
		if bInt == 0 {
			return 0, fmt.Errorf("Делить на ноль нельяза")
		}
		result = aInt / bInt
	}

	return result, nil
}

func parseInt(s string) int {
	arabicNumerals := map[rune]int{'I': 1, 'V': 5, 'X': 10}
	result := 0

	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && arabicNumerals[rune(s[i])] < arabicNumerals[rune(s[i+1])] {
			result -= arabicNumerals[rune(s[i])]
		} else {
			result += arabicNumerals[rune(s[i])]
		}
	}
	return result
}

func main() {
	var expression string
	fmt.Println("Формат операции  — два операнда и один оператор (+, -, /, *).")
	fmt.Scanln(&expression)

	result, err := calculate(expression)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Result: %v\n", result)
	}
}
