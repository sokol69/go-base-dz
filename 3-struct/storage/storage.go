package storage

import (
	"demo/bin/bins"
	"demo/bin/file"
	"encoding/json"
	"fmt"
)

func SaveStorage(bin *bins.Bin) error {
	data, err := json.Marshal(bin)
		if err != nil {
		return err
	}
	file.WriteFile(data, "data.json")
	return nil
}

func ReadStorage() error {
	file, err := file.ReadFile("data.json")
	if err != nil {
		return err
	}
	var bin bins.Bin
	err = json.Unmarshal(file, &bin)
	if err != nil {
		return err
	}
	fmt.Println(bin)
	return nil
}