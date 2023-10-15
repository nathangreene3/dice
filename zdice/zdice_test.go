package zdice

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
	// z is z*Hz, where Hz is the zth harmonic number. Here, we roll
	// each die 2*d*Hz to help ensure we observe the minimum and
	// maximum result of 0 and z-1, respectively.
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

	dice := []ZDie{1, Z4, Z6, Z8, Z10, Z12, Z20, Z100}
	for _, z := range dice {
		var (
			maxi int = 2 * int(float64(z)*harmonicNumber(int(z))) // 2*z*Hz
			min  int = 32 << (^uint(0) >> 63)
			max  int = -min
		)

		for i := 0; i < maxi; i++ {
			r := z.Roll(1)
			if r < min {
				min = r
			}

			if max < r {
				max = r
			}
		}

		if min != 0 {
			t.Errorf("expected min %d\nreceived %d", 1, min)
		}

		if max != int(z-1) {
			t.Errorf("expected max %d\nreceived %d", z, max)
		}
	}
}

func TestParseString(t *testing.T) {
	g := NewWithT(t)

	{
		goodDiceStrings := map[ZDie]string{
			0:    "Z0",
			1:    "Z1",
			Z4:   "Z4",
			Z6:   "Z6",
			Z8:   "Z8",
			Z10:  "Z10",
			Z12:  "Z12",
			Z20:  "Z20",
			Z100: "Z100",
		}

		for z, s := range goodDiceStrings {
			rz, err := Parse(s)
			g.Expect(err).To(BeNil())
			g.Expect(rz).To(Equal(z))

			rs := z.String()
			g.Expect(rs).To(Equal(s))
		}
	}

	{
		badDiceStrings := []string{
			"",
			"Z",
			"0",
			" Z0",
			"Z0 ",
		}

		for _, s := range badDiceStrings {
			z, err := Parse(s)
			g.Expect(errors.Is(err, ErrInvalidFmt)).To(BeTrue())
			g.Expect(z).To(BeZero())
		}
	}
}
