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

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1 := len(nums1)
	l2 := len(nums2)

	mid := (l1 + l2 - 1) / 2

	if l1 == 0 {
		if (l1+l2)%2 == 0 {
			return (float64(nums2[mid]) + float64(nums2[mid+1])) / 2
		} else {
			return float64(nums2[mid])
		}
	}

	if l2 == 0 {
		if (l1+l2)%2 == 0 {
			return (float64(nums1[mid]) + float64(nums1[mid+1])) / 2
		} else {
			return float64(nums1[mid])
		}
	}

	nums3 := []int{}

	i := 0
	j := 0
	for true {
		if j >= l2 {
			nums3 = append(nums3, nums1[i:]...)
			break
		} else if i >= l1 {
			nums3 = append(nums3, nums2[j:]...)
			break
		} else {
			if nums1[i] <= nums2[j] {
				nums3 = append(nums3, nums1[i])
				i++
			} else {
				nums3 = append(nums3, nums2[j])
				j++
			}
		}

	}

	if (l1+l2)%2 == 0 {
		return (float64(nums3[mid]) + float64(nums3[mid+1])) / 2
	} else {
		return float64(nums3[mid])
	}

}

func LengthOfLongestSubstring(s string) int {
	byteCount := map[byte]int{}

	max := 0
	i := 0
	j := 0

	for i = 0; i < len(s); i++ {
		if _, ok := byteCount[s[i]]; !ok {
			byteCount[s[i]] = 1
		} else {
			clear(byteCount)
			j++
			i = j
			byteCount[s[i]] = 1
		}
		if len(byteCount) > max {
			max = len(byteCount)
		}
	}

	return max
}
