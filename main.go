package main

import (
	tl "github.com/JoelOtter/termloop"
	// "time"
)

func main() {
	g := tl.NewGame()
	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlack,
	})

	// set the character at position (0, 0) on the entity.
	player := NewPlayer(15, 15, 1, 1)
	level.AddEntity(player)
	level.AddEntity(player.Weapon)

	g.Screen().SetFps(30)
	g.Screen().SetLevel(level)
	g.Start()
}

type Player struct {
	*tl.Entity
	prevX  int
	prevY  int
	level  *tl.BaseLevel
	Weapon Weapon
}
type Weapon struct {
	*tl.Entity
}

func NewPlayer(x, y, height, width int) *Player {
	weapon := Weapon{
		Entity: tl.NewEntity(x, y-1, height, width),
	}

	p := Player{
		Entity: tl.NewEntity(x, y, height, width),
		Weapon: weapon,
	}

	p.SetCell(0, 0, &tl.Cell{Fg: tl.ColorRed, Ch: '옷'})
	p.Weapon.SetCell(0, 0, &tl.Cell{Fg: tl.ColorRed, Ch: 'I'})

	return &p
}

type Invader struct {
	*tl.Entity
	x     int
	y     int
	level *tl.BaseLevel
}

// func (i *Invader) NewInvader() *Invader {
//   return &Invader{
//     Entity: tl.NewEntity(10, 10, 1, 1),
//   }
// }

func (p *Player) ShootWeapon() {
	pX, pY := p.Position()
	p.Weapon.SetPosition(pX, pY-1)

	for i := pY - 1; i > 0; i-- {
		p.Weapon.SetPosition(pX, i)
	}
}

// func (w *Weapon) Collide(collision tl.Physical) {
//   if _, ok := p.(*Invader);
//
// }

// Physical represents something that can collide with another
// Physical, but cannot process its own collisions.
// Optional addition to Drawable.
type Physical interface {
	Position() (int, int) // Return position, x and y
	Size() (int, int)     // Return width and height
}

// DynamicPhysical represents something that can process its own collisions.
// Implementing this is an optional addition to Drawable.
type DynamicPhysical interface {
	Position() (int, int) // Return position, x and y
	Size() (int, int)     // Return width and height
	Collide(Physical)     // Handle collisions with another Physical
}

func (player *Player) Tick(event tl.Event) {
	if event.Type == tl.EventKey { // Is it a keyboard event?
		x, y := player.Position()
		switch event.Key { // If so, switch on the pressed key.
		case tl.KeyArrowRight:
			player.SetPosition(x+1, y)
		case tl.KeyArrowLeft:
			player.SetPosition(x-1, y)
		case tl.KeyArrowUp:
			player.SetPosition(x, y-1)
		case tl.KeyArrowDown:
			player.SetPosition(x, y+1)
		case tl.KeySpace:
			player.ShootWeapon()
		}
	}
}

func (player *Player) Collide(collision tl.Physical) {
	// Check if it's a Rectangle we're colliding with
	if _, ok := collision.(*tl.Rectangle); ok {
		player.SetPosition(player.prevX, player.prevY)
	}
}

// func (player *Player) Draw(screen *tl.Screen) {
// 	screenWidth, screenHeight := screen.Size()
// 	x, y := player.Position()
// 	player.level.SetOffset(screenWidth/2-x, screenHeight/2-y)
//   // We need to make sure and call Draw on the underlying Entity.
// 	player.Entity.Draw(screen)
// }
