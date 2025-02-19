package Assignment1

/*
Even sum. [WEIGHT = 1]

Given a list of integers, return the sum of the even numbers within a list.
If the list doesn't contain any even numbers, return 0. The list will contain
at most 1000 integers between -1000 and 1000 (inclusive).

Examples
------------------------------------
Input: [0, 2, 3, 4]
Output: 6
Explanation: All numbers are even except 3, so the total is 0+2+4 = 6
------------------------------------
Input: [-11, 3, 7]
Output: 0
Explanation: All numbers are odd, so return 0
------------------------------------
*/

func EvenSum(nums []int) int {
	// TODO: Your code here
	result := 0

	for _, i := range nums {
		if i%2 == 0 {
			result += i
		}
	}

	if result != 0 {
		return result
	} else {
		return 0
	}
}
