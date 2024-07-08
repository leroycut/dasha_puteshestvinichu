package main

import (
	"fmt"
	"strings"
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

func task_21(arr []int) []int {
	res := []int{}
	for i := 0; i < len(arr); i++ {
		var tf = true
		for j := 0; j < i; j++ {
			if arr[i] == arr[j] {
				tf = false
			}
		}
		if tf == true {
			res = append(res, arr[i])
		}
	}
	return res
}

func task_22(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				var sup = arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = sup
			}
		}
	}
	return arr
}

func task_23(n int) {
	var f_last = 0
	var f_new = 1
	fmt.Println(f_new)
	for i := 1; i < n; i++ {
		var sup = f_new
		f_new = sup + f_last
		f_last = sup
		fmt.Println(f_new)
	}
}

func task_24(arr []int, el int) int {
	var cnt = 0
	for i := range arr {
		if el == arr[i] {
			cnt++
		}
	}
	return cnt
}

func task_25(arr1 []int, arr2 []int) []int {
	res := []int{}
	for i := 0; i < len(arr1); i++ {
		var tf = false
		for j := 0; j < len(arr2); j++ {
			if arr1[i] == arr2[j] {
				tf = true
			}
		}
		for j := 0; j < i; j++ {
			if arr1[i] == arr1[j] {
				tf = false
			}
		}
		if tf == true {
			res = append(res, arr1[i])
		}
	}
	return res
}

func task_26(s1, s2 string) {
	s1 = strings.ToLower(s1)
	s2 = strings.ToLower(s2)

	arr1 := make([]int, 26)
	arr2 := make([]int, 26)

	for i := range s1 {
		arr1[int(s1[i])-int('a')]++
	}

	for i := range s2 {
		arr2[int(s2[i])-int('a')]++
	}

	var tf = true
	for i := range arr1 {
		if arr1[i] != arr2[i] {
			tf = false
		}
	}

	if tf == true {
		fmt.Println("It's anagram")
	} else {
		fmt.Println("It's not anagram")
	}
}

func task_27(arr1, arr2 []int) []int {
	res := make([]int, 0)
	var point1 = 0
	var point2 = 0
	for point1 < len(arr1) {
		if point2 < len(arr2) && arr1[point1] > arr2[point2] {
			res = append(res, arr2[point2])
			point2++
		} else {
			res = append(res, arr1[point1])
			point1++
		}
	}
	for point2 < len(arr2) {
		res = append(res, arr2[point2])
		point2++
	}
	return res
}

// task_28

type vertex struct {
	key         string
	val         int
	next_vertex *vertex
}

type Hash_Table struct {
	sz     int
	matrix []*vertex
}

func get_hash(str string, sz int) int {
	hash := 0
	for i := range str {
		hash += int(str[i])
	}
	return hash % sz
}

func (ht *Hash_Table) add_el(str string, cur_val int) {
	hash := get_hash(str, ht.sz)
	if ht.matrix[hash] == nil {
		ht.matrix[hash] = &vertex{key: str, val: cur_val}
	} else {
		cur_ver := ht.matrix[hash]
		for cur_ver.next_vertex != nil {
			if cur_ver.key == str {
				cur_ver.val = cur_val
				return
			}
			cur_ver = cur_ver.next_vertex
		}
		cur_ver = &vertex{key: str, val: cur_val}
	}
}

func task_29(arr []int, el int) bool {
	var l = 0
	var r = len(arr)
	for r-l > 1 {
		var mid = (l + r) / 2
		if arr[mid] > el {
			r = mid
		} else {
			l = mid
		}
	}
	return arr[l] == el
}

// task 30

type queue struct {
	stack_1 []int
	stack_2 []int
}

func (q *queue) add_el_in_queue(val int) {
	q.stack_1 = append(q.stack_1, val)
}

func (q *queue) get_el_from_queue() int {
	if len(q.stack_2) == 0 {
		if len(q.stack_1) == 0 {
			fmt.Errorf("empty")
			return 0
		}
		for len(q.stack_1) > 0 {
			q.stack_2 = append(q.stack_2, q.stack_1[0])
			q.stack_1 = q.stack_1[1:]
		}
	}

	res := q.stack_2[0]
	q.stack_2 = q.stack_2[1:]
	return res
}

func main() {

	// arr1 := []int{1, 2, 2, 3, 5, 9, 14}
	// arr2 := []int{1, 3, 4, 7, 13, 26}
	// res := task_27(arr1, arr2)
	// for i := range res {
	// 	fmt.Println(res[i])
	// }

	// fmt.Println(task_29(arr1, 6))
}
