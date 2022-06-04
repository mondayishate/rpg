package unit

import (
	"fmt"
	"math/rand"
)

type Character struct {
	Name   string
	HP     int
	Damage int
	Items  []Item
}

func (r Character) Attack(r2 *Character) {
	d := rand.Intn(r.Damage)
	r2.HP -= d
	fmt.Printf("%s(은)는 %s에게 %d 피해를 입혔다\n", r.Name, r2.Name, d)
}

func Status(r1, r2 Character) {
	fmt.Printf("%s 체력 : %d, %s 체력 : %d\n", r1.Name, r1.HP, r2.Name, r2.HP)
}
