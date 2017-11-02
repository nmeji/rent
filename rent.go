package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/nmeji/rent/math"
	"go.uber.org/zap"
)

type Tenant struct {
	Name       string  `json:"name"`
	DaysStayed int     `json:"days"`
	Rent       float64 `json:"rent"`
}

type Summary struct {
	MonthlyTotal     float64  `json:"monthly_total"`
	MonthlyUtilities float64  `json:"monthly_utilities"`
	Tenants          []Tenant `json:"tenants"`
}

var AverageStay func(d, t int) float64 = math.Avg

func daysStayed(t []Tenant) []int {
	ds := make([]int, len(t))
	for i, tenant := range t {
		ds[i] = tenant.DaysStayed
	}
	return ds
}

func Calculate(summary *Summary) {
	monthly := summary.MonthlyTotal / float64(len(summary.Tenants))
	utilities := summary.MonthlyUtilities
	totalStay := math.SumInt(daysStayed(summary.Tenants)) // sum of all stays
	for i, tenant := range summary.Tenants {
		avg := AverageStay(tenant.DaysStayed, totalStay)
		summary.Tenants[i].Rent = math.TruncateFloat(avg*utilities+monthly, 4)
	}
}

func printUsage() {
	fmt.Println("usage: rent monthly_rent total_utilities head_count")
}

func Rent(w http.ResponseWriter, r *http.Request) {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	defer r.Body.Close()
	s := new(Summary)
	err := json.NewDecoder(r.Body).Decode(s)
	if err != nil {
		logger.Error("problem decoding JSON", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	Calculate(s)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(s); err != nil {
		w.WriteHeader(500)
		return
	}
}

func main() {
	if len(os.Args) < 4 {
		printUsage()
		os.Exit(0)
	}

	var monthlyTotal float64
	if total, err := strconv.ParseFloat(os.Args[1], 64); err == nil {
		monthlyTotal = total
	} else {
		fmt.Println("Please provide the correct monthly rent")
		os.Exit(1)
	}

	var utilsCost float64
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
	tenants := make([]Tenant, headCount)
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
		tenants[i] = Tenant{Name: name, DaysStayed: days}
	}

	summary := &Summary{
		MonthlyTotal:     monthlyTotal,
		MonthlyUtilities: utilsCost,
		Tenants:          tenants,
	}
	Calculate(summary)
	report, err := json.MarshalIndent(summary, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("\n%s\n", string(report))
}
