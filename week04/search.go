func search(nums []int, target int) int {
	var low, mid, high int
	low = 0
	high = len(nums) - 1
	for low <= high {
		mid = low + (high - low) / 2
		if nums[mid] == target {
			return mid
		}
		// 左边有序，如果有序满足情况，就在有序中找，否则在无序中找
		if nums[low] <= nums[mid] {
			if nums[low] <= target && target < nums[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else {
			// 右边有序
			if nums[mid] < target && target <= nums[high] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}
	return -1
}
