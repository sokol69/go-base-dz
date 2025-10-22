package file

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func ReadFile(name string) ([]byte, error) {
	isJson := strings.HasSuffix(name, ".json")
	if !isJson {
		return nil, errors.New("INVALID_FORMAT")
	}
	data, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func WriteFile(content []byte, name string) {
	file, err := os.Create(name)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	_, err = file.Write(content)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Success!")
}