package main

import "fmt"

func sortColors(nums []int) {
	length := len(nums)
	if length == 0 {
		return
	}
	zero := 0
	two := length - 1
	// for zero < length && nums[zero] == 0 {
	// 	zero++
	// }
	// if zero == length {
	// 	return
	// }
	// for two >= 0 && nums[two] == 2 {
	// 	two--
	// }
	// if two < 0 {
	// 	return
	// }

	i := zero

	for i <= two {
		if nums[i] == 0 {
			nums[i], nums[zero] = nums[zero], nums[i]
			if i == zero {
				i++
			}
			zero++
		} else if nums[i] == 2 {
			nums[i], nums[two] = nums[two], nums[i]
			two--
		} else {
			i++
		}
	}
}

func main() {
	ans := []int{0, 2, 1}
	sortColors(ans)
	fmt.Println(ans)
}
