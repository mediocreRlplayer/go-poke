package trainer

import (
	"fmt"
	Pokemon "github.com/mediocreRlplayer/go-poke"
)

type Trainer struct {
	Name string
	Pokemon [6]Pokemon
	Bag Bag
	Badges [8]string
}

func (t *Trainer) SelectName(string name) string{
	t.Name = name
	return t.Name
}