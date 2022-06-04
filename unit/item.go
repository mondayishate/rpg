package unit

import (
	"errors"
	"fmt"
)

type Item struct {
	Tool
	Name   string
	Count  int
	Value1 int
}

const (
	CURE           = iota
	WEAPON         = iota
	EquippedWeapon = iota
)

type Tool struct {
	Id   int
	Name string
}

func (i *Item) UseItem(c *Character) (*Item, error) {
	if i.Count == 0 {
		return &Item{}, errors.New("used item")
	}
	switch i.Tool.Id {
	case CURE:
		c.HP += i.Value1
	default:
		fmt.Print()
	}
	i.Count--
	return i, nil
}
