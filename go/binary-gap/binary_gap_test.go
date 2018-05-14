package binarygap

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type example struct {
	n     int
	zeros int
}

func (e example) String() string {
	return fmt.Sprintf("original: %d\nbinary: %b\nzeros count: %d", e.n, e.n, e.zeros)
}

var testCases = []example{
	{1, 0},  //      1
	{2, 1},  //     10
	{4, 2},  //    100
	{8, 3},  //   1000
	{3, 0},  //     11
	{5, 1},  //    101
	{6, 1},  //    110
	{7, 0},  //    111
	{9, 2},  //   1001
	{10, 1}, //   1010
	{34, 3}, // 100010
}

func TestLongestZeros(t *testing.T) {
	assert := assert.New(t)

	for _, testCase := range testCases {
		assert.Equal(testCase.zeros, LongestZeros(testCase.n), testCase.String())
	}
}

func BenchmarkLongestZeros(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range testCases {
			LongestZeros(c.n)
		}
	}
}
