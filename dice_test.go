package dice

import (
	"errors"
	"math/rand"
	"testing"
	"time"

	. "github.com/onsi/gomega"
)

// Testing random things is hard.

func TestMaxMinObserved(t *testing.T) {
	rand.Seed(time.Now().Unix())

	// The expected number of rolls to receive all roll results of die
	// d is d*Hd, where Hd is the dth harmonic number. Here, we roll
	// each die 2*d*Hd to help ensure we observe the minimum and
	// maximum result of 1 and d, respectively.
	//
	// https://en.wikipedia.org/wiki/Coupon_collector%27s_problem
	// https://en.wikipedia.org/wiki/Harmonic_number

	// harmonicNumber returns the nth harmonic number. That is, it returns
	// 	1 + 1/2 + 1/3 + ... + 1/n,
	// where n is in N.
	harmonicNumber := func(n int) float64 {
		var hn float64
		for i := 1; i <= n; i++ {
			hn += 1.0 / float64(i)
		}

		return hn
	}

	dice := []Die{1, D4, D6, D8, D10, D12, D20, D100}
	for _, d := range dice {
		var (
			maxi int = 2 * int(float64(d)*harmonicNumber(int(d))) // 2*d*Hd
			min  int = 32 << (^uint(0) >> 63)
			max  int = -min
		)

		for i := 0; i < maxi; i++ {
			r := d.Roll(1)
			if r < min {
				min = r
			}

			if max < r {
				max = r
			}
		}

		if min != 1 {
			t.Errorf("expected min %d\nreceived %d", 1, min)
		}

		if max != int(d) {
			t.Errorf("expected max %d\nreceived %d", d, max)
		}
	}
}

func TestParseString(t *testing.T) {
	g := NewWithT(t)

	{
		goodDiceStrings := map[Die]string{
			0:    "D0",
			1:    "D1",
			D4:   "D4",
			D6:   "D6",
			D8:   "D8",
			D10:  "D10",
			D12:  "D12",
			D20:  "D20",
			D100: "D100",
		}

		for d, s := range goodDiceStrings {
			rd, err := Parse(s)
			g.Expect(err).To(BeNil())
			g.Expect(rd).To(Equal(d))

			rs := d.String()
			g.Expect(rs).To(Equal(s))
		}
	}

	{
		badDiceStrings := []string{
			"",
			"D",
			"0",
			" D0",
			"D0 ",
		}

		for _, s := range badDiceStrings {
			d, err := Parse(s)
			g.Expect(errors.Is(err, ErrInvalidFmt)).To(BeTrue())
			g.Expect(d).To(BeZero())
		}
	}
}
