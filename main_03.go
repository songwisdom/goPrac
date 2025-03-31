package main

import "fmt"

func main03() {
	str := "chap03 - 연산자"
	fmt.Println(str)
	// 산술연산자
	// + -
	// 관계연산자
	// == != >=
	// 논리연산자
	// && || !

	// 할당연산자
	// = *= +=
	// >>= |= (??)
	// ( >> : 시프트연산, | : OR연산 )

	//포인터 연산자
	//변수의 주소를 얻거나
	//반대로 dereference(역참조) 할 때 사용
	//but Go에서는 포인터 산술 기능은 없음.

	k := 15
	kP := &k
	fmt.Println("k:", k, ",kP(var's val):", *kP)
	fmt.Println("kP(Pointer's val):", kP)

}
