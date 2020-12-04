package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func helloWorld()  {
	for i, r := range "Hello, 世界" {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}
}

func strangeString()  {
	fmt.Println(string(123456))
}

func basename(s string) string {
	slash := strings.LastIndex(s, "/") // -1 if "/" not found
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}


func changeIntString()  {
	// strconv.Itoa
	x := 123
	y := fmt.Sprintf("%d", x)
	fmt.Println(y, strconv.Itoa(x)) // 123 123

	// strconv.FormatInt
	fmt.Println(strconv.FormatInt(int64(x), 2)) // 1111011

	// Atoi
	z, _ := strconv.Atoi("123")
	fmt.Println(z) // 123

	// ParseInt
	w, _ := strconv.ParseInt("123", 10, 64) // 기수 10, 촤대 64비트
	fmt.Println(w) // 123
}

func main()  {
	helloWorld()
	strangeString()
	fmt.Println(basename("hihi.go"))
	fmt.Println(comma("12345.123"))
	changeIntString()
}
