package util

// BinarySearchFirstGreater 二分查找，返回第一个大于给定值的元素索引
func BinarySearchFirstGreater(arr []int, v int) int {
	size := len(arr)
	if size == 0 {
		return -1
	}

	return bsFirstGreaterRecursive(arr, v, 0, size-1)
}

// bsFirstGreaterRecursive 二分查找，递归实现
func bsFirstGreaterRecursive(arr []int, v, low, high int) int {
	maxIndex := len(arr) - 1

	if low == high && low != 0 {
		if low >= maxIndex {
			return -1
		}
		return low + 1
	} else if low > high {
		return -1
	}

	mid := (low + high) / 2
	if arr[mid] <= v && arr[mid+1] >= v {
		return mid + 1
	} else if arr[mid] < v {
		return bsFirstGreaterRecursive(arr, v, mid+1, high)
	} else {
		return bsFirstGreaterRecursive(arr, v, low, mid-1)
	}
}
