# 기본 테이터 타입

---

## 정수
  ### 1. 종류
  - 네 가지 크기(8, 16, 32, 64 비트) 존재
    - 부호화된 타입: `int8`, `int16`, `int32`, `int64`
    - 부호 없는 타입: `uint8`, `uint16`, `uint32`, `uint64`
  - `rune` 타입: (=uint8) 유니코드 코드 포인트
  - `uintptr` : 포인터 값의 모든 비트를 저장(저수준 프로그래밍에서 사용)
  - `int` 는 32 비트이지만 명시적으로는 `int32` 와 다르다.
  
  ```go
  var a int32 = 1
  var b int = 2

  // c = a + b 는 컴파일 에러
  var c = a + int32(b) // 다음과 같이 타입 변환 필요
  fmt.Println(c)
  ```
  
  ### 2. 이항 연산자
  - "*, /, %, <<, >>, &, &^, +, -" 등 여러가지 존재
  - 산술 연산
    - +, -, *, /, %
    - `%` 의 경우 정수에서만 사용 가능하며 부호는 피제수와 같다 (-5%2=-3, -5%-2=-3)
    - __오버플로우__ 가 발생하지 않도록 주의

  ```go
  // 오버플로우 예시
  var u uint8 = 255
  fmt.Println(u, u+1, u*u) // 255 0 1

  var i int8 = 127
  fmt.Println(i, i+1, i*i) // 127 -128 1
  ```

  - 비교 표현식
    - ==, >=, <= 등
    - 결과는 불리언 타입
  - 단항 연산자 (+, -)
    - +x (=0+x)
    - -x (=0-x)
  - 비트 단위 이항 연산자
    - &(AND), |(OR), ^(XOR), &^(AND NOT), <<(왼쪽 시프트), >>(오른쪽 시프트)
    - 처음 네 가지는 부호 개념이 없는 비트 패턴으로 취급
    - `^` 의 경우 비트-> XOR, 단항연산자->단위 부정 혹은 보수 연산
    - 부호없는 숫자의 경우 `<<, >>` 는 빈 비트를 0으로 채우지만, 부호있는 숫자를 `>>` 하면 빈 비트를 부호비트의 복사본으로 채운다. 따라서 정수를 비트패턴으로 사용한다면 부호 없는 산술 연산을 사용해야 한다. (_어떤 문제가 있을까????_)

  ```go
  // 오른쪽 시프트 - 부호 없음
  var y uint8 = 64>>2
  fmt.Printf("%08b\n", y) // 00010000
 
  // 오른쪽 시프트 - 부호 있음
  var w int8 = -64>>2
  fmt.Printf("%08b\n", -64) // -0010000 (11100000)
  fmt.Printf("%08b\n", z) // 00000010
  fmt.Printf("%08b\n", w) // -1000000 (11000000) -> (01100000 에서 부호비트 1 을 채움)
  fmt.Printf("%d\n", w) // -16
  ```

  - 부호 없는 숫자를 쓰면 안되는 경우도 존재 (uint 대신 int)
    - 아래와 같은 경우 uint 를 사용하면 i 의 최솟값은 0 (0-1 = 255) 이므로 `medals[255]` 가 되면서 panic 발생 

  ```go
  medals := []string{"gold", "silver", "bronze"}
  for i := len(medals) - 1; i >= 0; i-- {
    fmt.Println(medals[i])
  }
  ```

  ### 3. 타입 변환
  - 타입 변환을 명시적으로 해주어야 한다. 

  ```go
  var apples int32 = 1
  var oranges int16 = 2
  var compote int = apples + oranges // 컴파일 오류
  var compote2 int = int(apples) + int(oranges) // 타입 변환 필요
  ```

  - 부동 소수점의 경우 소수부분을 절삭
  - 피연산자가 대상 타립을 벗어나는 경우 그 동작이 구현에 의존 

  ```go
  f := 3.141
  fmt.Println(f, int(f)) // 3.131, 3

  f := 1e100
  i := int(f) // 구현별로 다름
  ```

  ### 4. 출력 포멧
  - `%d`: 10진수, `%o`: 8진수, `%x`: 16 진수
  - `#` 을 붙히면 0, 0x, 0X 등의 접두사를 붙혀준다.

  ```go
	o := 0666
	fmt.Printf("%d %[1]o, %#[1]o\n", o) // 438 666, 0666
  ```

---

## 부동소수점 수

  ### 1. 종류
  - float32, float64 존재
  - 최댓값: math.MaxFloat32(3.4e38), math.MaxFloat64(1.8e308)

  ### 2. NaN
  - 0/0 이나 Sqrt(-1) 과 같이 숫자가 아닌 값
  - NaN 의 비교 결과는 항상 __false__

  ```go
	var z float64
	fmt.Println(z, -z, 1/z, -1/z, z/z) // 0 -0 +Inf -Inf NaN

	nan := math.NaN()
	fmt.Println(nan == nan, nan < nan, nan> nan) // false false false
  ```

  - 따라서 실패가 예상된다면 실패를 별도로 만들도록 한다.

  ```go
  func div() (value float64, ok bool)  {
    if value == 0 { // 0으로 나눌 경우 실패!
      return 0, false
    }
    return 11 / value, true
  }
  ```

---

## 복소수
  ### 1. 종류
  - complex64, complex128 존재
  - `x := 1 + 2i` 와 같이 표현 가능
  - real, imag 함수는 각각 실수와 허수를 추출
  - math/cmplx 패키지에는 복소수 제곱근이나 지수 함수 등 처리하는 라이브러리 함수 존재

  ```go
  var x complex128 = complex(1, 2) // 1+2i
  var y complex128 = complex(3, 4) // 3+4i
  z := 1 + 2i

  fmt.Println(x * y + z) // (-4+12i)
  fmt.Println(real(x * y + z)) // -4
  fmt.Println(imag(x * y + z)) // 12
  fmt.Println(cmplx.Sqrt(-1)) // (0+1i)
  ```

---

## 불리언 (bool)
  ### 1. 기본
  - true, false 두가지 허용
  - &&, || 로 결합이 가능하며 단축연산 적용
    - 왼쪽 피연산자 값에 의해 답이 결정된 경우 오른쪽 계산하지 않는다.

  ```go
  s != "" && s[0] == 'x' // s[0] 에 빈 문자열을 적용하지 않으므로 패닉발생을 막을 수 있다.
  ```

  - && 는 || 보다 우선순위가 높다 (아래에 괄호 필요 없다.)
  
  ```go
  if 'a' <= c && c <= 'z' || 
      'A' <= c && c <= 'Z' || 
      '0' <= c && c <= '9' || 
      // ... 아스키 문자나 숫자
  ```

  - 묵시적으로 0 이나 1 로 변환되지 않으며 그 반대도 동일
    - 자주 필요하다면 함수를 작성해서 사용하자
  
  ```go
  func btoi(b bool) int {
    if b {
      return 1
    }
    return 0
  }
  ```

---

## 문자열
  ### 0. 기본
  - 통상적으로 유니코드 포인트(rune) 을 UTF-8 로 인코딩한 시퀀스로 해석
  - len 함수: 문자열 바이트 수 반환
  - s[i] 문자열의 i 번째 바이트 반환
  - 아스키코드가 아닌 문자열의 UTC-8 인코딩에는 2개 이상의 바이트 필요 (문자열의 i 번째 바이트가 i 번째 문자열 아님)
  - s[i:j]: i 번째 인덱스에서 j-1 번째 인덱스의 바이트 담은 새 문자열 생성
    - i 혹은 j 생략 가능하며 이때는 0 혹은 len(s) 로 간주
  - 비교 연산자 (<, ==) 사용가능, 비교는 사전순
  - 문자열 내의 데이터를 직접 변경하는 건 불가
    - `s[0] = 'L'` 과 같은 작업 불가능 

  ### 1. 문자열 리터럴
  - 큰따옴표로 묶인 바인트 시퀀스
  - 유니코드 코드 포인트 삽입 가능
  - escape sequence(\n 등) 사용 가능
  - 원시 문자열은 큰따옴표 대신 백쿼드 "\`" 사용
    - 정규식, HTML 템플릿, JSON 리터럴, 명령 사용법 메시지 등에 사용

  ```go
  const GoUsage = `Go is a tool for managing Go source code.

  Usage:
      go command [argument]
  `
  ```

  ### 2. 유니코드
  - `rune` 으로 이용 (int32)
  - 각 유니코드 코드 포인트의 인코딩은 모두 동일한 32 비트

  ### 3. UTF-8
  - 켄 톰슴, 롭 파이크가 발명
  - 현재 유니코드의 표준
    - Go 에서도 권장
  - unicode 패키지에 개별 룬 처리를 위한 함수 존재 
    - unicode/utf9 패키지에 룬을 바이트로 인코딩하거나 디코딩 하는 함수 존재

  ```
  0xxxxxxx                              rune 0-127 (아스키 -> 상위비트가 0)
  110xxxxx 10xxxxxx                     128-2047
  1110xxxx 10xxxxxx 10xxxxxx            2048-65535
  11110xxx 10xxxxxx 10xxxxxx 10xxxxxx   65536-0x10ffff (4바이트까지 이용)
  ```

  - 여러 문자열 연산에서 디코딩이 필요없게 만들어줌

  ```go
  s := "Hello, 世界"
  fmt.Println(len(s)) // "13"
  fmt.Println(utf8.RuneCountInString(s)) // "9"
  // 바이트수와 길이는 다를 수 있다!!
  ```

  - `range` 루프를 사용하면 묵지적으로 UTF-8 디코딩을 수행하여 글자 수만큼 반복한다. 

  ```go
  for i, r := range "Hello, 世界" {
    fmt.Println("%d\t%q\t%d\n", i, r, r)
  }
  ```

  ```
  0	'H'	72
  1	'e'	101
  2	'l'	108
  3	'l'	108
  4	'o'	111
  5	','	44
  6	' '	32
  7	'世'	19990
  10	'界'	30028
  ```

  - `utf8.DecodeRuneInString()` 은 예상치 못한 입력 바이트를 받으면 대체문자 `\uFFFD` 를 생성하며 이는 검은색 안에 ? 로 표시된다. (raange 도 포함) 
  - `[]rune` 변환을 적용하면 인코딩된 문자열의 유니코드 코드 포인트 시퀀스 반환
  - 룬의 슬라이스 문자열 변환 가능 (`string(rune)`)
    - _95p 뭐지?? 대체문자와 ? 표??_


  ### 4. 문자열과 바이트 슬라이스
  - `strings` 패키지: 문자열 검색, 교체, 비교 잘라내기, 쪼개기 합치기 등
  - `byte` 패키지: []byte 타입의 바이트 슬라이스 조작
  - `strconv` 패키지: 불리언, 부동소수점 수 값을 문자열 값으로 변환, 역변환하는 함수와 문자열을 참조하거나 역참조하는 함수 존재
  - `unicode` 패키지: rune 을 분류하기 위한 `IsDigit`, `IsLetter`, `IsUpper` 등 존재


  ```go
  // strings.LastIndex 활용
  func basename(s string) string {
    slash := strings.LastIndex(s, "/") // -1 if "/" not found
    s = s[slash+1:]
    if dot := strings.LastIndex(s, "."); dot >= 0 {
      s = s[:dot]
    }
    return s
  }
  ```

  ```go
  // comma 붙이기
  func comma(s string) string {
    n := len(s)
    if n <= 3 {
      return s
    }
    return comma(s[:n-3]) + "," + s[n-3:]
  }
  ```

  - 문자열은 바이트 슬라이스로 변환하고 되돌릴 수 있다.
    - 문자열을 변경 불가능하지만, 바이트 슬라이스의 원소는 변경 가능
  - `[]byte(s)` 변환은 s 바이트의 복사본을 갖는 새 바이트 배열을 할당하고 이 배열의 요소를 참조하는 슬라이스를 산출
  - `string(b)` 또한 복사가 일어남

  ```go
  s := "abc"
  b := []byte(s)
  s2 := string(b)
  ```

  - 불필요한 메뫼리 할당을 피하기 위해 특정한 유틸리티 함수들은 strings 함수와 직접 대응한다.

  ```
  func Contains(s, substr string) bool   <->   func Contains(b, subslice []byte]) bool
  func Count(s, sep string) int   <->   func Count(b, sep []byte]) int
  func Fields(s string) []string   <->   func Contains(s []byte]) [][]byte
  
  ...
  ```
  - bytes 패키지는 슬라이스의 효율적인 조작을 위한 `Buffer` 타입 제공

  ```go
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
  ```

  ### 5. 문자열과 숫자 사이의 변환
  - 숫자와 문자열 표현 사이의 변환이 필요한 경우 `strconv` 패키지를 이용 가능 
    - 정수 -> 문자열 변환은 `Itoa` 활용
    - 다른 기수를 가진 숫자를 포매팅 할때는 `FormatInt, FormatUint` 등 사용
    - 정수 파싱은 `Atoi` 나 `ParseInt` 를 활용

  ```go
	// strconv.Itoa
	x := 123
	y := fmt.Sprintf("%d", x)15
	fmt.Println(y, strconv.Itoa(x)) // 123 123

	// strconv.FormatInt
	fmt.Println(strconv.FormatInt(int64(x), 2)) // 1111011

	// strconv.Atoi
	z, error := strconv.Atoi("123")
	fmt.Println(z) // 123

	// strconv.ParseInt
	w, error := strconv.ParseInt("123", 10, 64) // 기수 10, 촤대 64비트 
	fmt.Println(w) // 123
  ```

---

## 상수
  ### 0. 기본
  - 컴파일 시 평가되는 표현 (변하지 않는 값)
  - `const` 로 상수 선언
  - 한번 선언에 여러 상수 표현 가능
    - 그룹 선언 시 첫 번째르 제외하고는 오른쪽 표현식 생략 가능
  - 상수에 대한 연산 결과도 상수
    - len, cap real 등의 결과도 상수

  ```go
  const pi = 3.14159
  const (
    e = 2.718281828
    pi2 = 3.14159
    d // 1 로 초기화된다.
  )
  ```

  ### 1. 상수 생성기 iota
  - `iota` 를 활용해 연관된 값들을 하나하나 명시하지 않고 생성 가능
  - 0 부터 시작
  - enum 타입처럼 활용 가능 

  ```go
  type Weekday int

  const (
    Sunday Weekday = iota // 0
    Mondy // 1
    Tuesday // 2
    Wednesday // ....
    Thursday
    Friday
    Saturaday
  )
  ```

  - net 패키지에서는 5비트의 부호 없는 정수에 까각 고유한 이름과 불리언 값 부여

  ```go
  type Flags uint

  const (
    FlagUp          Flags = 1 << iota // interface is up
    FlagBroadcast                      // interface supports broadcast access capability
    FlagLoopback                       // interface is a loopback interface
    FlagPointToPoint                   // interface belongs to a point-to-point link
    FlagMulticast                      // interface supports multicast access capability
  )
  ```

  - 단 `iota` 는 지수 연산자는 존재하지 않는다. (10 단위로 증가시키지는 못한다.)

  ### 2. 타입 없는 상수
  - 특정 타립으로 지정되지 않은 상수 표현 가능 (6가지)
    - 타입없는 불리언 (x = true)
    - 타입없는 정수 (x=111)
    - 타업없는 룬
    - 타입없는 부동소수점 (x = 1.1)
    - 타입없는 복소수 (x = 1i)
    - 타입없는 문자열 존재
  - 변수에 타입 없이 지정할 경우 묵지적으로 해당 타입으로 변환한다.
  - 타입없는 정수는 크기가 지정되지 않은 `int` 로 변환되지만 타입없는 부동소수점 수와 복소수는 크기가 지정된 `float64`, `complex128` 로 변환된다.

  ```go
  const ex1 = 123123123123123123123

  var f float64 = 3 + 0i // 타입없는 복소수 -> float64 변환
  f = 2 // 타입없는 정수 -> float64 변환
  ```

---