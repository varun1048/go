package main

import (
	"fmt"
	"math"
)

func cal(x, y int) (i, j int) {
	i = x + y
	j = x - y
	return
}
func main() {
	num, num2, num3 := 9, 4, 6
	i := 1
	const con = 10
	for i <= 4 {
		out1, out2 := cal(con, i)
		fmt.Println(out1, out2)
		i++
	}

	var fl float64 = 13
	fmt.Printf("from math function: %.3f", math.Sqrt(fl))
	fmt.Println("\n", num, num2, num3)

	x, y := "thala", "thalapathy"
	change(x, y)
	fmt.Println(x,y)

}


func change(x, y string) string {
	var result string

	result = x
	x = y
	y = result

	return result

}
