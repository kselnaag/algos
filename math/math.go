package math

import (
	I "github.com/kselnaag/algos/types"
)

func HashPirson16(mess []byte) uint16 {
	T := [256]byte{
		130, 4, 133, 49, 108, 178, 125, 95, 35, 126, 41, 129, 229, 48, 6, 94, 206,
		69, 20, 194, 236, 79, 156, 67, 100, 239, 152, 149, 93, 91, 56, 8, 183, 165,
		42, 148, 114, 59, 57, 5, 112, 151, 54, 97, 109, 145, 228, 196, 250, 104,
		169, 107, 86, 64, 98, 181, 200, 58, 199, 70, 138, 179, 60, 249, 34, 123,
		30, 22, 124, 240, 201, 132, 218, 21, 74, 83, 39, 223, 73, 88, 136, 27, 0,
		10, 89, 51, 215, 251, 255, 3, 235, 241, 19, 102, 71, 38, 166, 220, 110, 23,
		232, 25, 172, 210, 142, 211, 121, 242, 75, 208, 195, 203, 226, 253, 176, 17,
		66, 158, 231, 237, 99, 254, 173, 221, 117, 139, 213, 90, 85, 45, 187, 84, 92,
		44, 164, 247, 122, 32, 127, 177, 170, 155, 111, 185, 171, 61, 76, 184, 234, 192,
		16, 106, 160, 204, 153, 161, 186, 131, 28, 137, 37, 216, 248, 55, 72, 50, 26,
		46, 53, 224, 7, 217, 189, 120, 219, 167, 119, 11, 252, 65, 135, 96, 222, 68,
		144, 214, 227, 101, 207, 103, 212, 175, 157, 141, 168, 82, 163, 47, 52, 15, 113,
		230, 245, 116, 43, 80, 246, 33, 198, 197, 146, 193, 13, 31, 24, 143, 12, 18, 118,
		14, 62, 154, 78, 81, 134, 162, 105, 63, 244, 77, 190, 209, 150, 233, 159, 202, 191,
		40, 87, 180, 188, 36, 238, 9, 140, 128, 147, 174, 1, 2, 182, 243, 29, 115, 205, 225,
	}
	result := uint16(0)
	mlen := len(mess)
	hash := T[mlen%256]
	for j := 0; j < 2; j++ {
		for i := 0; i < mlen; i++ {
			hash = T[int(hash^mess[i])]
		}
		result += uint16(hash) << (8 * j)
	}
	return result
}

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
