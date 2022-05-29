package zdice

import (
	"math/rand"
	"strconv"
)

// ZDie represents a zero-based die. A die with n faces is numbered
// 0, 1, ..., n-1, where n > 0. Note `var z ZDie` is not defined.
type ZDie uint

// Commonly used dice
const (
	Z4   ZDie = 4
	Z6   ZDie = 6
	Z8   ZDie = 8
	Z10  ZDie = 10
	Z12  ZDie = 12
	Z20  ZDie = 20
	Z100 ZDie = 100
)

// Max returns the maximum result of rolling a die a given number of
// times.
func (z ZDie) Max(count int) int {
	var m int
	for ; count > 0; count-- {
		if r := rand.Intn(int(z)); r > m {
			m = r
		}
	}

	return m
}

// Min returns the maximum result of rolling a die a given number of
// times.
func (z ZDie) Min(count int) int {
	const maxInt = 1<<(32<<(^uint(0)>>63)-1) - 1 // TODO: Document this.

	if count == 0 {
		return 0
	}

	m := maxInt
	for ; count > 0; count-- {
		if r := rand.Intn(int(z)); r < m {
			m = r
		}
	}

	return m
}

// Parse returns a die parsed from a string.
func Parse(s string) (ZDie, error) {
	if len(s) < 2 || s[0] != 'Z' {
		err := strconv.NumError{
			Func: "Parse",
			Num:  s,
			Err:  strconv.ErrSyntax,
		}

		return 0, &err
	}

	z, err := strconv.Atoi(s[1:])
	return ZDie(z), err
}

// Roll returns the sum of a given number of dice rolls.
func (z ZDie) Roll(count int) int {
	var r int
	for ; count > 0; count-- {
		r += rand.Intn(int(z))
	}

	return r
}

// Roll returns the sum of several dice rolls.
func Roll(z ...ZDie) int {
	var r int
	for i := 0; i < len(z); i++ {
		r += rand.Intn(int(z[i]))
	}

	return r
}

// String returns a representation of a die.
func (z ZDie) String() string {
	return "Z" + strconv.Itoa(int(z))
}
