package main

import (
	"fmt"
	"math"
)

func cal(x,y int)(i,j int){
	i=x+y
	j=x-y
return
}
func main() {
	i:=1;	
	const con =10
		for i<=4 {
			out1, out2  := cal(con,i); 
			fmt.Println(out1,out2)
			i++
		}
		
		var fl float64 = 13
		fmt.Printf("from math function: %.3f",math.Sqrt(fl))
		
}
