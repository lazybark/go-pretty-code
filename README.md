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
**WARNING**
<br>
Colors will not work in standart Windows console. To get colors on Windows (instead of weird ANSI) use [Windows Terminal](https://docs.microsoft.com/en-us/windows/terminal/install)
## Colorful logging using Uber Zap
Pretty logger in this mod is just a wrap around go.uber.org/zap. It's designed to make logging more flexible for projects that require console+file base output but with different (customizable) rules for each: colors, formatting, etc.
<br>
There are some prepared variants in github.com/lazybark/go-pretty-code/logs but many more can be created just by combining existing functions.
<br>
Basic use: logs.Double(PATH_TO_LOGFILE, TRUNCATE_FILE, ZAP_LOGGING_LEVEL)
```
logger, _ := logs.FileOnly("some/log.txt", false, zap.InfoLevel) // Will use only specified file as output
logger, _ := logs.ConsoleOnly(zap.InfoLevel) // Console-only, for just pretty colorful logs
logger, _ := logs.Double("some/log.txt", false, zap.InfoLevel) // Will use both
```
