package chapter_5

import "math"

func EnumerateAllPrimes(x int) []int {
	if x < 2 {
		return []int{}
	}

	size := int(math.Floor(0.5*float64((x-3))) + 1)
	var primes []int
	primes = append(primes, 2)

	isPrime := make([]bool, size)
	for i := range isPrime {
		isPrime[i] = true
	}
	for i := 0; i < size; i++ {
		if isPrime[i] {
			p := (i * 2) + 3
			primes = append(primes, p)

			for j := 2*i*i + 6*i + 3; j < size; j += p {
				isPrime[j] = false
			}
		}
	}
	return primes
}
