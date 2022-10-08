package config

type Yaml struct {
	Server struct {
		Listen int       `yaml:"listen"`
		Path   string    `yaml:"path"`
		Tls    TlsConfig `yaml:"tls"`
	} `yaml:"server"`
	Databases DatabasesConfig `yaml:"databases"`
}
type DatabasesConfig struct {
	Host          string `yaml:"host"`
	User          string `yaml:"user"`
	Passwd        string `yaml:"passwd"`
	DatabasesName string `yaml:"databasesName"`
}
type TlsConfig struct {
	Enable   bool   `yaml:"enable"`
	CertFile string `yaml:"certFile"`
	KeyFile  string `yaml:"keyFile"`
}
