package console

import "fmt"

var (
	TextColorReset  = "\033[0m"
	TextColorRed    = "\033[31m"
	TextColorGreen  = "\033[32m"
	TextColorYellow = "\033[33m"
	TextColorBlue   = "\033[34m"
	TextColorPurple = "\033[35m"
	TextColorCyan   = "\033[36m"
	TextColorGray   = "\033[37m"
	TextColorWhite  = "\033[97m"
)

func ResetColor() string {
	return TextColorReset
}

func FormatRed(args ...interface{}) string {
	return TextColorRed + fmt.Sprint(args...) + TextColorReset
}

func FormatGreen(args ...interface{}) string {
	return TextColorGreen + fmt.Sprint(args...) + TextColorReset
}

func FormatYellow(args ...interface{}) string {
	return TextColorYellow + fmt.Sprint(args...) + TextColorReset
}

func FormatBlue(args ...interface{}) string {
	return TextColorBlue + fmt.Sprint(args...) + TextColorReset
}

func FormatPurple(args ...interface{}) string {
	return TextColorPurple + fmt.Sprint(args...) + TextColorReset
}

func FormatCyan(args ...interface{}) string {
	return TextColorCyan + fmt.Sprint(args...) + TextColorReset
}

func FormatGray(args ...interface{}) string {
	return TextColorGray + fmt.Sprint(args...) + TextColorReset
}

func FormatWhite(args ...interface{}) string {
	return TextColorWhite + fmt.Sprint(args...) + TextColorReset
}
