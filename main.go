package main

import (
	"fmt"

	"github.com/lazybark/go-pretty-code/console"
	"github.com/lazybark/go-pretty-code/logs"
	"go.uber.org/zap"
)

func main() {
	fmt.Println("This is simple prettifier for Go")
	fmt.Println("===Examples===")
	fmt.Println(console.ForeRed("ForeRed"))
	fmt.Println(console.ForeGreen("ForeGreen"))
	fmt.Println(console.ForeYellow("ForeYellow"))
	fmt.Println(console.ForeBlue("ForeBlue"))
	fmt.Println(console.ForeMagenta("ForeMagenta"))
	fmt.Println(console.ForeCyan("ForeCyan"))
	fmt.Println(console.ForeGray("ForeGray"))
	fmt.Println(console.ForeWhite("ForeWhite (ha-ha)"))
	fmt.Println()
	fmt.Println(console.BackRed("BackRed"))
	fmt.Println(console.BackGreen("BackGreen"))
	fmt.Println(console.BackYellow("BackYellow"))
	fmt.Println(console.BackBlue("BackBlue"))
	fmt.Println(console.BackMagenta("BackMagenta"))
	fmt.Println(console.BackBlack("BackBlack"))
	fmt.Println(console.BackCyan("BackCyan"))
	fmt.Println(console.BackWhite("BackWhite"))
	fmt.Println()
	fmt.Println(console.TextBold("TextBold"))
	fmt.Println(console.TextUnderline("TextUnderline"))
	fmt.Println(console.TextBlink("TextBlink"))
	fmt.Println()
	fmt.Println(console.ForeBlue(console.BackWhite("All parameters can be combined like: console.ForeBlue(console.BackWhite())")))
	fmt.Println(console.ForeGreen(console.TextBlink("Bold")))
	fmt.Println()
	fmt.Println(console.ForeBlue(console.BackWhite("Text can " + console.TextReverse("be reversed") + console.TextUnReverse(" and unreversed back, ") + "but in both cases needs to have" + console.ResetText() + " ResetText() at the end.")))

	fmt.Println()
	fmt.Println("\n" + console.TextColorRed + "Colored text can be splitted by " + console.ResetText() + "console.ResetText() and become default." + console.TextColorGreen + " Then it can change color again." + console.ResetText() + console.TextColorCyan + " You can even change color of console path if you leave a console.Fore___ var at the end (without console.ResetText())." + console.ResetText())

	logger, _ := logs.Double("some\\log.txt", false, zap.InfoLevel)
	logger.Fatal("That's bad")
}
