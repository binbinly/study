package conf

// PrometheusConfig
type PrometheusConfig struct {
	Enable bool
	Host   string
}

// SentryConfig
type SentryConfig struct {
	Enable bool
	Dsn    string
}
