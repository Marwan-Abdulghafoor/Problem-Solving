package main

func main() {

	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	answer := removeDuplicates(nums)
	println(answer)
}

func removeDuplicates(nums []int) int {
	temp := 1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] != nums[i+1] {
			nums[temp] = nums[i+1]
			temp++
		}
	}

	return temp

}
