package application

type RegionRepository struct {
}

func (r RegionRepository) getRangesForRegion(region string) []retentionRange {

	ranges := [7]retentionRange{
		{
			from:      0,
			to:        12450,
			retention: 9.5,
		},
		{
			from:      12450,
			to:        20199,
			retention: 12,
		},
		{
			from:      20200,
			to:        35199,
			retention: 15,
		},
		{
			from:      35200,
			to:        59999,
			retention: 18.5,
		},
		{
			from:      60000,
			to:        299999,
			retention: 22.5,
		},
		{
			from:      300000,
			to:        35199,
			retention: 24.5,
		},
	}
	return ranges[:]
}
