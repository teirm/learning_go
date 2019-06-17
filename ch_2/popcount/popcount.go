package popcount

// pc[i] is the population count of i.
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

// PopCount2 returns the population count (number of set bits) of x
// with a loop
func PopCount2(x uint64) int {
	var i uint
	var sum int
	for i < 8 {
		sum += int(pc[byte(x>>(i*8))])
		i++
	}
	return sum
}

// PopCount3 returns the population count (number of set bits) of x
// by shifting the value through all positions
func PopCount3(x uint64) int {
	var sum int
	for i := 0; i < 64; i++ {
		sum += int(x & 1)
		x = x >> 1
	}
	return sum
}

// PopCount4 returns the population count (number of set bits) of x
// by clearing the rightmost 1 repeatedly
func PopCount4(x uint64) int {
	var sum int
	for x != 0 {
		x = x & (x - 1)
		sum++
	}
	return sum
}
