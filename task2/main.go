package main

import "fmt"

func isThereSumOfTwoInArray(nums []int, sum int) bool {
	nums = quickSort(nums)
	fmt.Println(nums)
	for _, num := range nums {
		if num > (sum/2){
			return false
		}
		if _, ok := binarySearch(nums, sum - num); ok{
			return true
		}
	}
	return false
}

func partition(lo int, piv int, arr []int) int {
	is := lo

	for i := lo; i < piv; i++ {
		if arr[i] < arr[piv] {
			if i != is {
				arr[i], arr[is] = arr[is], arr[i]
			}

			is++
		}
	}

	arr[is], arr[piv] = arr[piv], arr[is]

	if is - 1 > lo {
		partition(lo, is - 1, arr)
	}
	if is + 1 < piv {
		partition(is + 1, piv, arr)
	}

	return is
}

func quickSort(arr []int) []int {
	l := len(arr)
	piv := l - 1

	partition(0, piv, arr)

	return arr
}

func binarySearch(haystack []int, needle int) (int, bool) {
	start := 0
	end := len(haystack) - 1

	for start <= end {

		median := (start + end) / 2

		if haystack[median] < needle {
			start = median + 1
		} else {
			end = median - 1
		}

	}

	if start == len(haystack) || haystack[start] != needle {
		return -1, false
	} else {
		return start, true
	}
}

func main() {
	test := []int{3,8,3,6,7,2,3,9,0,1,4}
	fmt.Println(isThereSumOfTwoInArray(test,12))//true
	fmt.Println(isThereSumOfTwoInArray(test,120))//false
}