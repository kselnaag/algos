package math

import (
	I "github.com/kselnaag/algos/types"
)

func Harmonic(n int) float64 {
	if n < 0 {
		panic("algos.math.Harmonic(x) -> 'x' can not be negative")
	}
	sum := 0.0
	for i := 1; i <= n; i++ {
		sum += 1.0 / float64(i)
	}
	return sum
}

// Newton square root
func Sqrt(c float64) float64 {
	if c < 0 {
		panic("algos.math.Sqrt(x) -> 'x' can not be negative")
	}
	err := 1e-15
	t := c
	for Abs(t-c/t) > (err * t) {
		t = (c/t + t) / 2.0
	}
	return t
}

func Abs[T I.Signed | I.Float](val T) T {
	if val < 0 {
		return val * (-1)
	} else {
		return val
	}
}

func IsPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; (i * i) <= n; i++ {
		if (n % i) == 0 {
			return false
		}
	}
	return true
}

// Evklid NOD
func Gcd(p, q int) int {
	p, q = Abs(p), Abs(q)
	if q == 0 {
		return p
	}
	r := p % q
	return Gcd(q, r)
}

func Min[T I.Ord](x, y T) T {
	if x < y {
		return x
	} else {
		return y
	}
}

func Max[T I.Ord](x, y T) T {
	if x > y {
		return x
	} else {
		return y
	}
}
