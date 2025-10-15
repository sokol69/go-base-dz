package main

import (
	"errors"
	"fmt"
	"time"
)

type Bin struct {
	id string
	private bool
	createdAt time.Time
	name string
}

type BinList = []Bin

func newBin(name, id string, private bool) (*Bin, error) {
	if name == "" {
		return nil, errors.New("INVALID_NAME")
	}

	if id == "" {
		return nil, errors.New("INVALID_ID")
	}

	return &Bin{
		id: id,
		private: private,
		name: name,
		createdAt: time.Now(),
	}, nil
}

func main() {
	bin, err := newBin("testBin", "qq-ww-ee", false)

	if err != nil {
		return
	}

	fmt.Println(bin)
}
