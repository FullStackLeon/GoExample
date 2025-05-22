package Basic

import (
	"GoLib/Basic/player"
	"GoLib/Basic/weapon"
	"fmt"
	"testing"
)

type IPeople interface {
	Talk()
}

type Woman struct {
	Name string
}

func (d *Woman) Talk() {}

func NewInterfacePeople() IPeople {
	var women *Woman
	return women
}

func NewStructWoman() *Woman {
	var women *Woman
	return women
}

func TestInterface(t *testing.T) {
	p1 := NewInterfacePeople()
	if people, ok := p1.(*Woman); ok {
		if people != nil {
			people.Talk()
		} else {
			fmt.Println("p1 is nil")
		}
	} else {
		fmt.Println("p1 is not nil")
	}

	p2 := NewStructWoman()
	if p2 != nil {
		p2.Talk()
		fmt.Println("p2 is not nil")
	} else {
		fmt.Println("p2 is nil")
	}
}

func TestPlayer(t *testing.T) {
	pl := player.Player{
		Health: 100,
	}
	wpn := weapon.Weapon{
		Damage: 10,
	}

	wpn.Attack(&pl)
	fmt.Println(pl.Health)
}
