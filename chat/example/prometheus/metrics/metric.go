package metrics

type Opts struct {

}

type Metric interface {
	Add(int64)
	Value() int64
}

type Aggregation interface {
	Min() float64
	Max() float64
	Avg() float64
	Sum() float64
}

type VectorOpts struct {
	Namespace string
	Subsystem string
	Name string
	Help string
	Labels []string
}