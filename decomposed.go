package main

import "fmt"

const maxDecomposedSize = 6000

type DecomposedInteger [maxDecomposedSize]uint16

func NewDecomposedInteger(n uint16) DecomposedInteger {
	started := n
	var ret DecomposedInteger
	for i := uint16(2); i < maxDecomposedSize; i++ {
		if n == 1 {
			break
		}
		//fmt.Println(n, i)
		if n%i == 0 {
			n /= i
			ret[i] += 1
			i = 1

			continue
		}
	}
	if n > maxDecomposedSize {
		panic(fmt.Sprintf("Cannot handle integer %d with only %d slots", started, maxDecomposedSize))
	}
	return ret
}

func (d DecomposedInteger) Pow(exp uint16) DecomposedInteger {
	var ret DecomposedInteger
	for i := 0; i < maxDecomposedSize; i++ {
		ret[i] = d[i] * exp
	}
	return ret
}
