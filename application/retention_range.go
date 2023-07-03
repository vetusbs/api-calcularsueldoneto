package application

type retentionRange struct {
	from      float32
	to        float32
	retention float32
}

func (r retentionRange) retentionOverSalary(salary float32) float32 {
	if r.from >= salary {
		return 0
	} else if r.to <= salary {
		return (r.to - r.from) * r.retention / 100
	} else {
		return (salary - r.from) * r.retention / 100
	}
}
