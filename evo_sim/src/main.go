package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const TPS int = 60
const dTimeTick float64 = 1.0 / float64(TPS)

const nSpecies = 8
const nInSpecies = 100
const nOrgs = nSpecies * nInSpecies

var organisms [nOrgs]Entity
var species [nSpecies]([nInSpecies]*Entity) // array nInSpecies in nSpecies, containing pointers to the corrosponding entities in entities

type Game struct {
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	// for idx := 0; idx < nOrgs; idx++ {
	// 	e := &organisms[idx]
	// 	angle := (math.Pi / 2.0 / 800.0 * float64(idx))
	// 	e.x += math.Sin(angle)
	// 	e.y += math.Cos(angle)
	// }
	return nil
}

// Write your game's rendering.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, fmt.Sprint(int(ebiten.ActualFPS()+0.1), " FPS"))
	for idx := 0; idx < nOrgs; idx++ {
		e := &organisms[idx]
		drawCircle(screen, int(e.x), int(e.y), int(e.r), color.White)
	}
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640 * 2, 480 * 2
}

func main() {
	game := &Game{}

	//Init the entities
	for s := 0; s < nSpecies; s++ {
		for org := 0; org < nInSpecies; org++ {
			e := Entity{x: float64(org * 30), y: float64(s * 30)}
			e.updateRadius(10)
			organisms[s*nInSpecies+org] = e
			species[s][org] = &e
		}
	}

	// Sepcify the window size as you like. Here, a doulbed size is specified.
	ebiten.SetWindowSize(640*2, 480*2)
	ebiten.SetWindowTitle("Evolution Simulator")
	ebiten.SetTPS(TPS)
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
