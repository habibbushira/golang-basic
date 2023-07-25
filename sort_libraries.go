package main

import (
	"fmt"
	"sort"
)

func main() {
	ages := []int{32, 12, 34, 33, 11, 9, 89, 23, 2}

	//This ulter the orginal variable
	sort.Ints(ages)
	fmt.Println(ages)

	// If value not found, index will be beyond the length of the slice
	index := sort.SearchInts(ages, 33)
	fmt.Println("Index of 33", index)

	names := []string{"org", "value", "value", "test", "ohio", "text"}
	sort.Strings(names)
	fmt.Println(names)

	fmt.Println(sort.SearchStrings(names, "value"))
}
