package config

import "github.com/jatis/oms/lib/database/sql"

type Database struct {
	MariaDB sql.DBConfig `yaml:"mariadb"`
}
