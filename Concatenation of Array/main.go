package main

func main() {

}

func getConcatenation(nums []int) []int {
	ans := make([]int, 2*len(nums))

	for i := 0; i < len(nums); i++ {
		ans[i], ans[i+len(nums)] = nums[i], nums[i]
	}
	return nums
}
