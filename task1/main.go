package main

import (
	"errors"
	"fmt"
	"log"
)

func pairsInGap(nums []int, gap int) (map[int]int, error) {
	if gap < 1{
		return nil, errors.New("Invalid gap")
	}

	result := map[int]int{}

	for i := 0; i < (len(nums) - gap); i++ {
		if nums[i] == nums[i+gap]{
			result[nums[i]]++
		}
	}
	return result, nil
}

func main() {
	test := []int{1, 3, 1, 3, 1, 3, 8, 8, 3}
	pairs, err := pairsInGap(test,2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(pairs)
}