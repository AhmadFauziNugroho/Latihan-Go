package util

import "fmt"

func MustPositive(n int) (int, error) {
	if n < 0 {
		return 0, fmt.Errorf("angka %d tidak boleh nol atau negatif", n)
	}
	return n, nil
}