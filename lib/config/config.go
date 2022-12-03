package config

import (
	"errors"
	"gopkg.in/yaml.v3"
	_ "gopkg.in/yaml.v3"
	"io/ioutil"
	"reflect"
	"strings"
)

func ReadModuleConfig(dest interface{}, source string) error {

	if dest == nil {
		return errors.New("destination cannot be nil")
	}

	if !strings.HasSuffix(source, ".yaml") {
		return errors.New("file ext not supported yet")
	}

	if reflect.ValueOf(dest).Kind() != reflect.Pointer {
		return errors.New("destination should be pointer")
	}

	yamlFile, err := ioutil.ReadFile(source)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(yamlFile, dest)

	return nil
}
