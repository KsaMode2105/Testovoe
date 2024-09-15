package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var ToArab = map[string]int{
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
var ToRome = map[int]string{
	1:   "I",
	4:   "IV",
	5:   "V",
	9:   "IX",
	10:  "X",
	40:  "XL",
	50:  "L",
	90:  "XC",
	100: "C",
}

const (
	Incorrect     = "Некорректное число"
	IncorrectRome = "Некорректное римское число"
	Mnogo         = "формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)."
	raznie        = "Используются одновременно разные системы счисления"
	MinusRome     = "Результатом работы калькулятора с римскими числами могут быть только положительные числа"
	Unknown       = "Не соответствует ни одной из представленных арифметических операций"
)

func ConvToRome(number int) string {
	roman := ""
	if number == 0 {
		return roman
	} // цикл для преобразования арабских в римские
	// первый цикл для
	for _, d := range []int{100, 90, 50, 40, 10, 9, 5, 4, 1} {
		for number >= d {
			roman += ToRome[d]
			number -= d
		}
	}
	return roman
}
func ConvToArab(RomanNumber string) (int, error) {
	if znach, ok := ToArab[RomanNumber]; ok {
		return znach, nil
	}
	return 0, errors.New("Некорректное римское число")
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	input := strings.Split(text, " ")
	//fmt.Println(input)
	if len(input) != 3 {
		panic(Mnogo)
	}
	operand1 := input[0]
	operator := input[1]
	operand2 := input[2]

	IsRoman := false
	// пробуем преобразовать в арабские
	Numeral1, err1 := strconv.Atoi(operand1)
	Numeral2, err2 := strconv.Atoi(operand2)
	// если оба числа не арабские проверим являются ли они римскими
	if err1 != nil && err2 != nil {
		Numeral1, err1 = ConvToArab(operand1)
		Numeral2, err2 = ConvToArab(operand2)
		if err1 != nil || err2 != nil {
			panic(Incorrect)
		}
		//если все хорошо меняем флаг
		IsRoman = true
	} else if err1 == nil && err2 == nil {
		//если операнды прошли проверку на арабские числа
	} else {
		panic(raznie)
	}
	var res int
	switch operator {
	case "+":
		res = Numeral1 + Numeral2
	case "-":
		res = Numeral1 - Numeral2
	case "*":
		res = Numeral1 * Numeral2
	case "/":
		res = Numeral1 / Numeral2
	default:
		panic(Unknown)
	}
	if IsRoman {
		// если были римские числа
		if res < 1 {
			panic(MinusRome)
		}
		RomanRes := ConvToRome(res) //преобразуем
		fmt.Println(RomanRes)
	} else {
		//если использовались арабские
		fmt.Println(res)
	}

}
