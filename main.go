package main

import "fmt"

func main() { //20250331
	str := `chap05 - 반복문`
	fmt.Println(str)

	//for 괄호 없음
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	//while과 비슷
	number := 10
	fmt.Print(sum, ", ")
	for sum > 0 {
		sum -= number
		number -= 1
		fmt.Print(sum, ", ")
	}

	//조건식 모두 생략하면 무한루프
	for sum < 5 {
		sum += 1
		fmt.Println("infinite~~")
	}
	//Go 컬렉션 타입
	//제네릭을 지원하지 않음음
	arrs := []string{"A", "B", "C"}
	for index, name := range arrs {
		fmt.Println(index, name)
	}
	//continue,goto,break
	a := 1
	for a < 15 {
		println("a:", a)
		if a == 5 {
			a += a
			println("a is 5:", a)
			continue // for루프 시작으로
		}
		a++
		if a > 10 {
			break //루프 빠져나옴
		}
	}
	if a == 11 {
		goto END //goto 사용예
	}
END:
	println("[End] a:", a)
	//break 레이블

}
