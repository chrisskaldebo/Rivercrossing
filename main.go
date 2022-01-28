package main

import (
	"fmt"
	"github.com/chrisskaldebo/rivercrossing/state"
)

func main() {
	fmt.Println(state.CrossRiver())

	fmt.Println(state.PutInBoat())

	fmt.Println(state.ViewState())

	fmt.Println(state.PutInBoat2())

	fmt.Println(state.ViewState2())
}
