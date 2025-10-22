package main

import (
	"demo/bin/bins"
	"demo/bin/storage"
)

func main() {
	bin, err := bins.NewBin("testBin", "qq-ww-ee", false)

	if err != nil {
		return
	}

	storage.SaveStorage(bin)
}
