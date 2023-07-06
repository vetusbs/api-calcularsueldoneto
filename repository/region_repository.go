package repository

import (
	"calcularsueldoneto/models"
)

type RegionRepository struct {
	regionRetentions map[string][]Bucket
}

func (r RegionRepository) GetRangesForRegion(region string) []models.RetentionRange {

	retentions := r.regionRetentions[region]

	arr := []models.RetentionRange{}
	for index, element := range retentions {
		if index == 0 {
			arr = append(arr, models.CreateRetention(0, element.UpTo, element.Retention))
		} else {
			bucket := models.CreateRetention(retentions[index-1].UpTo, element.UpTo, element.Retention)
			arr = append(arr, bucket)
		}
	}

	return arr
}
