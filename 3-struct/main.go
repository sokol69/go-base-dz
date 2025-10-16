package main

import (
	"demo/bin/bins"
	"fmt"
)

func main() {
	bin, err := bins.NewBin("testBin", "qq-ww-ee", false)

	if err != nil {
		return
	}

	fmt.Println(bin)
}
