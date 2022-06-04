package scenario

import (
	"awesomeProject/unit"
	"fmt"
	"time"
)

const (
	_       = iota
	BATTLE  = iota
	UseItem = iota
)

func Battle(hero *unit.Character, monster unit.Character) {
	var cmd int
	fmt.Printf("%s 앞에 %s가 나타났다. \n", hero.Name, monster.Name)
	for true {
		fmt.Printf("공격 : %d / 회복아이템 사용 : %d\n", BATTLE, UseItem)
		_, err := fmt.Scan(&cmd)
		if err != nil {
			fmt.Printf("battle err : %v", err)
			return
		}
		if cmd == BATTLE {
			fmt.Printf("%s(이)가 %s를 공격했다.\n", hero.Name, monster.Name)
			hero.Attack(&monster)
			time.Sleep(time.Second)

			fmt.Printf("%s가 %s(을)를 공격했다.\n", monster.Name, hero.Name)
			monster.Attack(hero)
			unit.Status(*hero, monster)
			time.Sleep(time.Second)

			fmt.Println()
			fmt.Println()

		} else if cmd == UseItem {
			item, err := hero.Items[0].UseItem(hero)
			if err != nil {
				fmt.Println("이미 아이템을 모두 써버렸습니다.(잼민이 수준)")
				continue
			}
			fmt.Printf("%d개의 %s(이)가 남았습니다\n", item.Count, item.Name)
			unit.Status(*hero, monster)
		}
		if hero.HP < 1 {
			fmt.Printf("%s가 쓰러졌다\n게임 패배\n", hero.Name)
			break
		}
		if monster.HP < 1 {
			fmt.Printf("%s를 쓰러트렸다\n게임 승리\n", monster.Name)
			break
		}
	}
}
