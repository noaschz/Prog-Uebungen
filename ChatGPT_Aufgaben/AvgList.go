package chatgptaufgaben

// AvgList gibt den ganzzahligen Durchschnitt (Abrundung) der Elemente zurück.
// Randfälle:
//   - nil oder leere Liste: 0
//   - ganzzahliges Abrunden durch Integer-Division
func AvgList(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum / len(nums)
}
