package main_test

import (
	"testing"

	. "github.com/franela/goblin"
	"github.com/nmeji/rent/math"
	rent "github.com/nmeji/rent"
)

func TestApp(t *testing.T) {
	g := Goblin(t)
	g.Describe("App", func() {
		g.It("should be able to compute average stay", func() {
			d := 24 // given the days stayed within the month
			t := 113 // and the total number of all stay
			avg := func(x,y int) float64 {
				return float64(x)/float64(y)
			}
			g.Assert(rent.AverageStay(d,t)).Equal(avg(d,t))
		})
		g.It("should sum all average stay to 1", func() {
			h := []int{24,20,20,20,20,15} // # of stay per head
			t := math.SumInt(h) // sum of all stays
			avg := make([]float64, len(h))
			for i, s := range h {
				avg[i] = rent.AverageStay(s, t)
			}
			fs := math.TruncateFloat(math.SumFloat(avg), 6)
			g.Assert(fs).Equal(float64(1))
		})
		g.It("should sum all per head rent to rent total", func() {
			tr := 21785.75 // given total monthly rent
			h := []int{24,20,20,20,20,15}  // # of stay per head
			hr := rent.SummaryOfContributions(tr, h)
			total := math.SumFloat(hr)
			g.Assert(total).Equal(tr)
		})
	})
}
