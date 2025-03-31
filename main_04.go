package main

import "fmt"

func main04() {
	str := "chap04 - 조건문"
	fmt.Println(str)

	//조건절 조건에 숫자(0,1) 못 씀, 무조건 boolean형
	//괄호 없어도 됨
	//대신 시작 브레이스가 조건절과 같은 라인에 있어야함
	k := 30 //case1) 조건절 밖 변수
	if k == 1 {
		fmt.Println("k is one")
	} else {
		fmt.Println("k isn't one")
	}

	//case2) 조건절 scope 안에서만 사용 변수
	if k2 := k * 3; k2 < k+10 {
		result := `<`
		fmt.Println(k2, result, k+10)
	} else {
		result := `>`
		fmt.Println(k2, result, k+10)
	}

	//switch
	//break 안써도 다음 case로 넘어가지 않음
	//switch 뒤에 expression이 없을 수 있음
	//case문에 expression을 쓸 수 있음
	var name string
	category := 4
	switch category {
	case 1:
		name = "Paper Book"
	case 2:
		name = "E Book"
	case 3:
		name = "NEWS"
	case 4:
		name = "other"
	}
	fmt.Println(name)
	//Type switch
	// name2 := "Demian"
	// switch name2.(type) {
	// case string:
	// 	fmt.Println("book name is string type. - ", name2)
	// case int:
	// 	fmt.Println("book name is int type. - ", name2)
	// default:
	// 	fmt.Println("unknown")
	// }
	//-> name2를 interface 타입으로 선언하라는데..뭔말이지
	var x int = 2
	check(x)
}

//break 안써도 다음 case로 넘어가지 않음 Go 컴파일러가 자동으로 break 문을 각 case문 블럭 마지막에 추가
func check(val int) {
	switch val {
	case 2:
		fmt.Println("2 이하")
		fallthrough
	case 1:
		fmt.Println("1 이하")
		fallthrough
	default:
		fmt.Println("default")
	}
}
