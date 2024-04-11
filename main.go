package main

import (
	"fmt"
	"time"
)

const FareOnePeak = 13
const FareTenPeak = 130
const FareTenOffPeak = 83
const FareMonthly = 253
const FareWeekly = 90

func main() {
	for i := 1; i <= 5; i++ {
		costPerWeek := FareOnePeak * 2 * i
		if costPerWeek > FareWeekly {
			fmt.Printf("Weekly Pass If %d Days In Office - $%d\n\n", i, (costPerWeek - FareWeekly))
			break
		}
	}

	daysToGoIn := [][]time.Weekday{
		{time.Tuesday},
		{time.Tuesday, time.Wednesday},
		{time.Tuesday, time.Wednesday, time.Monday},
		{time.Tuesday, time.Wednesday, time.Monday, time.Thursday},
		{time.Tuesday, time.Wednesday, time.Monday, time.Thursday, time.Friday},
	}

	monthArray := []time.Month{
		time.January, time.February, time.March, time.April, time.May, time.June, time.July, time.August, time.September, time.October, time.November, time.December}

	for _, month := range monthArray {
		counts := make(map[time.Weekday]int)
		for day := time.Date(2024, month, 1, 0, 0, 0, 0, time.Local); day.Local().Month() == month; day = day.AddDate(0, 0, 1) {
			counts[day.Local().Weekday()]++
		}
		for _, days := range daysToGoIn {
			numDaysIn := 0
			for _, d := range days {
				numDaysIn += counts[d]
			}
			costPerMonth := FareOnePeak * 2 * numDaysIn

			if costPerMonth > FareMonthly {
				fmt.Printf("%s Monthly Pass If %d Days In Office - $%d\n", month.String(), len(days), (costPerMonth - FareMonthly))
				break
			}
		}
	}

	// fmt.Printf("Trips per weekly: %f\n", FareWeekly/FareOnePeak)
	// fmt.Printf("Trips per monthly: %f\n", FareMonthly/FareOnePeak)
}
