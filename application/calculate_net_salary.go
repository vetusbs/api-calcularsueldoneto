package application

import (
	"calcularsueldoneto/models"
	"calcularsueldoneto/repository"
	"fmt"
)

type CalculateNetSalary struct {
	RegionRepository repository.RegionRepository
	StateRepository  repository.StateRepository
}

type CalculateNetSalaryInput struct {
	GrossSalary float32
	Region      string
	Children    int
	Babies      int
}

type CalculateNetSalaryOutput struct {
	MonthlyNet          float32
	YearlyNet           float32
	TotalWithholdings   float32
	TotalSocialSecurity float32
	ExtraMonth          float32
}

func (c CalculateNetSalary) Execute(input CalculateNetSalaryInput) CalculateNetSalaryOutput {

	stateRanges := c.StateRepository.GetRangesForState()
	regionRanges := c.RegionRepository.GetRangesForRegion(input.Region)

	bonus := calculateBonus(input)                                                // kids, parents and disabilities
	workerSSWithholdings := calculateSocialSecurityWithHolding(input.GrossSalary) // Employee social security
	taxBase := (input.GrossSalary * 0.94)                                         // base amount for taxes

	bonusRetention := calculateRetentions(regionRanges, bonus)
	fmt.Printf("retention bonus: %f \n", bonusRetention)

	stateRetention := calculateRetentions(stateRanges, taxBase)
	fmt.Printf("retention state: %f \n", stateRetention)

	regionRetention := calculateRetentions(regionRanges, taxBase)
	fmt.Printf("retention region: %f \n", regionRetention)

	totalWithholdings := regionRetention + stateRetention - bonusRetention
	netYear := taxBase - totalWithholdings
	netMonth := netYear / 12
	return CalculateNetSalaryOutput{
		MonthlyNet:          netMonth,
		YearlyNet:           netYear,
		TotalWithholdings:   totalWithholdings,
		TotalSocialSecurity: workerSSWithholdings,
		ExtraMonth:          float32(0),
	}
}

func calculateRetentions(retentionRanges []models.RetentionRange, input float32) float32 {
	var totalRetentions []float32

	for i := 0; i < len(retentionRanges); i++ {
		retentionRange := retentionRanges[i]

		retention := retentionRange.RetentionOverSalary(input)

		totalRetentions = append(totalRetentions, retention)
	}

	return sumArrayValues(totalRetentions)
}

type Props struct {
	childrenNumber       int
	babiesNumber         int
	disabilityPercentage float64
}

const BASE_BONUS = 5550

func calculateBonus(input CalculateNetSalaryInput) float32 {
	childrenBonus := calculateChildrenBonus(input) / 2
	babiesBonus := input.Babies * 2800 / 2
	//disabilityBonus := calculateDisabilityBonus(props.disabilityPercentage)

	return float32(BASE_BONUS) + float32(childrenBonus) + float32(babiesBonus)
}

const MIN_SOCIAL_SECURITY = float32(1260)
const MAX_SOCIAL_SECURITY = float32(4495.50)
const EMPLOYER_SS_WITHHOLDINGS = 4.7 + 1.55 + 0.1

func calculateSocialSecurityWithHolding(grossSalary float32) float32 {
	if grossSalary > MAX_SOCIAL_SECURITY*12 {
		return MAX_SOCIAL_SECURITY * 12 * (EMPLOYER_SS_WITHHOLDINGS / 100)
	} else if grossSalary < MIN_SOCIAL_SECURITY*12 {
		return MIN_SOCIAL_SECURITY * 12 * EMPLOYER_SS_WITHHOLDINGS
	} else {
		return grossSalary * EMPLOYER_SS_WITHHOLDINGS
	}
}

func calculateChildrenBonus(input CalculateNetSalaryInput) float32 {
	switch input.Children {
	case 0:
		return 2400
	case 1:
		return 2400
	case 2:
		return 2700
	case 3:
		return 4000
	case 4:
		return 4500
	default:
		return 4500
	}
}

func sumArrayValues(arr []float32) float32 {
	res := float32(0.0)
	for i := 0; i < len(arr); i++ {
		res += arr[i]
	}
	return res
}
