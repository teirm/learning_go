// implementations of commas program
package main

import (
	"fmt"
	"github.com/teirm/learning_go/ch_3/commas"
)

func main() {
	fmt.Println(commas.Comma1("12345"))
	fmt.Println(commas.Comma2("1234568910"))
	fmt.Println(commas.Comma3("-123456.8910"))
}
