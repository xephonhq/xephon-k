package config

type ClientConfig struct {
	Addr     string `yaml:"addr"`
	Prepare  bool   `yaml:"prepare"`
	Columnar bool   `yaml:"columnar"`
}
