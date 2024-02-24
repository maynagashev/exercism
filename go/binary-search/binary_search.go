package binarysearch

func SearchInts(list []int, key int) int {
	low, high := 0, len(list)-1
	for low <= high {
		mid := (low + high) / 2
		if list[mid] == key {
			return mid
		}
		if list[mid] < key {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
