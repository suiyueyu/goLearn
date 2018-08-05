package popcount

// The package below defines a function PopCount that returns the number of set bits, that is,
// bits whose value is 1, in a uint64 value, which is called its population count.

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// Exercise 2.3: Rewrite PopCount to use a loop instead of a single expression. Compare the per-
// formance of the two versions. (Section 11.4 shows how to compare the performance of differ-
// ent implementations systematically.)
func PopCount2(x uint64) int {
	count := 0
	var i uint = 0

	for i = 0; i < 7; i++ {
		count += int(pc[byte(x>>(uint(i*8)))])
	}

	return count
}
