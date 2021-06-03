package main
import (
	"fmt"
	"github.com/polarspetroll/gocolor"
)


func main() {
	fmt.Println(gocolor.ColorString("bold", "", "bold"))
	fmt.Println(gocolor.ColorString("italic", "", "italic"))
	fmt.Println(gocolor.ColorString("underline", "", "underline"))
	fmt.Println(gocolor.ColorString("blink", "", "blink"))
	fmt.Println(gocolor.ColorString("strike", "", "strike"))
}
