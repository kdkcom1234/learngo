// main: 컴파일러가 컴파일하는 패키지
package main

import (
	"fmt"
	"strings"
)

// func 함수명(매개변수 타입) 반환타입 { }
func multiply(a uint, b uint) uint {
	return a * b
}

func lenAndUpper(name string) (length int, uppercase string) {
	// 함수가 끝나기 직전에 실행되는 코드
	// 1. 리소스를 릴리즈하는 형태의 코드를 작성할 수 있다
	// 2. 가장 마지막에 지정된 defer부터 실행
	// 3. defer에 사용된 변수는 선언될 때 평가되므로 이후에 변수값이 바뀌더라도 defer에 영향을 주지 않음
	defer fmt.Println("I'm done")

	// 반환값에 이름을 붙이면 변수처럼 사용 가능하고 리턴만 처리하면 됨
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
	// return len(name), strings.ToUpper(name)
}

func repeatMe(words ...string) {
	fmt.Println((words))
}

func superAdd(numbers ...int) int {
	total := 0
	// // 일반적인 반복문
	// for i := 0; i < len(numbers); i++ {
	// 	fmt.Println(numbers[i])
	// }

	// // numbers 배열 크기만큼 반복
	// for index := range numbers {
	// 	fmt.Println(numbers[index])
	// }

	// numbers 배열 크기만큼 반복
	for _, number := range numbers {
		total += number
	}

	return total
}

// func canIDrink(age int) bool {
// 	// if 변수값 생성; 조건식
// 	if koreanAge := age + 2; koreanAge < 20 {
// 		return false
// 	}
// 	return true
// }

func canIDrink(age int) bool {
	switch age {
	case 10:
		return false
	case 18:
		return true
	}
	return false
}

type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	// fmt.Println("Hello world!")

	// 상수(const)
	// const name string = "nico"
	// const name = "nico"
	// name = "Lynn"

	// 변수(var)
	// var name string = "nico"
	// var name = "nico"
	// name = "lynn"

	// var name = "nico" 단축 표기
	// local variable(함수내부)일 때만 가능
	name := "nico"
	name = "lynn"

	fmt.Println(name)

	// 함수 호출
	fmt.Println(multiply(2, 2))

	// length, upperedName := lenAndUpper("nico")
	// fmt.Println(length, upperedName)

	// _: 무시된 값
	length, _ := lenAndUpper("nico")
	fmt.Println(length)

	repeatMe("nico", "lynn", "dal", "marl", "flynn")

	total := superAdd(1, 2, 3, 4, 5, 6)
	fmt.Println(total)

	fmt.Println(canIDrink(10))
	fmt.Println(canIDrink(18))

	// 값의 복사
	// a := 2
	// b := a
	// a = 10
	// 참조 복사
	// &: 변수의 주소, *: 포인터 변수가 참조하고 있는 주소의 값
	a := 2
	b := &a // 포인터변수 생성
	a = 10
	*b = 20 // 포인터변수로 값 변경
	fmt.Println(a, *b)

	// array: 길이 선언이 있음
	// names := [5]string{"nico", "lynn", "dal"}
	// array: 인덱스로 접하여 추가
	// names[3] = "alala"
	// names[4] = "blabla"

	// slice: 길이 선언이 없음
	names := []string{"nico", "lynn", "dal"}
	// slice: append 메서드로 접근하여 추가
	// 새로운 slice를 반환
	names = append(names, "alala")

	fmt.Println(names[0])
	fmt.Println(names)

	// // map
	// // map[key타입]value타입
	// nico := map[string]string{"name": "nico", "age": "12"}
	// fmt.Println(nico)
	// // map 반복문
	// for key, value := range nico {
	// 	fmt.Println(key, value)
	// }

	// struct
	favFood := []string{"kimchi", "ramen"}
	nico := person{name: "nico", age: 18, favFood: favFood}
	fmt.Println(nico.name)
}
