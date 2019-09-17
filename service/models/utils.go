package models

import (
	"math/rand"
	"time"
)

func generateNUniqueRandomNumbers(n int, max int) []int {
	res := make([]int, max)
	for i := 0; i < max; i++ {
		res[i] = i
	}
	rand.Seed(time.Now().Unix())
	for i := 0; i < n; i++ {
		r := rand.Intn(max - i)
		res[max-i-1], res[r] = res[r], res[max-i-1]

	}
	return res[:n]
}
