package main

import (
    "fmt"
    "strings"
    "strconv"
)
//Функция преобразует строку римского числа в целое.
//Она не учитывает некоторые особенности римских чисел, и не является на 100% правильной!!!
func rom_num_to_int(rom_num string) int {
	result := 0
    var prev_value byte

	for i := 0; i < len(rom_num); i++ {
		value := rom_num[i]
		switch value {
		case 'I':
			result += 1
		case 'V':
			if prev_value == 'I' {
				result += 3 // 5 - 1 = 4
			} else {
				result += 5
			}
		case 'X':
			if prev_value == 'I' {
				result += 8 // 10 - 1 = 9
			} else {
				result += 10
			}
		case 'L':
			if prev_value == 'X' {
				result += 30 // 50 - 10 = 40
			} else {
				result += 50
			}
		case 'C':
			if prev_value == 'X' {
				result += 80 // 100 - 10 = 90
			} else {
				result += 100
			}
		case 'D':
			if prev_value == 'C' {
				result += 300 // 500 - 100 = 400
			} else {
				result += 500
			}
		case 'M':
			if prev_value == 'C' {
				result += 800 // 1000 - 100 = 900
			} else {
				result += 1000
			}
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
    //var input string
    input := "X - IX"
    var num1_is_romanian, num2_is_romanian bool
	//fmt.Print("Введите арифметическую опеарцию (разделяйте числа и арифметические операторы пробелами!): ")
	//fmt.Scanln(&input)
    parts := strings.Fields(input)
    if len(parts) != 3 {
        fmt.Println(len(parts));
		panic("Некорректный ввод! Пожалуйста, используйте формат 'число операция число'. Например, '1 + 20'")
	}
    num1, err := strconv.Atoi(parts[0])
	if err != nil {
        num1 = rom_num_to_int(parts[0])
        num1_is_romanian = true;
	}

	num2, err := strconv.Atoi(parts[2])
	if err != nil {
        num2 = rom_num_to_int(parts[2])
        num2_is_romanian = true;
	}
    if !(num1 <= 10 && num2 <= 10){
        panic("Числа должны быть меньше или равны 10!")
    }
    answer_is_romanian := false;
    if num1_is_romanian && num2_is_romanian{
        answer_is_romanian = true;
    } else if !num1_is_romanian && !num2_is_romanian{
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


    // int_from_rom_num := rom_num_to_int(input)
    // fmt.Println(int_from_rom_num)
    // rom_num_from_int := int_to_rom_num(int_from_rom_num)
    // fmt.Println(rom_num_from_int)
}