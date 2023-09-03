package leetcode

import "fmt"

//TrappingRainWaterIterate https://leetcode.cn/problems/trapping-rain-water/
//height = [0,1,0,2,1,0,1,3,2,1,2,1] ans 6
func TrappingRainWaterIterate(nums []int) int {
	sum := 0
	max := getMax(nums)
	for j := 1; j <= max; j++ {
		startFlag := false
		tmp := 0
		for i := 0; i < len(nums); i++ {
			fmt.Println(i, nums[i], startFlag, tmp)
			if startFlag && nums[i] < j {
				tmp++
			}
			if nums[i] >= j {
				sum += tmp
				tmp = 0
				startFlag = true
			}
		}
		fmt.Println("level", j, sum)
	}

	return sum
}

func getMax(nums []int) int {
	tmp := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > tmp {
			tmp = nums[i]
		}
	}
	return tmp
}
