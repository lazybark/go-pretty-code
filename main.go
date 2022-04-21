package main

import (
	"fmt"

	"github.com/lazybark/go-pretty-code/console"
)

func main() {
	fmt.Println("This is simple prettifier for Go")
	fmt.Println("===Examples===")
	fmt.Println(console.FormatRed("TextColorRed"))
	fmt.Println(console.FormatGreen("TextColorGreen"))
	fmt.Println(console.FormatYellow("TextColorYellow"))
	fmt.Println(console.FormatBlue("TextColorBlue"))
	fmt.Println(console.FormatPurple("TextColorPurple"))
	fmt.Println(console.FormatCyan("TextColorCyan"))
	fmt.Println(console.FormatGray("TextColorGray"))
	fmt.Println(console.FormatWhite("TextColorWhite (ha-ha)"))
	fmt.Println("\n" + console.TextColorRed + "Some red text can be splitted by " + console.ResetColor() + "console.ResetColor() and become default." + console.TextColorGreen + " Then it can change color again." + console.ResetColor() + console.TextColorCyan + " You can even change color of console path!" + console.TextColorPurple)
}
