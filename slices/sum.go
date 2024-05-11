package slices

// Sum 计算一个整数切片（slice）中所有元素的和。
// 它接受一个整数切片 `numbers` 并返回它们的总和。
func Sum(numbers []int) int {
	sum := 0                         // 初始化 sum 变量来存储和的结果。
	for _, number := range numbers { // 遍历切片中的每个元素。
		sum += number // 将当前元素加到 sum 上。
	}
	return sum // 返回计算得到的总和。
}

// SumAll 接受多个整数切片，计算每个切片的和，然后返回一个包含这些和的切片。
// `numbersToSum` 是一个可变参数，允许传入任意数量的整数切片。
func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}
