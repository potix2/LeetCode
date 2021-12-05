package main

import "sort"

func findEvenNumbers(digits []int) []int {

	ret := []int {}
	table := map[int]bool {}
	for i := range digits {
		if digits[i] == 0 {
			continue
		}

		for j := range digits {
			if i == j {
				continue
			}

			for k := range digits {
				if j == k || i == k {
					continue
				}

				ijk := digits[i] * 100 + digits[j] * 10 + digits[k]
				if _, ok := table[ijk]; !ok && ijk % 2 == 0 {
					table[ijk] = true
					ret = append(ret, ijk)
				}
			}
		}
	}
	sort.Ints(ret)
	return ret
}