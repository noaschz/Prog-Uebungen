package chatgptaufgaben

// IndexOf liefert den ersten Index von x in nums oder -1, falls x nicht enthalten ist.
// Randfälle:
//   - nil oder leere Liste: -1
func IndexOf(nums []int, x int) int {
	for i, v := range nums {
		if v == x {
			return i
		}
	}
	return -1
}

