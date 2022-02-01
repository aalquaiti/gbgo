package main

import (
	"fmt"
	"github.com/aalquaiti/gbgo/cartridge"
	"github.com/aalquaiti/gbgo/cpu"
)

const file = "./roms/blargg/cpu_instrs/individual/01-special.gb"

func main() {
	cart, err := cartridge.NewCartridge(file)
	if err != nil {
		panic(err)
	}
	dsm := cpu.NewDisassembler()

	result, err := dsm.DisassembleAll(cart)

	if err != nil {
		panic(err)
	}

	for _, line := range result {
		fmt.Println(line)
	}
}
