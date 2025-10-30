package bins

import (
	"errors"
	"time"
)

type Bin struct {
	Id string `json:"id"`
	Private bool `json:"private"`
	CreatedAt time.Time `json:"createdAt"`
	Name string `json:"name"`
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
		Id: id,
		Private: private,
		Name: name,
		CreatedAt: time.Now(),
	}, nil
}
