package main

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
func SumAll(numbersToSum ...[]int) (sums []int) {
	lengthOfNumbers := len(numbersToSum) // 获取传入切片的数量。
	sums = make([]int, lengthOfNumbers)  // 根据切片数量创建一个新的整数切片，用于存储每个切片的和。

	for i, numbers := range numbersToSum { // 遍历每个切片。
		sums[i] = Sum(numbers) // 调用 Sum 函数计算当前切片的和，并将结果存储到 sums 切片中。
	}
	return sums // 返回包含所有和的切片。
}
