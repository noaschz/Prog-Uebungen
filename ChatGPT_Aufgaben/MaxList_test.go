package chatgptaufgaben

import "fmt"

func ExampleMaxList() {
	fmt.Println(MaxList(nil))
	fmt.Println(MaxList([]int{}))
	fmt.Println(MaxList([]int{5}))
	fmt.Println(MaxList([]int{3, 7, 1}))
	fmt.Println(MaxList([]int{2, 9, 4, 9, 1}))
	// Output:
	// 0
	// 0
	// 5
	// 7
	// 9
}
