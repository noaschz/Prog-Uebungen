package chatgptaufgaben

import "fmt"

func ExampleAvgList() {
	fmt.Println(AvgList(nil))
	fmt.Println(AvgList([]int{}))
	fmt.Println(AvgList([]int{5}))
	fmt.Println(AvgList([]int{2, 3}))       // 2
	fmt.Println(AvgList([]int{-3, 7, -1}))  // 1
	fmt.Println(AvgList([]int{2, 9, 4, 9})) // 6
	// Output:
	// 0
	// 0
	// 5
	// 2
	// 1
	// 6
}
