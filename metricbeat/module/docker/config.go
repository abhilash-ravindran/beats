package docker

type TlsConfig struct {
	Enabled  bool   `config:"enable"`
	CaPath   string `config:"ca_path"`
	CertPath string `config:"cert_path"`
	KeyPath  string `config:"key_path"`
}

type Config struct {
	Socket string `config:"socket"`
	Tls    TlsConfig
}

func GetDefaultConf() *Config {
	return &Config{
		Socket: "",
		Tls: TlsConfig{
			Enabled: false,
		},
	}
}
