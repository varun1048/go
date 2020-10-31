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
	next := closure()
		for i <= 4 {
			out1, out2 := cal(con, i)
			fmt.Println(out1, out2)
			
			fmt.Println(next())		
		i++
				
		}
	
	
	var fl float64 = 13
	fmt.Printf("from math function: %.3f", math.Sqrt(fl))
	fmt.Println("\n", num, num2, num3)

	x, y := "thala", "thalapathy"
	change(&x, &y)
	fmt.Println(x,y)

}
func closure() func() int {
   i:=0
   return func() int {
      i+=1
	  return i  
   }
}
func change(x, y *string) {
	var result string

	result = *x
	*x = *y
	*y = result

	//return result

}
