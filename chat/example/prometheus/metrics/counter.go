package metrics

import "github.com/prometheus/client_golang/prometheus"

type CounterVecOpts VectorOpts

type CounterVec interface {
	Inc(Labels ...string)
	Add(v float64, labels ...string)
}

type promCounterVec struct {
	counter *prometheus.CounterVec
}

func NewCounterVec(cfg *CounterVecOpts) CounterVec {
	if cfg == nil {
		return nil
	}

	vec := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: cfg.Namespace,
			Subsystem: cfg.Subsystem,
			Name: cfg.Name,
			Help: cfg.Help,
		}, cfg.Labels)
	prometheus.MustRegister(vec)
	return &promCounterVec{
		counter: vec,
	}
}

func (p promCounterVec) Inc(Labels ...string) {
	p.counter.WithLabelValues(Labels...).Inc()
}

func (p promCounterVec) Add(v float64, labels ...string) {
	p.counter.WithLabelValues(labels...).Add(v)
}