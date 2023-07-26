package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	romanToArabic := map[string]int{
		"I":        1,
		"II":       2,
		"III":      3,
		"IV":       4,
		"V":        5,
		"VI":       6,
		"VII":      7,
		"VIII":     8,
		"IX":       9,
		"X":        10,
		"XI":       11,
		"XII":      12,
		"XIII":     13,
		"XIV":      14,
		"XV":       15,
		"XVI":      16,
		"XVII":     17,
		"XVIII":    18,
		"XIX":      19,
		"XX":       20,
		"XXI":      21,
		"XXII":     22,
		"XXIII":    23,
		"XXIV":     24,
		"XXV":      25,
		"XXVI":     26,
		"XXVII":    27,
		"XXVIII":   28,
		"XXIX":     29,
		"XXX":      30,
		"XXXI":     31,
		"XXXII":    32,
		"XXXIII":   33,
		"XXXIV":    34,
		"XXXV":     35,
		"XXXVI":    36,
		"XXXVII":   37,
		"XXXVIII":  38,
		"XXXIX":    39,
		"XL":       40,
		"XLI":      41,
		"XLII":     42,
		"XLIII":    43,
		"XLIV":     44,
		"XLV":      45,
		"XLVI":     46,
		"XLVII":    47,
		"XLVIII":   48,
		"XLIX":     49,
		"L":        50,
		"LI":       51,
		"LII":      52,
		"LIII":     53,
		"LIV":      54,
		"LV":       55,
		"LVI":      56,
		"LVII":     57,
		"LVIII":    58,
		"LIX":      59,
		"LX":       60,
		"LXI":      61,
		"LXII":     62,
		"LXIII":    63,
		"LXIV":     64,
		"LXV":      65,
		"LXVI":     66,
		"LXVII":    67,
		"LXVIII":   68,
		"LXIX":     69,
		"LXX":      70,
		"LXXI":     71,
		"LXXII":    72,
		"LXXIII":   73,
		"LXXIV":    74,
		"LXXV":     75,
		"LXXVI":    76,
		"LXXVII":   77,
		"LXXVIII":  78,
		"LXXIX":    79,
		"LXXX":     80,
		"LXXXI":    81,
		"LXXXII":   82,
		"LXXXIII":  83,
		"LXXXIV":   84,
		"LXXXV":    85,
		"LXXXVI":   86,
		"LXXXVII":  87,
		"LXXXVIII": 88,
		"LXXXIX":   89,
		"XC":       90,
		"XCI":      91,
		"XCII":     92,
		"XCIII":    93,
		"XCIV":     94,
		"XCV":      95,
		"XCVI":     96,
		"XCVII":    97,
		"XCVIII":   98,
		"XCIX":     99,
		"C":        100,
	}
	stringToint := map[string]int{
		"0":  0,
		"1":  1,
		"2":  2,
		"3":  3,
		"4":  4,
		"5":  5,
		"6":  6,
		"7":  7,
		"8":  8,
		"9":  9,
		"10": 10,
	}
	fmt.Print("Введите выражение:")
	reader := bufio.NewReader(os.Stdin)
	operation, _ := reader.ReadString('\n')
	words := strings.Fields(operation)
	if len(words) > 3 {
		fmt.Println("формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)")
	} else if len(words) < 3 {
		fmt.Println("строка не является математической операцией.")
	}
	if value1, inMap := romanToArabic[words[0]]; inMap {
		if value2, inMap := romanToArabic[words[2]]; inMap {
			number1 := value1
			number2 := value2
			var result int
			switch words[1] {
			case "+":
				result = number1 + number2
			case "-":
				result = number1 - number2
			case "/":
				result = number1 / number2
			case "*":
				result = number1 * number2
			default:
				fmt.Println("Введите правильный символ операции")
			}
			if result < 0 {
				fmt.Println("В римской системе нет отрицательных чисел")
			} else {
				var romanNum string
				for key, value := range romanToArabic {
					if value == result {
						romanNum = key

					}
				}
				fmt.Printf(romanNum)
			}
		} else {
			fmt.Println("используются одновременно разные системы счисления")
		}

	} else {
		if value1, inMap := stringToint[words[0]]; inMap {
			if value2, inMap := stringToint[words[2]]; inMap {
				number1 := value1
				number2 := value2
				var result int
				switch words[1] {
				case "+":
					result = number1 + number2
				case "-":
					result = number1 - number2
				case "/":
					result = number1 / number2
				case "*":
					result = number1 * number2
				default:
					fmt.Println("Введите правильный символ операции")
				}
				fmt.Println(result)
			} else {
				fmt.Println("используются одновременно разные системы счисления")
			}
		}
	}
}
