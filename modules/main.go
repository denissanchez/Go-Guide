package main

import (
	utils "github.com/denissanchez/mygolib"
	"github.com/donvito/hellomod"
	hellomod2 "github.com/donvito/hellomod/v2"
)

func main() {
	hellomod.SayHello()
	hellomod2.SayHello("Platzi")
	utils.Hello("Denis")
}
