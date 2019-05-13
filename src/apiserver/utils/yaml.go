package utils

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

func MarshalToYAML(data interface{}, targetPath string) error {
	output, err := yaml.Marshal(data)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(targetPath, output, 0644)
}
