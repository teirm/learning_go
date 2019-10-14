package main

import (
	"flag"
	"fmt"
	"github.com/teirm/learning_go/ch_7/tempconv"
)

var temp = tempconv.CelsiusFlag("temp", 20.0, "the temperature")
var kTemp = tempconv.KelvinFlag("ktemp", 30.0, "the temperature in kelvin")

func main() {
	flag.Parse()
	fmt.Println(*temp)
	fmt.Println(*kTemp)
}
