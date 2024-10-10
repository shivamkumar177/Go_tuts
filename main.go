package main

import "fmt"

type fuelEngine struct {
	avg  uint8
	tank uint8
}

type electricEngine struct {
	avg    uint8
	charge uint8
}

func (e fuelEngine) distanceLeft() uint8 {
	return e.avg * e.tank
}

func (e electricEngine) distanceLeft() uint8 {
	return e.avg * e.charge
}

type engine interface {
	distanceLeft() uint8
}

func canMakeIt(e engine, distance uint8) {
	if distance <= e.distanceLeft() {
		fmt.Println("GO AHEAD")
	} else {
		fmt.Println(("NOT POSSIBLE"))
	}
}

func main() {
	var myEngine fuelEngine = fuelEngine{12, 1}
	canMakeIt(myEngine, 23)
	var myFriendEngine electricEngine = electricEngine{34, 2}
	canMakeIt(myFriendEngine, 23)
}
