package gocolor

import "fmt"

var colors = map[string]string{
	"red":     "31",
	"green":   "32",
	"blue":    "34",
	"yellow":  "33",
	"magenta": "35",
	"gray":    "243",
	"cyan":    "36",
	"black":   "30",
	"white":   "37",
	"reset":   "0",
}

var mode = map[string]string{
	"bold":      "1",
	"italic":    "3",
	"underline": "4",
	"blink":     "5",
	"strike":    "9",
	"reset":     "0",
}

func ColorizeText(str *string, color, format string) {
	inputcheck(&color, &format)
	*str = fmt.Sprintf("%s%s%s", color, *str, "\x1b[0;0m")
}


func Colorize(str *string, color, format string) {
	inputcheck(&color, &format)
	*str = fmt.Sprintf("%s%s", color, *str)
}

func SetColor(color, format string) {
	inputcheck(&color, &format)
	fmt.Printf(color)
}

func ColorString(str, color, format string ) string {
	inputcheck(&color, &format)
	return fmt.Sprintf("%s%s%s", color, str, "\x1b[0;0m")
}

func inputcheck(color, format *string) {
	*format = mode[*format]
	*color = colors[*color]
	if *color == "" {
		*color = "37"
	}
	if *format == "" {
		*format = "0"
	}
	*color = fmt.Sprintf("\x1b[%s;%sm", *format, *color)
}