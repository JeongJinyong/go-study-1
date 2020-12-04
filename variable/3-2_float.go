package main

import (
	"fmt"
	"math"
)

func check1()  {
	fmt.Println(math.MaxFloat32)
	fmt.Println(math.MaxFloat64)
}

func nanTest()  {
	var z float64
	fmt.Println(z, -z, 1/z, -1/z, z/z)

	nan := math.NaN()
	fmt.Println(nan == nan, nan < nan, nan> nan)
}

func div1() (value float64, ok bool)  {
	if value == 0 {
		return 0, false
	}
	return 11 / value, true
}

func main()  {
	//check1()
	//nanTest()
}