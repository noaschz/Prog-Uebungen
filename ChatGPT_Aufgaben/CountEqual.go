package chatgptaufgaben

// CountEqual zählt, wie oft x in nums vorkommt.
// Randfälle:
//   - nil oder leere Liste: 0
func CountEqual(nums []int, x int) int {
	count := 0
	for _, v := range nums {
		if v == x {
			count++
		}
	}
	return count
}
