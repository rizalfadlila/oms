package config

type ServerConfig struct {
	Rest RestConfig `yaml:"rest"`
}

type RestConfig struct {
	Host            string `yaml:"host"`
	Port            int    `yaml:"port"`
	GracefulTimeout string `yaml:"gracefulTimeout"`
	WriteTimeout    string `yaml:"writeTimeout"`
	ReadTimeout     string `yaml:"readTimeout"`
	RouterTimeout   string `yaml:"routerTimeout"`
	EnableSwagger   bool   `yaml:"enableSwagger"`
}
