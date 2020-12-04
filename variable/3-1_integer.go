package main

import "fmt"

func funcVar1()  {
	var a int32 = 1
	var b int = 2

	// 변환을 해줘야 한다.
	// c = a + b 는 컴파일 에러
	var c = a + int32(b)
	fmt.Println(c)
}

func funcOverflow()  {
	var u uint8 = 255
	fmt.Println(u, u+1, u*u) // 255 0 1

	var i int8 = 127
	fmt.Println(i, i+1, i*i) // 127 -128 1
}

func bitLeftTest()  {
	// 왼쪽 시프트 - 부호 없음
	var x uint8 = 1<<1 | 1<<7
	var y uint8 = 1<<1 | 1<<2
	fmt.Printf("%08b\n", x)
	fmt.Printf("%08b\n", y)

	// 왼쪽 시프트 - 부호 있음
	var z int8 = 1<<1 | 1<<6 // 에러 발생
	var w int8 = 1<<1 | 1<<2

	fmt.Printf("%08b\n", z)
	fmt.Printf("%08b\n", w)
}

func bitRightTest()  {
	// 오른쪽 시프트 - 부호 없음
	var x uint8 = 64>>2
	fmt.Printf("%08b\n", x)

	// 오른쪽 시프트 - 부호 있음
	var y int8 = -64>>2
	fmt.Printf("%08b\n", y)
	fmt.Printf("%d\n", y)
	fmt.Printf("%08b\n", -64)
}

func medalTest()  {
	medals := []string{"gold", "silver", "bronze"}
	for i := len(medals) - 1; i >= 0; i-- {
		fmt.Println(medals[i])
	}
}

func formattingTest()  {
	o := 0666
	fmt.Printf("%d %[1]o, %#[1]o\n", o)
}


func main()  {
	//funcVar1()
	//funcOverflow()

	//bitLeftTest()
	//bitRightTest()
	//medalTest()
	//formattingTest()
}
