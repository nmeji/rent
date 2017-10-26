package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/nmeji/rent/math"
)

type Tenant struct {
	Name       string  `json:"name"`
	DaysStayed int     `json:"days"`
	Rent       float64 `json:"rent"`
}

var AverageStay func(d, t int) float64 = math.Avg

func SummaryOfContributions(totalRent float64, daysStayed []int) []float64 {
	totalStay := math.SumInt(daysStayed) // sum of all stays
	rent := make([]float64, len(daysStayed))
	for i, ds := range daysStayed {
		avg := AverageStay(ds, totalStay)
		rent[i] = math.TruncateFloat(avg*totalRent, 4)
	}
	return rent
}

func printUsage() {
	fmt.Println("usage: rent monthly_rent total_utilities head_count")
}

func main() {
	if len(os.Args) < 4 {
		printUsage()
		os.Exit(0)
	}

	monthlyRent := float64(0)
	if total, err := strconv.ParseFloat(os.Args[1], 64); err == nil {
		monthlyRent = total
	} else {
		fmt.Println("Please provide the correct monthly rent")
		os.Exit(1)
	}

	utilsCost := float64(0)
	if uc, err := strconv.ParseFloat(os.Args[2], 64); err == nil {
		utilsCost = uc
	} else {
		fmt.Println("Please provide the correct utilities cost (electricity + water)")
		os.Exit(1)
	}

	headCount := 0
	if count, err := strconv.Atoi(os.Args[3]); err == nil {
		headCount = count
	} else {
		fmt.Println("Please provide the correct head count")
		os.Exit(1)
	}

	scanner := bufio.NewScanner(os.Stdin)
	tenant := make([]Tenant, headCount)
	daysStayed := make([]int, headCount)
	for i := 0; i < headCount; i++ {
		name := fmt.Sprintf("Tenant #%d", i+1)
		fmt.Printf("[%s] Name: ", name)
		if scanner.Scan() {
			name = scanner.Text()
		}
		days := 0
		fmt.Print("# of days stayed: ")
		if scanner.Scan() {
			if num, err := strconv.Atoi(scanner.Text()); err == nil {
				days = num
			}
		}
		tenant[i] = Tenant{Name: name, DaysStayed: days}
		daysStayed[i] = days
	}
	rents := SummaryOfContributions(utilsCost, daysStayed)
	for i, rent := range rents {
		tenant[i].Rent = monthlyRent + rent
	}
	summary, err := json.MarshalIndent(tenant, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("\n%s\n", string(summary))
}
