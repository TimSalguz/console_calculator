package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//Функция преобразует строку римского числа в целое.
//Она не учитывает некоторые особенности римских чисел, и не является на 100% правильной!!!
func rom_num_to_int(rom_num string) int {
	romans := []byte{'I', 'V', 'X', 'L', 'C', 'D', 'M'}
	values := []int{1, 5, 10, 50, 100, 500, 1000}

	result := 0
	prev_value := 0

	for i := len(rom_num) - 1; i >= 0; i-- {
		value := 0
		for index, roman := range romans {
			if rom_num[i] == roman {
				value = values[index]
				break
			}
		}

		if value < prev_value {
			result -= value
		} else {
			result += value
		}
		prev_value = value
	}
	if int_to_rom_num(result) != rom_num{
        panic("Введено некорректное римское число!");
    }

	return result
}

//Функция преобразует целое число в римское.
func int_to_rom_num(num int) string {
	if num <= 0 || num > 3999 {
        panic("Недопустимое число");
	}

	romanSymbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	romanValues := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	var result strings.Builder

	for i := 0; i < len(romanSymbols); i++ {
		for num >= romanValues[i] {
			result.WriteString(romanSymbols[i])
			num -= romanValues[i]
		}
	}

	return result.String()
}

func main() {
	fmt.Println("Введите математическую операцию. Между числами и операторами должны быть пробелы.")
	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')

    var num1_is_roman, num2_is_roman bool

    parts := strings.Fields(input)
    if len(parts) != 3 {
        fmt.Println(len(parts));
		panic("Некорректный ввод! Пожалуйста, используйте формат 'число операция число'. Например, '1 + 20'")
	}
    num1, err := strconv.Atoi(parts[0])
	if err != nil {
        num1 = rom_num_to_int(parts[0])
        num1_is_roman = true;
	}

	num2, err := strconv.Atoi(parts[2])
	if err != nil {
        num2 = rom_num_to_int(parts[2])
        num2_is_roman = true;
	}
    if !(num1 <= 10 && num2 <= 10){
        panic("Числа должны быть меньше или равны 10!")
    }
    answer_is_romanian := false;
    if num1_is_roman && num2_is_roman{
        answer_is_romanian = true;
    } else if !num1_is_roman && !num2_is_roman{
            answer_is_romanian = false;
    } else{
        panic("Римскими должны быть или все числа, или ни одно!");
    }

	result := 0

	switch parts[1] {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		result = num1 / num2
	default:
		panic("Некорректная операция. Пожалуйста, используйте '+', '-', '*' или '/'")
	}

    if answer_is_romanian{
        if result <= 0{
            panic("Ответ в римских числах не может быть <= 0!")
        }
        fmt.Println(int_to_rom_num(result))
    } else{
        fmt.Println(result)
    }
}