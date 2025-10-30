package chatgptaufgaben

import "fmt"

func ExampleReverseInPlace() {
	a := []int{1, 2, 3, 4}
	ReverseInPlace(a)
	fmt.Println(a)
	b := []int{}
	ReverseInPlace(b)
	fmt.Println(b)
	// Output:
	// [4 3 2 1]
	// []
}
