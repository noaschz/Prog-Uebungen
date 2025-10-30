package chatgptaufgaben

import "fmt"

func ExampleCountEqual() {
	fmt.Println(CountEqual(nil, 1))
	fmt.Println(CountEqual([]int{}, 1))
	fmt.Println(CountEqual([]int{1, 1, 2, 1}, 1))
	fmt.Println(CountEqual([]int{2, 3, 4}, 1))
	// Output:
	// 0
	// 0
	// 3
	// 0
}
