package storage

import (
	"demo/bin/bins"
	"demo/bin/file"
	"encoding/json"
)

func SaveStorage(bin *[]bins.Bin) error {
	data, err := json.Marshal(bin)
		if err != nil {
		return err
	}
	file.WriteFile(data, "data.json")
	return nil
}

func ReadStorage() ([]bins.Bin, error) {
	file, err := file.ReadFile("data.json")
	if err != nil {
		return nil, err
	}
	var bins []bins.Bin
	err = json.Unmarshal(file, &bins)
	if err != nil {
		return nil, err
	}
	return bins, nil
}