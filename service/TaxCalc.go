package service

import (
	"fmt"
	"math"
)

var quickCalculator = make([]TaxMap, 0)

type TaxMap struct {
	sum             float64
	withholdingRate float64
	deduction       float64
}

func init() {
	quickCalculator = append(quickCalculator, TaxMap{
		36000, 0.03, 0,
	})
	quickCalculator = append(quickCalculator, TaxMap{
		144000, 0.1, 2520,
	})
	quickCalculator = append(quickCalculator, TaxMap{
		300000, 0.2, 16920,
	})
	quickCalculator = append(quickCalculator, TaxMap{
		420000, 0.25, 31920,
	})
	quickCalculator = append(quickCalculator, TaxMap{
		660000, 0.3, 52920,
	})
	quickCalculator = append(quickCalculator, TaxMap{
		960000, 0.35, 85920,
	})
	quickCalculator = append(quickCalculator, TaxMap{
		math.MaxFloat64, 0.45, 18920,
	})
}

// SeparateTax 单独计税
func SeparateTax(salary float64, yearEndAwards float64) {
	oneYear := (salary - salary*0.185 - 5000 - 1500) * 12
	var tax float64 = 0
	var tax2 float64 = 0
	for _, value := range quickCalculator {
		if oneYear <= value.sum {
			tax = (oneYear - value.deduction) * value.withholdingRate
			fmt.Printf("工资预扣率%f， 速算扣除数%f", value.withholdingRate, value.deduction)
			fmt.Println()
			break
		}
	}
	for _, value := range quickCalculator {
		if yearEndAwards <= value.sum {
			tax2 = yearEndAwards * value.withholdingRate
			fmt.Printf("年终奖预扣率%f", value.withholdingRate)
			fmt.Println()
			break
		}
	}
	fmt.Printf("综合计税工资扣税%f，年终奖扣税%f， 合计扣税%f", tax, tax2, tax+tax2)
	fmt.Println()
}

// ComprehensiveTax 综合计税
func ComprehensiveTax(salary float64, yearEndAwards float64) {
	oneYear := (salary - salary*0.185 - 5000 - 1500) * 12
	for _, value := range quickCalculator {
		if oneYear <= value.sum {
			tax := oneYear*value.withholdingRate - value.deduction
			tax2 := yearEndAwards * value.withholdingRate
			fmt.Printf("综合计税工资扣税%f，年终奖扣税%f， 合计扣税%f， 预扣率%f， 速算扣除数%f", tax, tax2, tax+tax2, value.withholdingRate, value.deduction)
			fmt.Println()
			break
		}
	}
}
