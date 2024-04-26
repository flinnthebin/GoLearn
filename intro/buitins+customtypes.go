package main

import "fmt"

var (
	floatVar32 float32 = 0.1
	floatVar64 float64 = 0.1
	name       string  = "Foo"
	intVar32   int32   = 1
	intVar64   int64   = 484848
	intVar     int     = -1
	uintVar    uint    = 1
	uintVar32  uint32  = 1
	uintVar64  uint64  = 10
	uintVar8   uint8   = 0x1
	byteVar    byte    = 0x2
	runeVar    rune    = 'a'
)

type Player struct {
	name        string
	health      int
	attackPower float64
}

func (player Player) getHealthAttached() int {
	return player.health
}

func getHealth(player Player) int {
	return player.health
}

type version int

type Weapon string

func getWeapon(weapon Weapon) string {
	return string(weapon)
}

func main() {
	player := Player{
		name:        "Captain Jack",
		health:      100,
		attackPower: 45.1,
	}

	fmt.Printf("This is the player health: %d\n", getHealth(player))
	fmt.Printf("This is the player health: %d\n", player.getHealthAttached())

	users := map[string]int{} // empty map
	users["foo"] = 10         // add to map
	users["bar"] = 11

	//users2 := map[string]int{ // value initialized
	//    "chris" : 33,
	//}

	//users3 := make(map[string]int) // less idiomatic golang empty map init

	age := users["10"]
	fmt.Printf("age: %d\n", age)
	users["baz"] = 12
	ages, ok := users["baz"]
	if !ok {
		fmt.Println("baz does not exist")
	} else {
		fmt.Println("baz does exist")
	}
	fmt.Printf("ages: %d\n", ages)

	delete(users, "foo") // delete from map

	for k, v := range users {
		fmt.Printf("The key %s and the value %d\n", k, v)
	}

	numbers := []int{} // create a slice
	otherNumbers := make([]int, 2)
	numArr := [2]int{} // create an array

	fmt.Println(numbers)
	fmt.Println(otherNumbers)
	fmt.Println(numArr)

	numbers = append(numbers, 1)
	numbers = append(numbers, 2)
	numbers = append(numbers, 3)

	fmt.Println(numbers)
}
