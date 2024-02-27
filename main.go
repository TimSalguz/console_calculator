package main

import "fmt"
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

	return result
}

func main() {
    rom_number := rom_num_to_int("MCMLIV")
    fmt.Println(rom_number)
}