package config

import (
	"encoding/json"
	"io/ioutil"
)

//Load parse json to object
func Load(fileName string, object interface{}) (interface{}, error) {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		return object, err
	}
	json.Unmarshal(file, &object)
	return object, nil
}
