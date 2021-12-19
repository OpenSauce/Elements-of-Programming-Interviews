package chapter_5

import (
	"math/rand"
	"time"
)

func SampleOfflineData(data []int, k int) []int {
	x1 := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < k; i++ {
		rndIndex := x1.Intn(len(data))
		data[i], data[rndIndex] = data[rndIndex], data[i]
	}
	return data
}
