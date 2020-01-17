package amanar

import "gopkg.in/yaml.v2"

func UnmarshalYamlAmanarConfiguration(data []byte) (AmanarConfiguration, error) {
	var r AmanarConfiguration
	err := yaml.Unmarshal(data, &r)
	return r, err
}

func (r *AmanarConfiguration) MarshalYaml() ([]byte, error) {
	return yaml.Marshal(r)
}