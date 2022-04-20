package main

import (
	"fmt"

	"github.com/lazybark/go-pretty-code/console"
)

func main() {
	fmt.Println("This is simple prettifier for Go")
	fmt.Println("===Examples===")
	fmt.Println(console.TextColorRed + "TextColorRed" + console.TextColorReset)
	fmt.Println(console.TextColorGreen + "TextColorGreen" + console.TextColorReset)
	fmt.Println(console.TextColorYellow + "TextColorYellow" + console.TextColorReset)
	fmt.Println(console.TextColorBlue + "TextColorBlue" + console.TextColorReset)
	fmt.Println(console.TextColorPurple + "TextColorPurple" + console.TextColorReset)
	fmt.Println(console.TextColorCyan + "TextColorCyan" + console.TextColorReset)
	fmt.Println(console.TextColorGray + "TextColorGray" + console.TextColorReset)
	fmt.Println(console.TextColorWhite + "TextColorWhite (ha-ha)" + console.TextColorReset)

}
