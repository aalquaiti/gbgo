package main

import (
	"flag"
	"fmt"
	"github.com/aalquaiti/gbgo/cpu"
	"github.com/aalquaiti/gbgo/io"
)

func main() {
	flag.Parse()
	cart := io.Cartridge{}
	//file := "./roms/delme.gb"
	file := "./roms/tetris.gb"
	// file := "./roms/DMG_ROM.bin"
	if err := cart.OpenRom(file); err != nil {
		panic(err)
	}
	bus := io.Bus{Rom: &cart}
	cpu.Init(cpu.DMG_MODE, bus)

	fmt.Println(cart.Header)

	//for i := 0; i < 100000; i++ {
	//	cpu.Step()
	//	//if bus.Read(0xFF02) == 0x81 {
	//	//	log.Fatal("Found it")
	//	//}
	//}

	for {
		cpu.Step()
	}

	//for {
	//	cpu.Step()
	//	if bus.Read(0xFF02) == 0x81 {
	//		log.Fatal("Found it")
	//	}
	//}

}
