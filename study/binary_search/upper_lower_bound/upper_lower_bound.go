package main

import "fmt"

func main() {
	arr := []int{
		1, 3, 3, 3, 4, 5, 6,
	}
	/*
		lowerBound: x 이상의 값이 처음 발견되는 index를 return
		upperBound: x 초과하는 값이 처음 발견되는 index
	*/
	fmt.Println(lowerBound(arr, 2)) // return 1
	fmt.Println(upperBound(arr, 2)) // return 1
	fmt.Println(lowerBound(arr, 3)) // return 1
	fmt.Println(upperBound(arr, 3)) // return 4

}
// lower, upperbound는 많이 헷갈리니 양식을 외우고, 그냥 가져다 쓰자 
/*
 1. start < end while loop를 돈다.
 2. 원하는 값 <= arr[mid] 이면 lower, 원하는 값 >= arr[mid]이면 upper
 3. start = mid + 1, end = mid는 변하지 않고, 2번만 바꿔주면 된다.
*/ 
func lowerBound(arr []int, x int) int {
	start := 0
	end := len(arr)
	for start < end {
		mid := (start + end) >> 1
		if x > arr[mid] {
			start = mid + 1
		} else {
			end = mid
		}
	}
	return start
}

func upperBound(arr []int, x int) int {
	start := 0
	end := len(arr)
	for start < end {
		mid := (start + end) >> 1
		if x >= arr[mid] {
			start = mid + 1

		} else {
			end = mid
		}
	}
	return start
}
