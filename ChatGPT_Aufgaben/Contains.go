package chatgptaufgaben

// Contains liefert true, wenn x mindestens einmal in nums vorkommt.
// Randfälle:
//   - nil oder leere Liste: false
func Contains(nums []int, x int) bool {
	for _, v := range nums {
		if v == x {
			return true
		}
	}
	return false
}
