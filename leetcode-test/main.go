package main

import "fmt"

func main() {
	// result := minSubArrayLen(11, []int{1, 2, 3, 4, 5})
	var arg = [][]int{{2, 3}, {4, 5}, {6, 7}, {8, 9}, {1, 10}}
	// result := arg[1:3]
	result := merge(arg)

	fmt.Printf("result %v %d %d %v\n", result, len(result), cap(result), result[0])
}

func minSubArrayLen(target int, nums []int) int {
	sum := 0
	minLen := -1
	for j := 0; j < len(nums); j++ {
		if nums[j] >= target {
			return 1
		} else {
			sum = nums[j]
		}
		for k := j + 1; k < len(nums); k++ {
			sum += nums[k]
			if sum >= target {
				if k-j < minLen || minLen == -1 {
					minLen = k - j
				}
			}
		}
	}
	return minLen + 1
}

// 杨辉三角
func generate(numRows int) [][]int {
	if numRows == 1 {
		row := []int{1}
		row2 := []int{1, 2}
		array := [][]int{}
		array = append(array, row)
		array = append(array, row2)
		return array
	}
	return nil
}

func removeDuplicates(nums []int) int {
	dupCount := 0
	if len(nums) <= 1 {
		return len(nums)
	}

	slow := len(nums) - 1
	fast := len(nums) - 2
	for fast >= 0 {
		if nums[fast] == nums[slow] {
			fmt.Println("remove")

			if slow+1 == len(nums) {
				nums = nums[0 : fast+1]
			} else {
				nums = append(nums[0:fast+1], nums[slow+1:]...)
			}
			dupCount++
		}
		slow = fast
		fast--
	}
	return len(nums)
}

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}
	sortedM := fastSortMatrix(intervals)
	merged := [][]int{}
	slow := -1
	fast := 0
	var maxValue int = -1
	for fast < len(intervals) {
		fmt.Printf("fast %d maxValue %d sortedM[fast][0]%d\n", fast, maxValue, sortedM[fast][0])
		if maxValue < sortedM[fast][0] {
			merged = append(merged, sortedM[fast])
			slow++
			maxValue = sortedM[fast][1]
		} else {
			if sortedM[fast][1] > maxValue {
				maxValue = sortedM[fast][1]
			}
			merged[slow][1] = maxValue
		}
		fast++
	}

	return merged

}
func fastSortMatrix(m [][]int) [][]int {
	if len(m) <= 1 {
		return m
	}
	base := m[0][0]
	left := 0
	right := len(m) - 1
	for left < right {
		for left < right && m[right][0] >= base {
			right--
		}
		if right == left {
			break
		}
		m[left], m[right] = m[right], m[left]
		for left < right && m[left][0] <= base {
			left++
		}
		if left == right {
			break
		}
		m[left], m[right] = m[right], m[left]
	}
	fastSortMatrix(m[0:left])
	fastSortMatrix(m[right+1:])
	return m
}
