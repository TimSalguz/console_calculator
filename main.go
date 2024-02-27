package main

import "fmt"
import "strings"
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
    int_from_rom_num := rom_num_to_int("XXIV")
    fmt.Println(int_from_rom_num)
    rom_num_from_int := int_to_rom_num(int_from_rom_num)
    fmt.Println(rom_num_from_int)
}