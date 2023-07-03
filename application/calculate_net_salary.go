package application

import "fmt"

type CalculateNetSalary struct {
	RegionRepository RegionRepository
	StateRepository  StateRepository
}

type CalculateNetSalaryInput struct {
	GrossSalary float32
	Region      string
}

func (c CalculateNetSalary) Execute(input CalculateNetSalaryInput) float32 {

	stateRetentions := []float32{}
	regionRetentions := []float32{}

	stateRanges := c.StateRepository.getRangesForState()
	regionRanges := c.RegionRepository.getRangesForRegion(input.Region)

	for i := 0; i < len(stateRanges); i++ {
		retentionRange := stateRanges[i]

		retention := retentionRange.retentionOverSalary(input.GrossSalary)

		fmt.Printf("retention state: %f \n", retention)
		stateRetentions = append(stateRetentions, retention)
	}

	for i := 0; i < len(regionRanges); i++ {
		retentionRange := regionRanges[i]

		retention := retentionRange.retentionOverSalary(input.GrossSalary)

		fmt.Printf("retention region: %f \n", retention)
		regionRetentions = append(regionRetentions, retention)
	}

	return (input.GrossSalary * 0.94) - (sumArrayValues(stateRetentions) + sumArrayValues(regionRetentions))
}

func sumArrayValues(arr []float32) float32 {
	res := float32(0.0)
	for i := 0; i < len(arr); i++ {
		res += arr[i]
	}
	return res
}
