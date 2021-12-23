package main

import (
	"flag"
	"github.com/aalquaiti/gbgo/cartridge"
	"github.com/aalquaiti/gbgo/cpu"
	"github.com/aalquaiti/gbgo/io"
)

func main() {
	flag.Parse()
	file := "./roms/delme.gb"
	//file := "./roms/tetris.gb"
	// file := "./roms/DMG_ROM.bin"
	cart, err := cartridge.NewCartridge(file)
	if err != nil {
		panic(err)
	}
	cpu.Init(cpu.DMG_MODE, io.NewBus(cart))

	//fmt.Println(cart.Header)

	for i := 0; i < 100000; i++ {
		cpu.Step()
		//if io.Read(0xFF02) == 0x81 {
		//	log.Fatal("Found it")
		//}
	}

	//for {
	//	cpu.Step()
	//}

	//for {
	//	cpu.Step()
	//	if io.Read(0xFF02) == 0x81 {
	//		log.Fatal("Found it")
	//	}
	//}

}
