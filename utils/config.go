package utils

import (
	"encoding/json"
	"os"
)

func ReadConfig(parseTo interface{}, names ...string) error {
	var name string
	if len(names) == 0 {
		name = "config.json"
	} else {
		name = names[0]
	}

	d, err := os.ReadFile(name)

	if err != nil {
		return err
	}

	return json.Unmarshal(d, parseTo)
}
