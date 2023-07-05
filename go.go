package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	var rimNum = map[string](int){
		"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10, "XI": 11, "XII": 12, "XIII": 13, "XIV": 14, "XV": 15, "XVI": 16, "XVII": 17, "XVIII": 18, "XIX": 19, "XX": 20, "XXI": 21, "XXII": 22, "XXIII": 23, "XXIV": 24, "XXV": 25, "XXVI": 26, "XXVII": 27, "XXVIII": 28, "XXIX": 29, "XXX": 30, "XXXI": 31, "XXXII": 32, "XXXIII": 33, "XXXIV": 34, "XXXV": 35, "XXXVI": 36, "XXXVII": 37, "XXXVIII": 38, "XXXIX": 39, "XL": 40, "XLI": 41, "XLII": 42, "XLIII": 43, "XLIV": 44, "XLV": 45, "XLVI": 46, "XLVII": 47, "XLVIII": 48, "XLIX": 49, "L": 50, "LI": 51, "LII": 52, "LIII": 53, "LIV": 54, "LV": 55, "LVI": 56, "LVII": 57, "LVIII": 58, "LIX": 59, "LX": 60, "LXI": 61, "LXII": 62, "LXIII": 63, "LXIV": 64, "LXV": 65, "LXVI": 66, "LXVII": 67, "LXVIII": 68, "LXIX": 69, "LXX": 70, "LXXI": 71, "LXXII": 72, "LXXIII": 73, "LXXIV": 74, "LXXV": 75, "LXXVI": 76, "LXXVII": 77, "LXXVIII": 78, "LXXIX": 79, "LXXX": 80, "LXXXI": 81, "LXXXII": 82, "LXXXIII": 83, "LXXXIV": 84, "LXXXV": 85, "LXXXVI": 86, "LXXXVII": 87, "LXXXVIII": 88, "LXXXIX": 89, "XC": 90, "XCI": 91, "XCII": 92, "XCIII": 93, "XCIV": 94, "XCV": 95, "XCVI": 96, "XCVII": 97, "XCVIII": 98, "XCIX": 99, "C": 100,
	}

	var isRomanNum1 bool = false
	var isRomanNum2 bool = false
	var rimResult string
	var error1 = errors.New("Возникла ошибка:")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите математическое выражение")
	equation, _ := reader.ReadString('\n')
	equation = strings.TrimSpace(equation)
	components := strings.Split(equation, " ")

	if len(components) != 3 {
		fmt.Println(error1, "неправильный формат уравнения!")
		return
	}

	operand1 := strings.TrimSpace(components[0])
	operand2 := strings.TrimSpace(components[2])
	operator := components[1]

	operand1Int, err1 := strconv.Atoi(components[0])
	operand2Int, err2 := strconv.Atoi(components[2])

	if err1 != nil {
		operand1Int = rimNum[operand1]
		if operand1Int == 0 {
			fmt.Println(error1, "необходимо вводить только римские или арабские числа")
			return
		} else {
			isRomanNum1 = true
		}
	}

	if err2 != nil {
		operand2Int = rimNum[operand2]
		if operand2Int == 0 {
			fmt.Println(error1, "необходимо вводить только римские или арабские числа")
			return
		} else {
			isRomanNum2 = true
		}
	}

	if operand1Int < 1 || operand1Int > 10 || operand2Int < 1 || operand2Int > 10 {
		fmt.Println(error1, "вводимые операнды не могут быть меньше 1 и больше 10!")
		return
	}

	var result int

	switch operator {
	case "+":
		result = operand1Int + operand2Int
	case "-":
		result = operand1Int - operand2Int
	case "*":
		result = operand1Int * operand2Int
	case "/":
		result = operand1Int / operand2Int

	default:
		fmt.Println(error1, "неверный оператор!")
		return
	}

	if isRomanNum1 && isRomanNum2 {
		if result > 0 {
			for key, value := range rimNum {
				if value == result {
					rimResult = key
					fmt.Println("Результат: ", rimResult)
					return
				}
			}
		} else {
			fmt.Println(error1, "римские цифры не могут быть равны или меньше 0")
		}
	} else if (isRomanNum1 == true && isRomanNum2 == false) || (isRomanNum1 == false && isRomanNum2 == true) {
		fmt.Println(error1, "либо только римские либо только арабские")
	} else {
		fmt.Println("Результат: ", result)
	}
}
