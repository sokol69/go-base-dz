package main

import (
	"demo/bin/bins"
	"demo/bin/storage"
	"fmt"
)

func main() {
	bin, err := bins.NewBin("testBin", "qq-ww-ee", false)

	if err != nil {
		fmt.Println(err)
		return
	}
	bins := []bins.Bin{}
	bins = append(bins, *bin)

	storage.SaveStorage(&bins)
	binsFromFile, err := storage.ReadStorage()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binsFromFile)
}
