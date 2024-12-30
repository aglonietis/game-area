package visualisation

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
	"image"
	"image/color"
	"soul/objects"
)

type Game struct {
	world        *objects.World
	screenWidth  int
	screenHeight int
	gridSize     int
}

func NewGame(world *objects.World) *Game {
	game := &Game{
		world:        world,
		screenWidth:  800,
		screenHeight: 800,
	}

	game.gridSize = game.screenWidth / (world.Radius * 2)

	ebiten.SetWindowSize(game.screenWidth, game.screenHeight)
	ebiten.SetWindowTitle("2D World Map")
	ebiten.SetTPS(10)
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}

	return game
}

func (g *Game) Update() error {
	for _, soul := range g.world.Souls {
		soul.Act()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Draw background
	screen.Fill(color.RGBA{0, 0, 0, 255})

	// Draw a grid
	for x := 0; x < g.screenWidth; x += g.gridSize {
		ebitenutil.DrawLine(screen, float64(x), 0, float64(x), float64(g.screenHeight), color.RGBA{255, 255, 255, 255})
	}
	for y := 0; y < g.screenHeight; y += g.gridSize {
		ebitenutil.DrawLine(screen, 0, float64(y), float64(g.screenWidth), float64(y), color.RGBA{255, 255, 255, 255})
	}

	for _, soul := range g.world.Souls {
		textImage := ebiten.NewImage(
			g.calculateTextWidth(soul.Initials, basicfont.Face7x13),
			basicfont.Face7x13.Metrics().Height.Ceil(),
		) // Adjust the size of this image based on your text length
		textImage.Fill(color.Transparent)

		// Render the text onto the offscreen image
		d := &font.Drawer{
			Dst:  textImage,
			Src:  image.NewUniform(color.White),
			Face: basicfont.Face7x13,
			Dot:  fixed.P(0, 0+basicfont.Face7x13.Metrics().Height.Ceil()), // , // Text position inside the offscreen image
		}
		d.DrawString(soul.Initials)

		// Scale the text image
		options := &ebiten.DrawImageOptions{}

		options.GeoM.Scale(
			float64(g.gridSize/g.calculateTextWidth(soul.Initials, basicfont.Face7x13)),
			float64(g.gridSize/basicfont.Face7x13.Metrics().Height.Ceil()),
		) // Scale text by 3x in both X and Y directions
		options.GeoM.Translate(float64(soul.Location.X*g.gridSize), float64(soul.Location.Y*g.gridSize)) // Position the scaled text on the screen

		// Draw the scaled text onto the screen
		screen.DrawImage(textImage, options)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.screenWidth, g.screenHeight
}

func (g *Game) calculateTextWidth(text string, face font.Face) int {
	width := 0
	for _, r := range text {
		advance, ok := face.GlyphAdvance(r)
		if ok {
			width += int(advance.Round())
		}
	}
	return width
}
