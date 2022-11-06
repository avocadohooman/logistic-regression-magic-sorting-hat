package stats

type Stats struct {
	FeatureName string
	Count       int
	Mean        float64
	Std         float64
	Min         float64
	Max         float64
	TwentyFive  float64
	Fifty       float64
	SeventyFive float64
}

type StatsArray []Stats
