package config

type Main struct {
	Server   ServerConfig `yaml:"server"`
	Database Database     `yaml:"database"`
}
