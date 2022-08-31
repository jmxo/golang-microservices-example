package main

type config struct {
	API        apiConfig        `yaml:"api"`
	Prometheus prometheusConfig `yaml:"prometheus"`
}

type apiConfig struct {
	Port string `yaml:"port"`
}

type prometheusConfig struct {
	MetricsPort int `yaml:"metricsPort"`
}
