package repository

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Bucket struct {
	UpTo      float32 `yaml:"upTo"`
	Retention float32 `yaml:"retention"`
}

func NewRegionRepository() RegionRepository {
	filename, _ := filepath.Abs("./config/region_retentions.yaml")
	yamlFile, _ := ioutil.ReadFile(filename)

	retentions := make(map[string][]Bucket)
	err := yaml.Unmarshal(yamlFile, &retentions)

	if err != nil {
		fmt.Println(err)
	}

	return RegionRepository{regionRetentions: retentions}
}

func NewStateRepository() StateRepository {
	filename, _ := filepath.Abs("./config/state_retentions.yaml")
	yamlFile, _ := ioutil.ReadFile(filename)

	retentions := make(map[string][]Bucket)
	err := yaml.Unmarshal(yamlFile, &retentions)

	if err != nil {
		fmt.Println(err)
	}

	return StateRepository{
		retentions: retentions["spain"],
	}
}
