package removeelement

func RemoveElement(nums []int, val int) int {
	if len(nums) < 1 {
		return 0
	}
	j := 0
	for i := 0; i < len(nums); i++ {
		if val != nums[i] {
			nums[j] = nums[i]
			j++
		}
	}
	return j
}
