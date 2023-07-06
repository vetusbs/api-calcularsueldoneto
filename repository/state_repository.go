package repository

import (
	"calcularsueldoneto/models"
)

type StateRepository struct {
	retentions []Bucket
}

func (r StateRepository) getRetentions() []Bucket {
	return r.retentions
}

func (r StateRepository) GetRangesForState() []models.RetentionRange {

	arr := []models.RetentionRange{}

	for index, element := range r.retentions {
		if index == 0 {
			arr = append(arr, models.CreateRetention(0, element.UpTo, element.Retention))
		} else {
			bucket := models.CreateRetention(r.retentions[index-1].UpTo, element.UpTo, element.Retention)
			arr = append(arr, bucket)
		}
	}

	return arr
}
