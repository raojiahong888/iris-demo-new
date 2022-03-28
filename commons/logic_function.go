package commons

import (
	"sort"
)

// replace the fixed element from a simple array into a new array
func ReplaceFixedElementFromSimpleArray(arr []int, element int) (newArr []int) {
	newArr = make([]int, 0, len(arr))
	for _,v := range arr {
		if element != v {
			newArr = append(newArr, v)
		}
	}
	return
}

func Intersection(arrOne, arrTwo []int ) (newArr []int) {
	arrOneTemp := make(map[int]bool, len(arrOne))
	for _,oneValue := range arrOne {
		arrOneTemp[oneValue] = true
	}

	arrTwoTemp := make(map[int]bool, len(arrTwo))
	for _,twoValue := range arrTwo {
		arrTwoTemp[twoValue] = true
	}

	if len(arrOneTemp) > len(arrTwoTemp) {
		arrOneTemp, arrTwoTemp = arrTwoTemp, arrOneTemp
	}

	newArr = make([]int, 0, len(arrTwoTemp))
	for k, _ := range arrOneTemp {
		if _,ok := arrTwoTemp[k]; ok {
			newArr = append(newArr, k)
		}
	}

	return
}

func IntersectionV2(nums1 []int, nums2 []int) (res []int) {
	sort.Ints(nums1)
	sort.Ints(nums2)
	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		x, y := nums1[i], nums2[j]
		if x == y {
			if res == nil || x > res[len(res)-1] {
				res = append(res, x)
			}
			i++
			j++
		} else if x < y {
			i++
		} else {
			j++
		}
	}
	return
}

// a simple method used to search the key of the fixed element by using mid search
func SimpleMidSearch(arr []int, element int) (mid int) {
	sort.Ints(arr)
	low := 0
	high := len(arr) - 1
	for low <= high {
		mid = low + (high - low)/2
		if arr[mid] == element {
			return
		}
		if arr[mid] > element {
			high = mid - 1
			continue
		}
		if arr[mid] < element {
			low = mid + 1
			continue
		}
	}
	return
}

func TheFirstEqualSearch(arr []int, element int) (mid int) {
	sort.Ints(arr)
	low := 0
	high := len(arr) - 1
	for low <= high {
		mid = low + (high - low)/2
		if arr[mid] > element {
			high = mid - 1
			continue
		}
		if arr[mid] < element {
			low = mid + 1
			continue
		}
		if mid==0 || arr[mid-1] != element {
			return
		} else {
			high = mid - 1
		}
	}
	return
}


func TheEndEqualSearch(arr []int, element int) (mid int) {
	sort.Ints(arr)
	low := 0
	high := len(arr) - 1
	for low <= high {
		mid = low + (high - low)/2
		if arr[mid] > element {
			high = mid - 1
			continue
		}
		if arr[mid] < element {
			low = mid + 1
			continue
		}
		if mid==0 || arr[mid+1] != element {
			return
		} else {
			low = mid + 1
		}
	}
	return
}