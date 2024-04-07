package main

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

// https://leetcode.cn/problems/first-bad-version/description/
func isBadVersion(version int) bool

func firstBadVersion(n int) int {
	start := 1
	end := n
	for start <= end {
		mid := (start + end) / 2
		if isBadVersion(mid) {
			if mid == 0 || mid-1 >= 0 && !isBadVersion(mid-1) {
				return mid
			}
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return -1
}
