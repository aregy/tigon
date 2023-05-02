package main

func BitString(n int) string {
	u := uint(n)
	k := uint(1)
	repr := ""
	for i := 0; i < 32; i++ {
		if i > 0 && i%4 == 0 {
			repr = " " + repr
		}
		if u&k == k {
			repr = "1" + repr
		} else {
			repr = "0" + repr
		}
		k <<= 1
	}
	return repr
}
func ReverseBits(n int) int {
	u := uint(n)
	c := uint(0)
	for i := 0; i < 32; i++ {
		if u%2 == 1 {
			c += 1
		}
		u /= 2
		c *= 2
	}
	return int(c)
}

func main() {

}
