package chatgptaufgaben

// Erwartet eine Liste von Zahlen und gibt das Maximum zurück.
// Wenn die Liste leer ist, wird 0 zurückgegeben.
func MaxList(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}
