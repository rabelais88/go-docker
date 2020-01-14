// 테스트용 모듈
package sum

// 모든 인자를 더하여 정수 합을 반환한다.
func mySum(arg ...int) int {
	count := 0
	for _, v := range arg {
		count += v
	}
	return count
}
