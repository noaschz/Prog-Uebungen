package chatgptaufgaben

// SumList gibt die Summe aller Elemente von nums zurück.
// Randfälle:
//   - nil oder leere Liste: 0
//   - int-Überläufe werden nicht speziell behandelt
func SumList(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}
