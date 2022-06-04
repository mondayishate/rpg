package main

import (
	"awesomeProject/scenario"
	"awesomeProject/unit"
	"fmt"
)

var ToolSetting []unit.Tool

func gameSetting() {
	ToolSetting = []unit.Tool{
		{Id: unit.CURE, Name: "cure"},
		{Id: unit.WEAPON, Name: "weapon"},
		{Id: unit.EquippedWeapon, Name: "equipped weapon"},
	}
}

func starterItem() []unit.Item {
	return []unit.Item{
		{
			Tool: ToolSetting[unit.CURE], Name: "잼민이의 알약", Value1: 13, Count: 8,
		},
	}
}

// todo : 이제 할일 - 회복아이템
func main() {
	gameSetting()

	var name string
	fmt.Print("당신의 이름은?")
	fmt.Scan(&name)

	hero := unit.Character{
		Name: name, HP: 50, Damage: 16,
		Items: starterItem(),
	}

	monster := unit.Character{
		Name: "뒤틀린 잼민이", HP: 15, Damage: 6,
	}

	scenario.Battle(&hero, monster)

	monster = unit.Character{
		Name: "흑화한 잼민이", HP: 25, Damage: 6,
	}

	scenario.Battle(&hero, monster)
}
