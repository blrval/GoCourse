package main

import (
	"fmt"
	"math"
)	

func main()  {
	var leg1 float64 = 3  // катет1
	var leg2 float64 = 4  // катет2
	hypotenuse := math.Sqrt(math.Pow(leg1,2) + math.Pow(leg2,2))
	fmt.Println("Гипотенуза равна: ", hypotenuse)
	perimeter := leg1 + leg2 + hypotenuse
	fmt.Println("Перимерт треугольника: ", perimeter)
	area := leg1 * leg2 / 2 
	fmt.Println("Площадь треугольника: ", area)
}