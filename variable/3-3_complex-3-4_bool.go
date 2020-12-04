package main

import (
	"fmt"
	"math/cmplx"
)

func basic()  {
	var x complex128 = complex(1, 2) // 1+2i
	var y complex128 = complex(3, 4) // 3+4i
	z := 1 + 2i

	fmt.Println(x * y + z) // (-4+12i)
	fmt.Println(real(x * y + z)) // -4
	fmt.Println(imag(x * y + z)) // 12
	fmt.Println(cmplx.Sqrt(-1)) // (0+1i)
}

func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

func main()  {
	basic()
}
