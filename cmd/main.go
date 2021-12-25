package main

import (
	"fmt"
	"github.com/aalquaiti/gbgo/cartridge"
	"github.com/aalquaiti/gbgo/cpu"
	"github.com/aalquaiti/gbgo/io"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	log "github.com/sirupsen/logrus"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

const file = "./roms/delme.gb"

// gui Represents ebiten game
type gui struct {
	str string
}

func (g *gui) Update() error {
	for i := 0; i < cpu.DMG_HZ/60; i++ {
		g.str = cpu.Tick()
	}

	for g.str == "" {
		g.str = cpu.Tick()
	}

	return nil
}

func (g *gui) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %.2f", ebiten.CurrentFPS()))
	//ebitenutil.DebugPrint(screen, g.str)
	ebitenutil.DebugPrintAt(screen, g.str, 0, 20)
}

func (g *gui) Layout(width, height int) (int, int) {
	return 160, 144
}

func init() {
	log.SetOutput(os.Stdout)
	cart, err := cartridge.NewCartridge(file)
	if err != nil {
		panic(err)
	}
	cpu.Init(cpu.DMG_MODE, io.NewBus(cart))
	log.WithField("Cart Header", cart.Header).Info()
}

func main() {

	ebiten.SetMaxTPS(60)
	ebiten.SetWindowSize(640, 480)
	if err := ebiten.RunGame(&gui{}); err != nil {
		log.Fatal(err)
	}
}
