package main

import (
	"soul/visualisation"
)
import "soul/objects"

func main() {
	world := objects.NewWorld(10)
	world.AddSoul(objects.NewSoul("Alice", world))
	world.AddSoul(objects.NewSoul("Ben", world))
	world.AddSoul(objects.NewSoul("Charlie", world))
	world.AddSoul(objects.NewSoul("Dave", world))
	world.AddSoul(objects.NewSoul("Emerald", world))
	world.AddSoul(objects.NewSoul("Falcon", world))
	world.AddSoul(objects.NewSoul("Gladiator", world))
	world.AddSoul(objects.NewSoul("Hector", world))
	world.AddSoul(objects.NewSoul("Ivan", world))
	world.AddSoul(objects.NewSoul("Jayce", world))
	world.AddSoul(objects.NewSoul("Karl", world))

	visualisation.NewGame(world)
}
