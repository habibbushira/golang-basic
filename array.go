package main

import "fmt"

func main() {
	//Array variable declaration
	var ages [3]int = [3]int{2, 3, 2}
	fmt.Println(ages, len(ages))

	//short hand definition
	var shortAges = [3]int{20, 23, 23}
	fmt.Println(shortAges, len(shortAges))

	// short hand assignment
	names := [3]string{"a", "b", "c"}
	names[1] = "Lewije"
	fmt.Println("Names: ", names)

	// slices (use array under the hood)
	var scores = []int{2, 3, 5}
	scores[2] = 3
	// add a value to a slice
	scores = append(scores, 83)
	fmt.Println(scores, len(scores))

	// slices ranges
	// inclusive of index one but not index three
	rangeOne := names[1:3]
	fmt.Println(rangeOne)

	//from position two (including) upto the last position (including)
	rangeTwo := names[2:]
	fmt.Println(rangeTwo)

	//from the first element up to position two excluding the last element
	rangeThree := names[:3]
	fmt.Println(rangeThree)

	rangeOne = append(rangeOne, "kuper")
	fmt.Println(rangeOne)
}
