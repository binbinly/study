package metrics

import "github.com/prometheus/client_golang/prometheus"

type HistogramVecOpts struct {
	Namespace string
	SubSystem string
	Name string
	Help string
	Labels []string
	Buckets []float64
}

type HistogramVec interface {
	Observe(v int64, labels ...string)
}

type promHistogramVec struct {
	histogram *prometheus.HistogramVec
}

func NewHistogramVec(cfg *HistogramVecOpts) HistogramVec {
	if cfg == nil {
		return nil
	}
	vec := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: cfg.Namespace,
			Subsystem: cfg.SubSystem,
			Name: cfg.Name,
			Help: cfg.Help,
			Buckets: cfg.Buckets,
		}, cfg.Labels)
	prometheus.MustRegister(vec)
	return &promHistogramVec{
		histogram: vec,
	}
}

func (p promHistogramVec) Observe(v int64, labels ...string) {
	p.histogram.WithLabelValues(labels...).Observe(float64(v))
}