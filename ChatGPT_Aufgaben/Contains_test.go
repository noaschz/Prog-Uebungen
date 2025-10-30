package chatgptaufgaben

import "fmt"

func ExampleContains() {
	fmt.Println(Contains(nil, 5))
	fmt.Println(Contains([]int{}, 5))
	fmt.Println(Contains([]int{1, 2, 3}, 2))
	fmt.Println(Contains([]int{1, 2, 3}, 5))
	// Output:
	// false
	// false
	// true
	// false
}
