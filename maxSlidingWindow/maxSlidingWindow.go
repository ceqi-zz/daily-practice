package main

import (
	"container/list"
	"fmt"
)

func maxValues(nums []int, k int) []int {
	q := list.New()
	maxValues := []int{}
	// as always, be aware of boundary when playing with arrays
	for i := 0; i < len(nums); i++ {
		// maintaining a monoque, a descrsing / increasing queue
		if q.Len() != 0 && i-q.Front().Value.(int) >= k {
			q.Remove(q.Front())
		}

		for q.Len() != 0 && nums[q.Back().Value.(int)] < nums[i] {
			q.Remove(q.Back())
		}

		q.PushBack(i)

		if q.Len() != 0 && i >= k-1 {
			maxValues = append(maxValues, nums[q.Front().Value.(int)])
		}

	}

	return maxValues

}

func main() {
	nums := []int{3, 5, 8, 7, 9, 4, 1}
	k := 3
	fmt.Println(maxValues(nums, k))
}
