package main

import "fmt"

type WeaponType int

func (w WeaponType) String() string {
    switch w {
    case Sword:
        return "SWORD"
    case Axe:
        return "AXE"
    default:
        return ""
    }
}

const (
    Axe         WeaponType = iota // construct enum list
    Sword
    WoodenStick
    Knife
)

func getDamage(weaponType WeaponType) int {
    switch weaponType {
    case Axe:
        return 100
    case Sword:
        return 90
    case WoodenStick:
        return 1
    case Knife:
        return 40
    default:
        panic("weapon does not exist")
    }
}

func main() {
    fmt.Printf("damage of the given weapon (%d) (%d):\n", WoodenStick, getDamage(WoodenStick))
    fmt.Printf("damage of the given weapon (%s) (%d):\n", Sword, getDamage(Sword))
}
