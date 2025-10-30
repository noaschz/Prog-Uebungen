package chatgptaufgaben

import "fmt"

func ExampleIndexOf() {
	fmt.Println(IndexOf(nil, 7))
	fmt.Println(IndexOf([]int{}, 7))
	fmt.Println(IndexOf([]int{1, 7, 7, 2}, 7))
	fmt.Println(IndexOf([]int{1, 2, 3}, 4))
	// Output:
	// -1
	// -1
	// 1
	// -1
}
