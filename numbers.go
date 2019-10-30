package main

import "fmt"

func main() {
	var number int64
	fmt.Println("Введите целое число: ")
	fmt.Scanln(&number)
	if number%2 == 0 {
		fmt.Println(number, "Это четное число")
	} else {
		fmt.Println(number, "Это нечетное число")
	}
	second()
	printFib()
}

func second() {
	var number int64
	fmt.Println("Введите целое число: ")
	fmt.Scanln(&number)
	if number % 3 == 0 {
		fmt.Println(number, "Это число делится на 3 без остатка")
	} else {
		fmt.Println(number, "Это число не делится на 3 без остатка")
	}
}

func fibi(n int) int {
    var a, b int = 0, 1
    for i := 0; i < n; i++ {
        a, b = b, a+b
    }
    return a
}

func printFib() {
	fmt.Println("Первые 100 чисел Фибоначчи: ")
    for i := 0; i < 100; i++ {
        fmt.Printf("%d, ", fibi(i))
    }
    fmt.Println("")
}