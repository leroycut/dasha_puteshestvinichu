package main

import (
	"fmt"
)

func task_1() {
	fmt.Println("Hello, World!")
}

func task_2(a, b int) int {
	return a + b
}

func task_3(a int) {
	if (a & 1) == 1 {
		fmt.Println("This is odd number")
	} else {
		fmt.Println("This is even number")
	}
}

func task_4(a, b, c int) int {
	return max(a, b, c)
}

func task_5(a int) int {
	var res = 1
	for i := 1; i <= a; i++ {
		res *= i
	}
	return res
}

func task_6(sym byte) {
	if sym == 'a' || sym == 'e' || sym == 'i' || sym == 'o' || sym == 'u' || sym == 'y' {
		fmt.Println("This is vowel letter")
	} else {
		fmt.Println("This isn't vowel letter")
	}
}

func task_7(n int) {

	prime := make([]int, n+1)

	for i := range prime {
		prime[i] = 1
	}

	prime[0] = 0
	prime[1] = 0

	for i := 2; i <= n; i++ {
		if prime[i] == 1 {
			fmt.Println(i)
			for j := i * i; j <= n; j += i {
				prime[j] = 0
			}
		}
	}
}

func task_8(str string) string {
	var sz = len(str)
	res := ""
	for i := 0; i < sz; i++ {
		res += string(str[sz-i-1])
	}
	return res
}

func task_9(arr []int) int {
	var sum = 0
	for i := range arr {
		sum += arr[i]
	}
	return sum
}

// 10 таска

type Rectangle struct {
	Hight int
	Width int
}

func get_area(r Rectangle) int {
	return r.Hight * r.Width
}

func task_11(a float32) float32 {
	return (a * 1.8) + 32
}

func task_12(n int) {
	for i := n; i > 0; i-- {
		fmt.Println(i)
	}
}

func task_13(str string) int {
	var size = 0
	for i := range str {
		size++
		i = i
	}
	return size
}

func task_14(arr []int, el int) {
	for i := range arr {
		if el == arr[i] {
			fmt.Println("In array")
			return
		}
	}
	fmt.Println("Out array")
}

func task_15(arr []int) float32 {
	sum := 0
	for i := range arr {
		sum += arr[i]
	}
	res := float32(sum)
	return res / float32(len(arr))
}

func task_16(a int) {
	for i := 1; i < 10; i++ {
		fmt.Println(a * i)
	}
}

func task_17(str string) {
	corect := true
	sz := len(str)
	for i := 0; i < sz/2; i++ {
		if str[i] != str[sz-1-i] {
			corect = false
		}
	}
	if corect == true {
		fmt.Println("This is palindrome")
	} else {
		fmt.Println("This isn't palindrome")
	}
}

func task_18(arr []int) []int {
	Max := arr[0]
	Min := arr[0]
	for i := range arr {
		Max = max(Max, arr[i])
		Min = min(Min, arr[i])
	}
	res := make([]int, 2)
	res[0] = Max
	res[1] = Min
	return res
}

func task_19(arr []int, ind int) []int {
	return append(arr[:ind], arr[ind+1:]...)
}

func task_20(arr []int, val int) int {
	res := -1
	for i := range arr {
		if arr[i] == val {
			res = i
		}
	}
	return res
}

func main() {

	fmt.Println(task_13("arr"))

}
