package main

import (
	"fmt"
	"sort"
)

func main() {
	ts := []int{-1, -12, 14, -6, 4, -11, 2, -7, 13, -15, -1, 11, -15, -15, 13, -9, -11, -10, -7, -13, 7, 9, -13, -3, 10, 1, -5, 12, 11, -9, 2, -4, -6, -7, 5, 5, -2, 14, 13, -12, 14, -3, 14, 14, -11, 5, 12, -6, 10, -9, -4, -6, -15, -9, -1, 2, -1, 9, -9, -5, -15, 8, -2, -6, 9, 10, 7, 14, 9, 5, -13, 9, -12, 8, 8, -8, -2, -2, 1, -15, -4, 5, -13, -4, 3, 1, 5, -13, -13, 11, -11, 10, 5, 3, 11, -9, -2, 3, -10, -7, -6, 14, 9, -11, 7, 2, -4, 13, 7, 5, 13, -12, -13, 7, -9, 5, -7, 11, -1, -12, 8, 3, 13, -10, 13, 1, -4, 6, 7, 8, -6, -8, -23, 8, 5, 3, 9, 12, -9, 8, -9, 13, -10, -14, -13, 5, -15, -4, 2, 8, -11, -6, 12, 9, -15, 13, 11, 13, 13, 6, -12, -15, -4, -6, 0, -14, 5, -14, 5, 3, 2, 4, 2, 7, 5, 4, -10, -3, 7, 7, -9, 4, -14, 10, -2, -13, 8, -6, 7, -1, 7, 11, -9, -12, -10, 6, 12, 10, 7, 2, -9, -6, 13, 8, 9, 3, -11, 14, -14, 11, -2, 14, 0, -1, 1, 6, -7, -5, 7, -14, 9, 0, 4, 7, -5, 1, -2, 14, -3, 12, -6, -5, 14, -8, -12, 0, 3, -8, -1}
	fmt.Println(threeSum(ts))
}

func threeSum(nums []int) [][]int {

	sol := [][]int{}

	if len(nums) < 3 {
		return [][]int{}
	}

	comb := combinations(len(nums), 3)

	//fmt.Println(comb)
	for i := range comb {
		var sum int
		sumSet := make([]int, 0)
		for j := range comb[i] {
			sum += nums[comb[i][j]]
			sumSet = append(sumSet, nums[comb[i][j]])
		}

		if sum == 0 {
			sort.Ints(sumSet)
			sol = append(sol, sumSet)
		}
	}

	return remDup(sol)
}

func remDup(intSlice [][]int) [][]int {
	keys := make(map[[3]int]bool)
	uniq := [][]int{}

	//fmt.Println(keys)
	for _, entry := range intSlice {
		arrEntry := [3]int{}
		copy(arrEntry[:], entry)

		//fmt.Println(keys[arrEntry])

		if keys[arrEntry] == false {
			uniq = append(uniq, entry)
			keys[arrEntry] = true
		}
	}
	return uniq
}

func combinations(n, k int) [][]int {
	combins := binomial(n, k)
	data := make([][]int, combins)
	if len(data) == 0 {
		return data
	}
	data[0] = make([]int, k)
	for i := range data[0] {
		data[0][i] = i
	}
	for i := 1; i < combins; i++ {
		next := make([]int, k)
		copy(next, data[i-1])
		nextCombination(next, n, k)
		data[i] = next
	}
	return data
}

func nextCombination(s []int, n, k int) {
	for j := k - 1; j >= 0; j-- {
		if s[j] == n+j-k {
			continue
		}
		s[j]++
		for l := j + 1; l < k; l++ {
			s[l] = s[j] + l - j
		}
		break
	}
}

func binomial(n, k int) int {

	// (n,k) = (n, n-k)
	if k > n/2 {
		k = n - k
	}
	b := 1
	for i := 1; i <= k; i++ {
		b = (n - k + i) * b / i
	}
	return b
}
