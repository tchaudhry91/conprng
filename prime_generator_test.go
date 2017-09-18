package conprng

import (
	"runtime"
	"sort"
	"testing"
)

func TestReturnLength(t *testing.T) {
	testLengths := []int{100, 599, 1000, 2322}
	for _, length := range testLengths {
		lenFound := len(GeneratePrimes(length))
		if length != lenFound {
			t.Logf("Found %d primes instead of the requested %d", lenFound, length)
			t.Fail()
		}
	}
}

func TestPrimeAvailability100(t *testing.T) {
	primes := GeneratePrimes(100)
	presentNums := []int{2, 3, 17, 457, 331, 281, 541}
	for _, num := range presentNums {
		index := sort.SearchInts(primes, num)
		found := true
		if index < len(primes) && primes[index] != num {
			found = false
		} else if index >= len(primes) {
			found = false
		} else {
			found = true
		}
		if found == false {
			t.Logf("Did not find %d", num)
			t.Fail()
		}
	}
}

func TestCompositeAbsence100(t *testing.T) {
	primes := GeneratePrimes(100)
	presentNums := []int{4, 600, 242, 92, 99, 287, 498}
	for _, num := range presentNums {
		index := sort.SearchInts(primes, num)
		found := true
		if index < len(primes) && primes[index] != num {
			found = true
		} else if index >= len(primes) {
			found = true
		} else {
			found = false
		}
		if found == false {
			t.Logf("Found %d", num)
			t.Fail()
		}
	}
}

func BenchmarkCoresQuarter(b *testing.B) {
	runtime.GOMAXPROCS(runtime.NumCPU() / 4)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GeneratePrimes(10000)
	}
}

func BenchmarkCoresFull(b *testing.B) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GeneratePrimes(10000)
	}
}
