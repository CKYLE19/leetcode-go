package leetcode

import "fmt"

func lemonadeChange(bills []int) bool {
	cost := 5
	changeToGive := 0
	for _, customer := range bills {
		diff := customer - cost
		fmt.Printf("Cost: %d\nHas: %d\nChange To Give: %d\n\n", cost, customer, changeToGive)
		if diff == 0 {
			changeToGive += customer
		} else {
			if diff > changeToGive {
				return false
			}
			changeToGive -= cost
			changeToGive += diff
		}
	}
	return true
}
