package models

import "fmt"

type RetentionRange struct {
	from      float32
	to        float32
	retention float32
}

func (r RetentionRange) RetentionOverSalary(salary float32) float32 {
	if r.from >= salary {
		fmt.Println(r.from)
		return 0
	} else if r.to <= salary {
		return (r.to - r.from) * r.retention / 100
	} else {
		return (salary - r.from) * r.retention / 100
	}
}

func CreateRetention(from float32, to float32, retention float32) RetentionRange {
	return RetentionRange{
		from:      from,
		to:        to,
		retention: retention,
	}
}
