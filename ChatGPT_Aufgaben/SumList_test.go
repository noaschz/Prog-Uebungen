package chatgptaufgaben

import "fmt"

func ExampleSumList() {
	fmt.Println(SumList(nil))
	fmt.Println(SumList([]int{}))
	fmt.Println(SumList([]int{5}))
	fmt.Println(SumList([]int{-3, 7, -1}))
	fmt.Println(SumList([]int{2, 9, 4, 9, 1}))
	// Output:
	// 0
	// 0
	// 5
	// 3
	// 25
}
