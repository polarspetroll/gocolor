## functions

[func Colorize(str *string, color, format string)](#Colorize)

[func ColorizeText(str *string, color, format string)](#ColorizeText)

[func SetColor(color, format string)](#SetColor)

[func ColorString(str, color, format string ) string ](#ColorString)


## variables
[var colors map[string]string](#colors)

[var mode map[string]string](#modes)



---

##### [Colorize](https://github.com/polarspetroll/gocolor/blob/00079bb93bc4f0a7c218eedda5b8a98a783b86d2/gocolor.go#L33)
colorize a variable (type: string)
##### [ColorizeText](https://github.com/polarspetroll/gocolor/blob/00079bb93bc4f0a7c218eedda5b8a98a783b86d2/gocolor.go#L27)
same as colorize but the function will reset the color at the end
##### [SetColor](https://github.com/polarspetroll/gocolor/blob/00079bb93bc4f0a7c218eedda5b8a98a783b86d2/gocolor.go#L38)
set the console color

##### [ColorString](https://github.com/polarspetroll/gocolor/blob/00079bb93bc4f0a7c218eedda5b8a98a783b86d2/gocolor.go#L43)

gets a string and returns the colorized string

---

##### [colors](https://github.com/polarspetroll/gocolor/blob/00079bb93bc4f0a7c218eedda5b8a98a783b86d2/gocolor.go#L5)
```go

var colors = map[string]string{
    //ANSI escape codes
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
```

##### [modes](https://github.com/polarspetroll/gocolor/blob/00079bb93bc4f0a7c218eedda5b8a98a783b86d2/gocolor.go#L18)

```go
var mode = map[string]string{
	"bold":      "1",
	"italic":    "3",
	"underline": "4",
	"blink":     "5",
	"strike":    "9",
	"reset":     "0",
}
```
