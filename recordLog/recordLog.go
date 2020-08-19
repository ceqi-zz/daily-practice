package main

import (
	"container/list"
	"fmt"
)

var l = list.New()

func record(orderID int) {
	l.PushBack(orderID)
}

func getLast(i int, e *list.Element) *list.Element {

	if l.Len() == 0 {
		return nil
	}

	if i == 1 {
		return e
	}

	return getLast(i-1, e.Prev())

	// last = l.Back()

	// if i == 1 {
	// 	return last
	// }

	// for j := 1; j < i; j++ {
	// 	last = last.Prev()
	// }

}

func main() {
	record(1)
	record(2)
	record(3)
	record(300)

	last := getLast(3, l.Back())

	fmt.Println(last.Value)

}
