package main

import "fmt"

func main01() {
	st := "chap01 - Go 변수와 상수"
	fmt.Println(st)

	//변수
	var a int
	var f float32 = 11.
	var b int
	a = 5
	b = 10
	fmt.Println(a, f, b)
	//변수 선언 순서, 자료형 종류, 선언했으면 사용 필수(강타입 언어)
	var a1, b1, c1 int = 10, 20, 30
	fmt.Println(a1, b1, c1)
	i := "string" //Short Assignment Statement ( := ) - var 생략 가능

	fmt.Println(i)

	//상수 - 변수와의 차이 : 변하지 않음
	const c int = 100
	fmt.Println("const c:", c)
	c2 := 101
	fmt.Println("const c:", c2)

	const (
		gl  = "Golang"
		jvm = "java virtual machine"
		jre = "java runtime environment"
	)

	const (
		apple = iota //걍 순회
		banana
		coconut
	)
	fmt.Println(apple, banana, coconut)

	//예약 키워드
	// break default func interface select
	// case defer go map struct chan else
	// goto package switch const fallthrough
	// if range type continue for import return var
}
