package Pokemon

import (
	"fmt"
	"math"
	"math/rand"
)

type Pokemon struct {
	Name   string
	Stats Stats
}

type Stats struct {
	Attack  int
	Defense int
	SpecAtk int
	SpecDef int
	Speed   int
	Health int
	Level int
	Exp int
}

func (p *Pokemon) LevelUp() uint {
	p.stats.level++
	return uint(p.stats.level)
}