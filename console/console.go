package console

import "fmt"

var (
	TextReset = "\033[0m"

	TextColorRed     = "\033[31m"
	TextColorGreen   = "\033[32m"
	TextColorYellow  = "\033[33m"
	TextColorBlue    = "\033[34m"
	TextColorMagenta = "\033[35m"
	TextColorCyan    = "\033[36m"
	TextColorGray    = "\033[37m"
	TextColorWhite   = "\033[97m"

	BackColorBlack   = "\033[40m"
	BackColorRed     = "\033[41m"
	BackColorGreen   = "\033[42m"
	BackColorYellow  = "\033[43m"
	BackColorBlue    = "\033[44m"
	BackColorMagenta = "\033[45m"
	BackColorCyan    = "\033[46m"
	BackColorWhite   = "\033[47m"

	Bold      = "\033[1m"
	Underline = "\033[4m"
	Blink     = "\033[5m"
	Reverse   = "\033[7m"
	UnReverse = "\033[27m"
)

// ResetText sets text parameters to default
func ResetText() string {
	return TextReset
}

func TextReverse(args ...interface{}) string {
	return Reverse + fmt.Sprint(args...)
}

func TextUnReverse(args ...interface{}) string {
	return UnReverse + fmt.Sprint(args...)
}

func TextBold(args ...interface{}) string {
	return Bold + fmt.Sprint(args...) + TextReset
}

func TextUnderline(args ...interface{}) string {
	return Underline + fmt.Sprint(args...) + TextReset
}

func TextBlink(args ...interface{}) string {
	return Blink + fmt.Sprint(args...) + TextReset
}

//Foreground

func ForeRed(args ...interface{}) string {
	return TextColorRed + fmt.Sprint(args...) + TextReset
}

func ForeGreen(args ...interface{}) string {
	return TextColorGreen + fmt.Sprint(args...) + TextReset
}

func ForeYellow(args ...interface{}) string {
	return TextColorYellow + fmt.Sprint(args...) + TextReset
}

func ForeBlue(args ...interface{}) string {
	return TextColorBlue + fmt.Sprint(args...) + TextReset
}

func ForeMagenta(args ...interface{}) string {
	return TextColorMagenta + fmt.Sprint(args...) + TextReset
}

func ForeCyan(args ...interface{}) string {
	return TextColorCyan + fmt.Sprint(args...) + TextReset
}

func ForeGray(args ...interface{}) string {
	return TextColorGray + fmt.Sprint(args...) + TextReset
}

func ForeWhite(args ...interface{}) string {
	return TextColorWhite + fmt.Sprint(args...) + TextReset
}

// Background

func BackRed(args ...interface{}) string {
	return BackColorRed + fmt.Sprint(args...) + TextReset
}

func BackGreen(args ...interface{}) string {
	return BackColorGreen + fmt.Sprint(args...) + TextReset
}

func BackYellow(args ...interface{}) string {
	return BackColorYellow + fmt.Sprint(args...) + TextReset
}

func BackBlue(args ...interface{}) string {
	return BackColorBlue + fmt.Sprint(args...) + TextReset
}

func BackMagenta(args ...interface{}) string {
	return BackColorMagenta + fmt.Sprint(args...) + TextReset
}

func BackCyan(args ...interface{}) string {
	return BackColorCyan + fmt.Sprint(args...) + TextReset
}

func BackBlack(args ...interface{}) string {
	return BackColorBlack + fmt.Sprint(args...) + TextReset
}

func BackWhite(args ...interface{}) string {
	return BackColorWhite + fmt.Sprint(args...) + TextReset
}
