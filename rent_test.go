package main_test

import (
	"testing"

	. "github.com/franela/goblin"
	rent "github.com/nmeji/rent"
	"github.com/nmeji/rent/math"
)

func allRents(s *rent.Summary) []float64 {
	r := make([]float64, len(s.Tenants))
	for i, t := range s.Tenants {
		r[i] = t.Rent
	}
	return r
}

func TestApp(t *testing.T) {
	g := Goblin(t)
	g.Describe("App", func() {
		g.It("should be able to compute average stay", func() {
			d := 24  // given the days stayed within the month
			t := 113 // and the total number of all stay
			avg := func(x, y int) float64 {
				return float64(x) / float64(y)
			}
			g.Assert(rent.AverageStay(d, t)).Equal(avg(d, t))
		})
		g.It("should sum all average stay to 1", func() {
			h := []int{24, 20, 20, 20, 20, 15} // # of stay per head
			t := math.SumInt(h)                // sum of all stays
			avg := make([]float64, len(h))
			for i, s := range h {
				avg[i] = rent.AverageStay(s, t)
			}
			fs := math.TruncateFloat(math.SumFloat(avg), 6)
			g.Assert(fs).Equal(float64(1))
		})
		g.It("should sum all per head rent to rent total", func() {
			u := 21785.75 // given total monthly utilities
			t := []rent.Tenant{
				{
					DaysStayed: 24,
				},
				{
					DaysStayed: 20,
				},
				{
					DaysStayed: 20,
				},
				{
					DaysStayed: 20,
				},
				{
					DaysStayed: 20,
				},
				{
					DaysStayed: 15,
				},
			}
			s := &rent.Summary{
				MonthlyTotal:     float64(0),
				MonthlyUtilities: u,
				Tenants:          t,
			}
			rent.Calculate(s)
			r := allRents(s)
			total := math.SumFloat(r)
			g.Assert(u).Equal(total)
		})
	})
}
