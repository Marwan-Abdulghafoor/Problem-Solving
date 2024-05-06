package main

func main() {
	nums := []int{3, 2, 2, 3}
	ans := removeElements(nums, 3)
	println(ans)
}

func removeElements(nums []int, val int) int {
	temp := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] != val {
			nums[temp] = nums[i]
			temp++
		}
	}
	return temp
}
