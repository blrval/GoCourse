package main

import (
	"fmt"
)	

func converter()  {
	const currencies float64 = 64.4
	var rub float64
	//var dollar float64 = rub/currencies
	fmt.Println("Введите сумму в рублях для конвертации в доллары.")
	fmt.Scanln(&rub)
	fmt.Println("Конвертируемая сумма в долларах: ", rub/currencies)
}


