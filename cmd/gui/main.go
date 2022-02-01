package main

import (
	"fmt"
	"github.com/aalquaiti/gbgo/cpu"
	"github.com/aalquaiti/gbgo/ppu"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/sirupsen/logrus"
	"log"
	"os"
)

const file = "./roms/blargg/cpu_instrs/individual/01-special.gb"

// gui Represents ebiten game
type gui struct {
	ppu ppu.PPU
}

func (g *gui) Update() error {
	for i := 0; i < cpu.DMG_HZ/60; i++ {

	}

	return nil
}

func (g *gui) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %.2f", ebiten.CurrentTPS()))
	//ebitenutil.DebugPrint(screen, g.str)
	//ebitenutil.DebugPrintAt(screen, g.str, 0, 20)
}

func (g *gui) Layout(width, height int) (int, int) {
	return 160, 144
}

func main() {
	f, err := os.OpenFile("debug.log", os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		log.Panicf("failed to open debug.log %v", err)
	}
	defer f.Close()
	ebiten.SetMaxTPS(60)

	gui := &gui{}
	gui.ppu = ppu.NewPPU()

	//cart, err := cartridge.NewCartridge(file)
	if err != nil {
		panic(err)
	}
	//cpu.Init(cpu.DMG_MODE, io.NewBus(cart, gui.ppu))
	//log.WithField("Cart Header", cart.Header).Info()

	logrus.SetOutput(f)
	logrus.SetLevel(logrus.DebugLevel)

	ebiten.SetWindowSize(640, 480)
	if err := ebiten.RunGame(gui); err != nil {
		logrus.Fatal(err)
	}
}
