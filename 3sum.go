package main

import (
	"fmt"
	"sort"

	"github.com/gonum/stat/combin"
)

func main() {
	ts := []int{-1, 0, 1, 2, -1, -4, 2}
	fmt.Println(threeSum(ts))
}

func threeSum(nums []int) [][]int {

	sol := [][]int{}
	comb := combin.Combinations(len(nums), 3)

	fmt.Println(comb)
	for i := range comb {
		sum := 0
		sumSet := make([]int, 0)
		for j := range comb[i] {
			sum += nums[comb[i][j]]
			sumSet = append(sumSet, nums[comb[i][j]])
		}

		sort.Ints(sumSet)
		if sum == 0 {
			sol = append(sol, sumSet)
		}
	}
	sol = remDup(sol)
	return [][]int{sol[1], sol[0]}
}

func remDup(intSlice [][]int) [][]int {
	keys := make(map[[3]int]bool)
	uniq := [][]int{}

	fmt.Println(keys)
	for _, entry := range intSlice {
		arrEntry := [3]int{}
		copy(arrEntry[:], entry)

		fmt.Println(keys[arrEntry])

		if keys[arrEntry] == false {
			uniq = append(uniq, entry)
			keys[arrEntry] = true
		}
	}
	return uniq
}

func Combinations(n, k int) [][]int {
	combins := Binomial(n, k)
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

func Binomial(n, k int) int {

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
