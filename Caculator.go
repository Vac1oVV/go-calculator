package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Карта для перевода римских чисел в арабские
var romanToInt = map[string]int{
	"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
	"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
}

// Карта для перевода арабских чисел в римские
var intToRoman = []string{
	"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X",
}

// Функция для проверки, является ли строка римским числом
func isRoman(input string) bool {
	_, exists := romanToInt[input]
	return exists
}

// Функция для выполнения арифметических операций
func calculate(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, errors.New("деление на ноль")
		}
		return a / b, nil
	default:
		return 0, errors.New("недопустимая операция")
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите выражение:")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	parts := strings.Fields(input)
	if len(parts) > 3 {
		panic("Выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
		return
	}
	if len(parts) != 3 {
		panic("Выдача паники, так как строка не является математической операцией.")
		return
	}

	aStr, op, bStr := parts[0], parts[1], parts[2]

	isRomanInput := isRoman(aStr) && isRoman(bStr)

	if (!isRoman(aStr) && isRoman(bStr)) || (isRoman(aStr) && !isRoman(bStr)) {
		panic("Выдача паники, так как используются одновременно разные системы счисления.")
		return
	}

	var a, b int
	var err error

	if isRomanInput {
		a = romanToInt[aStr]
		b = romanToInt[bStr]
	} else {
		a, err = strconv.Atoi(aStr)
		if err != nil {
			panic("Ошибка: неверный формат числа.")
			return
		}
		b, err = strconv.Atoi(bStr)
		if err != nil {
			panic("Ошибка: неверный формат числа.")
			return
		}
	}

	if a < 1 || a > 10 || b < 1 || b > 10 {
		panic("Ошибка: числа должны быть в диапазоне от 1 до 10.")
		return
	}

	result, err := calculate(a, b, op)
	if err != nil {
		panic(err)
		return
	}

	if isRomanInput {
		if result < 1 {
			panic("Выдача паники, так как в римской системе нет отрицательных чисел.")
			return
		}
		fmt.Println("Результат:", intToRoman[result])
	} else {
		fmt.Println("Результат:", result)
	}
}
