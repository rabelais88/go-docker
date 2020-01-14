package sum

func mySum(arg ...int) int {
	count := 0
	for _, v := range arg {
		count += v
	}
	return count
}
