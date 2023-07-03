package application

type StateRepository struct {
}

func (r StateRepository) getRangesForState() []retentionRange {

	ranges := [7]retentionRange{
		{
			from:      0,
			to:        12960.45,
			retention: 8.5,
		},
		{
			from:      12960.45,
			to:        18433.20,
			retention: 10.7,
		},
		{
			from:      18433.20,
			to:        34360.50,
			retention: 12.8,
		},
		{
			from:      34360.50,
			to:        55596.90,
			retention: 17.4,
		},
		{
			from:      55596.90,
			to:        100000000,
			retention: 20.5,
		},
	}
	return ranges[:]
}
