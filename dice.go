package dice

import (
	"math/rand"
	"strconv"
)

// Die represents a game die. A die with n faces is numbered from
// 1, 2, ..., n. Note `var d Die` is not defined.
type Die uint

// Commonly used dice
const (
	D4   Die = 4
	D6   Die = 6
	D8   Die = 8
	D10  Die = 10
	D12  Die = 12
	D20  Die = 20
	D100 Die = 100
)

// Max returns the maximum result of rolling a die a given number of
// times.
func (d Die) Max(count int) int {
	var m int
	for ; count > 0; count-- {
		if r := rand.Intn(int(d)) + 1; r > m {
			m = r
		}
	}

	return m
}

// Min returns the maximum result of rolling a die a given number of
// times.
func (d Die) Min(count int) int {
	const maxInt = 1<<(32<<(^uint(0)>>63)-1) - 1 // TODO: Document this

	if count == 0 {
		return 0
	}

	m := maxInt
	for ; count > 0; count-- {
		if r := rand.Intn(int(d)) + 1; r < m {
			m = r
		}
	}

	return m
}

// Parse returns a die parsed from a string.
func Parse(s string) (Die, error) {
	if len(s) < 2 || s[0] != 'D' {
		err := strconv.NumError{
			Func: "Parse",
			Num:  s,
			Err:  strconv.ErrSyntax,
		}

		return 0, &err
	}

	d, err := strconv.Atoi(s[1:])
	return Die(d), err
}

// Roll returns the sum of a given number of dice rolls.
func (d Die) Roll(count int) int {
	if count < 1 {
		return 0
	}

	r := count
	for ; count > 0; count-- {
		r += rand.Intn(int(d))
	}

	return r
}

// Roll returns the sum of several dice rolls.
func Roll(d ...Die) int {
	r := len(d)
	for i := 0; i < len(d); i++ {
		r += rand.Intn(int(d[i]))
	}

	return r
}

// String returns a representation of a die.
func (d Die) String() string {
	return "D" + strconv.Itoa(int(d))
}
