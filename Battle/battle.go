package battle

import (
	"fmt"
	"math"
	"math/rand"
)

func (p *Pokemon) AttackPokemon(p2 *Pokemon) {
	fmt.Printf("%s attacks %s!\n", p.name, p2.name)
	damage := p.stats.attack - p2.stats.defense
	if damage < 0 {
		damage = 0
	}
	p2.stats.health -= damage
	fmt.Printf("%s takes %d damage!\n", p2.name, damage)
	battleEnd(&p, &p2)
}

func battleEnd(p *Pokemon, p2 *Pokemon) {
	if p2.stats.health <= 0 {
		fmt.Printf("%s has fainted!\n", p2.name)
		p.stats.exp += 100
		exit(0)
	}
	else{
		fmt.Printf("%s has fainted!\n", p.name)
		p2.stats.exp += 100
		exit(0)
	}
}
