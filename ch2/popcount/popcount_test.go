package popcount

import (
	"math/rand"
	"testing"
)

var tests = []struct {
	num      uint64
	popcount int
}{
	{8, 1},
	{1024, 1},
	{1537, 3},
}

const benchmarkTestValue uint64 = 1232134
const numRandomTests int = 1000

var randomTests [numRandomTests]uint64

func init() {
	for i := 0; i < numRandomTests; i++ {
		randomTests[i] = uint64(rand.Int63())
	}
}

func TestPopCount(t *testing.T) {
	for _, test := range tests {
		//Testint Popcount()
		if PopCount(test.num) != test.popcount {
			t.Errorf("PopCount(%v) != %d, expected %d", test.num, PopCount(test.num), test.popcount)
		}

		//Testing PopCountEx23
		if PopCountEx23(test.num) != test.popcount {
			t.Errorf("PopCount(%v) != %d, expected %d", test.num, PopCountEx23(test.num), test.popcount)
		}

	}
}

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(benchmarkTestValue)
	}

}

func BenchmarkPopCountEx23(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountEx23(benchmarkTestValue)
	}
}
