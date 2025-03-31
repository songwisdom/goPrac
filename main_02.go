package main

import "fmt"

func main02() {
	st := "chap02 - Go 데이타타입"
	fmt.Println(st)

	//데이타 타입이 정말 다양함
	//bool string(immutable; 수정불가)
	//int int8 int16 int32 int64
	//uint uint8 uint16 uint32 uint64 uintptr
	//float32 float64 complex64 complex128
	//byte(uint8과 동일, 바이트코드에 사용)
	//rune(int32와 동일, 유니코드 코드포인트)

	//문자열
	// ``(back quote) = raw string literal, 탈출문자 적용x, 복수라인 o
	// "" interpreted string literal, 탈출문자O, 복수라인X && ""+""

	rawLiteral := `메롱
	롱메 \n 안먹음?ㅋ`
	interLiteral := "나는" + "너를" + "좋아해" //+할때도 다음라인에 + 안됨
	fmt.Println(rawLiteral + "\n" + interLiteral)

	//타입 변환
	//(강타입 언어지만 암묵적 변환이 불가능한 거고, 명시적 지정만 하면 가능능)
	i := 100
	u := uint(i)
	fmt.Println(i, "(int)->(uint)", u)

	//ex) 문자열을 바이트배열로 변환
	strBefore := "ABC"
	strAfter := []byte(strBefore)
	fmt.Println(strBefore, "->", strAfter)
	strFinal := string(strAfter)
	fmt.Println("toString: ", strFinal)

}
