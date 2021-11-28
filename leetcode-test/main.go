package main

import (
	"fmt"
	"log"
)

func main() {
	// result := minSubArrayLen(11, []int{1, 2, 3, 4, 5})
	var arg = []int{0, 3, -3, 4, -1}
	// result := arg[1:3]

	fmt.Printf("result %v %d %d\n", twoSum(arg, -1), len(twoSum(arg, -1)), cap(twoSum(arg, -1)))
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
func twoSum(nums []int, target int) []int {
	quickSort(nums, 0, len(nums)-1)
	for j := 0; j < len(nums); j++ {
		res := binarySearch(nums, target-nums[j])
		if res != -1 && res != j {
			return []int{j, res}
		}
	}
	return nil
}

func quickSort(nums []int, left int, right int) {
	if left >= right {
		return
	}
	mid := partition(nums, left, right)
	quickSort(nums, left, mid-1)
	quickSort(nums, mid+1, right)
}

func partition(nums []int, left int, right int) int {

	i, j := left, right
	base := nums[left]
	for i < j {
		for i < j && nums[j] >= base {
			j--
		}
		if i == j {
			break
		}
		for i < j && nums[i] <= base {
			i++
		}
		if i == j {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	// i==j
	if nums[j] > base {
		j--
	}
	nums[j], nums[left] = nums[left], nums[j]
	log.Printf("partition %v %d\n", nums, j)

	return j
}

func binarySearch(nums []int, target int) int {
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}
	if len(nums) == 0 {
		return -1
	}
	left, right := 0, len(nums)-1
	for left < right {
		mid := (right-left)/2 + left
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	if left == right && nums[left] == target {
		return left
	} else {
		return -1
	}
}
