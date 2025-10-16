package bins

import (
	"errors"
	"time"
)

type Bin struct {
	id string
	private bool
	createdAt time.Time
	name string
}

type BinList = []Bin

func NewBin(name, id string, private bool) (*Bin, error) {
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
