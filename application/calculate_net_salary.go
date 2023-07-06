package application

import (
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

func (c CalculateNetSalary) Execute(input CalculateNetSalaryInput) float32 {

	stateRetentions := []float32{}
	regionRetentions := []float32{}

	stateRanges := c.StateRepository.GetRangesForState()
	regionRanges := c.RegionRepository.GetRangesForRegion(input.Region)

	baseSalary := (input.GrossSalary * 0.94) - calculateBonus(input)

	for i := 0; i < len(stateRanges); i++ {
		retentionRange := stateRanges[i]

		retention := retentionRange.RetentionOverSalary(baseSalary)

		fmt.Printf("retention state: %f bucket %v\n", retention, retentionRange)
		stateRetentions = append(stateRetentions, retention)
	}

	for i := 0; i < len(regionRanges); i++ {
		retentionRange := regionRanges[i]

		retention := retentionRange.RetentionOverSalary(baseSalary)

		fmt.Printf("retention region: %f bucket %v\n", retention, retentionRange)
		regionRetentions = append(regionRetentions, retention)
	}

	return (input.GrossSalary * 0.94) - (sumArrayValues(stateRetentions) + sumArrayValues(regionRetentions))
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
