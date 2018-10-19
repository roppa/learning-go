package arithmetic

func Add(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}
