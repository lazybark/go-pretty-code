# Go pretty code
A pack of simple solutions for making console output prettier.
## Installation
```
go get github.com/lazybark/go-pretty-code
```
## Text colors for console output
```
go get github.com/lazybark/go-pretty-code/console
```
Foreground & background:
- ![#f03c15](https://via.placeholder.com/15/f03c15/000000?text=+) `console.ForeRed / BackRed`
- ![#00FF00](https://via.placeholder.com/15/00FF00/000000?text=+) `console.ForeGreen / BackGreen`
- ![#FFFF00](https://via.placeholder.com/15/FFFF00/000000?text=+) `console.ForeYellow / BackYellow`
- ![#0000FF](https://via.placeholder.com/15/0000FF/000000?text=+) `console.ForeBlue / BackBlue`
- ![#FF00FF](https://via.placeholder.com/15/FF00FF/000000?text=+) `console.ForeMagenta / BackMagenta`
- ![#00FFFF](https://via.placeholder.com/15/00FFFF/000000?text=+) `console.ForeCyan / BackCyan`
- ![#C0C0C0](https://via.placeholder.com/15/C0C0C0/000000?text=+) `console.ForeGray`
- ![#FFFFFF](https://via.placeholder.com/15/FFFFFF/000000?text=+) `console.ForeWhite / BackWhite`
- ![#000000](https://via.placeholder.com/15/000000/000000?text=+) `console.BackBlack`

### Combined formating
```
fmt.Println(console.ForeBlue(console.BackWhite("Blue & white")))
fmt.Println(console.ForeGreen(console.TextBlink("Blinking green")))
```
### Using variables to set color manually
```
console.TextColorRed + "Colored text can be splitted by " + console.TextReset
console.TextColorRed + "Colored text can be splitted by " + console.ResetText()
```
## Colorful logging using Uber Zap
