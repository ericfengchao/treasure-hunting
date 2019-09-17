package models

import (
	"math/rand"
	"time"
)

func generateNUniqueRandomNumbers(n int) []int {
	rand.Seed(time.Now().Unix())
	return rand.Perm(n)
}
