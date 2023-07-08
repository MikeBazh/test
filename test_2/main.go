package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var Input1, Input2 string
var Numb1, Numb2 int
var e1 interface{} = nil
var e2 interface{} = nil
var rims1 bool = false                                                             //для проверки, является ли первое число римским
var rims2 bool = false                                                             //для проверки, является ли второе число римским
var err bool = false                                                               //для проверки, нарушен ли формат ввода
var rim = [10]string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}  //римские цифры (порядок "единиц")
var rim2 = [10]string{"X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC", "C"} //римские цифры (порядок "десятков")
// функция для вывода результата операции с римскими числами
func outRim(out int) string {
	var r, N1, N2 string
	if out < 10 {
		r = rim[out-1]
	}
	if out > 10 {
		N1 = rim2[(out/10)-1]
		if (out % 10) == 0 {
			N2 = ""
		}
		if (out % 10) != 0 {
			N2 = rim[(out%10)-1]
		}
		r = N1 + N2
	}
	return r
}

func main() {
	// считывание строки
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("введите выражение")
	text, _ := reader.ReadString('\n')
	data := strings.Fields(text)
	if len(data) > 3 {
		fmt.Println("Ошибка, формат операции не удовлетворяет заданию")
		err = true
	}
	if len(data) < 3 {
		fmt.Println("Ошибка, введите строку в формате 'число оператор число'")
		err = true
	}
	//Input1 := strings.TrimSpace(data[0])
	//Input2 := strings.TrimSpace(data[2])
	//проверка и считываение римских чисел
	if !err {
		Input1 = strings.TrimSpace(data[0])
		Input2 = strings.TrimSpace(data[2])
		for i := 1; i <= 10; i++ {
			switch rim[i-1] {
			case Input1:
				Numb1 = i
				if Input1 == Input2 {
					Numb2 = i
					rims2 = true
				}
				rims1 = true
			case Input2:
				Numb2 = i
				rims2 = true
			}
		}
	}
	//проверка, введены ли арабские числа
	_, e1 = strconv.Atoi(Input1)
	_, e2 = strconv.Atoi(Input2)
	if !err && (e1 != nil || e2 != nil) && !(rims1 && rims2) {
		fmt.Println("ошибка, введено неверное число")
		err = true
	}
	if rims1 && rims2 {
		switch data[1] {
		case "+":
			fmt.Println(data[0], "+", data[2], "=", outRim(Numb1+Numb2))
		case "-":
			if Numb1-Numb2 < 1 {
				fmt.Println("Ошибка операции вычитания римских чисел, результат меньше 1")
			}
			if Numb1-Numb2 >= 1 {
				fmt.Println(data[0], "-", data[2], "=", outRim(Numb1-Numb2))
			}
		case "*":
			fmt.Println(data[0], "*", data[2], "=", outRim(Numb1*Numb2))
		case "/":
			if Numb1/Numb2 < 1 {
				fmt.Println("Ошибка операции деления римских чисел, результат меньше 1")
			}
			if Numb1-Numb2 >= 0 {
				fmt.Println(data[0], "/", data[2], "=", outRim(Numb1/Numb2))
			}
		default:
			fmt.Println("Ошибка, неверный формат операции, введите строку в формате 'число оператор число'")
		}
	}
	if !err && (rims1 && (e2 == nil) || rims2 && (e1 == nil)) && !(rims1 && rims2) {
		fmt.Println("ошибка, калькулятор не работает с римскими и арабскими числами одновременно")
		err = true
	}
	if !err && !(rims1 && rims2) {
		Numb1, e1 = strconv.Atoi(Input1)
		Numb2, e2 = strconv.Atoi(Input2)
		if Numb1 <= 10 && Numb2 <= 10 && Numb1 > 0 && Numb2 > 0 {
			err = false
			switch data[1] {
			case "+":
				fmt.Println(data[0], "+", data[2], "=", Numb1+Numb2)
			case "-":
				fmt.Println(data[0], "-", data[2], "=", Numb1-Numb2)
			case "*":
				fmt.Println(data[0], "*", data[2], "=", Numb1*Numb2)
			case "/":
				fmt.Println(data[0], "/", data[2], "=", Numb1/Numb2)
			default:
				fmt.Println("Ошибка, неверный формат операции, введите строку в формате 'число оператор число'")
			}
		}
		if Numb1 > 10 || Numb2 > 10 || Numb1 <= 0 || Numb2 <= 0 {
			fmt.Println("Ошибка, калькулятор принимает на вход только числа от 1 до 10")
		}
	}
}
