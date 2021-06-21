package metrics

import "github.com/prometheus/client_golang/prometheus"

type GaugeVecOpts VectorOpts

type GaugeVec interface {
	Set(v float64, labels ...string)

	Inc(labels ...string)

	Add(v float64, labels ...string)
}

type promGaugeVec struct {
	gauge *prometheus.GaugeVec
}

func NewGaugeVec(cfg *GaugeVecOpts) GaugeVec {
	if cfg == nil {
		return nil
	}
	vec := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: cfg.Namespace,
			Subsystem: cfg.Subsystem,
			Name: cfg.Name,
			Help: cfg.Help,
		}, cfg.Labels)
	prometheus.MustRegister(vec)
	return &promGaugeVec{
		gauge: vec,
	}

}

func (p promGaugeVec) Set(v float64, labels ...string) {
	p.gauge.WithLabelValues(labels...).Set(v)
}

func (p promGaugeVec) Inc(labels ...string) {
	p.gauge.WithLabelValues(labels...).Inc()
}

func (p promGaugeVec) Add(v float64, labels ...string) {
	p.gauge.WithLabelValues(labels...).Add(v)
}