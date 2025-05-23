package problems

func extraChar(str1 string, str2 string) int {
	sum1 := 0
	sum2 := 0

	i := 0
	for i = range len(str1) {
		sum1 += int(str1[i])
	}

	for i = range len(str2) {
		sum2 += int(str2[i])
	}

	return sum1 - sum2
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	i := 0
	for i = range nums {
		difference := target - nums[i]
		j, ok := m[difference]
		if ok {
			return []int{j, i}
		} else {
			m[nums[i]] = i
		}
	}
	return nil
}

func MaxArea(height []int) int {
	n := len(height)
	i := 0
	j := 0
	area := 0

	for i != n-1-j {
		newArea := min(height[i], height[n-1-j]) * (n - 1 - i - j)
		if newArea > area {
			area = newArea
		}

		if height[i] <= height[n-1-j] {
			i++
			continue
		} else {
			j++
			continue
		}
	}
	return area
}

func MaxSubArray(nums []int) int {
	maxSum := nums[0]
	sum := 0

	for i := range nums {
		if sum < 0 && nums[i] > sum {
			sum = nums[i]
			if sum > maxSum {
				maxSum = sum
			}
			continue
		}
		if sum > 0 && sum+nums[i] < 0 || sum < 0 && sum+nums[i] < sum {
			sum = 0
			i++
		} else {
			sum += nums[i]
			if sum > maxSum {
				maxSum = sum
			}
		}
	}

	return maxSum
}
