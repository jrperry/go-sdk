package iland

type PerformanceCounter struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Group string `json:"group"`
}

type Performance struct {
	ID       string              `json:"uuid"`
	Name     string              `json:"name"`
	Type     string              `json:"type"`
	Group    string              `json:"group"`
	Summary  string              `json:"summary"`
	Interval int                 `json:"interval"`
	Unit     string              `json:"unit"`
	Samples  []PerformanceSample `json:"samples"`
}

type PerformanceSample struct {
	Value     int `json:"value"`
	Timestamp int `json:"time"`
}

func (p *Performance) GetMaxValue() int {
	max := 0
	for _, sample := range p.Samples {
		if sample.Value > max {
			max = sample.Value
		}
	}
	return max
}

func (p *Performance) GetAvgValue() int {
	total := 0
	for _, sample := range p.Samples {
		total += sample.Value
	}
	return int(total / len(p.Samples))
}
