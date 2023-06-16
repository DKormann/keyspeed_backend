package util

import "fmt"

var m map[string]string = map[string]string{
	"green":     "\033[92m",
	"red":       "\033[91m",
	"yellow":    "\033[93m",
	"grey":      "\033[90m",
	"blue":      "\033[94m",
	"pink":      "\033[95m",
	"cyan":      "\033[96m",
	"white":     "\033[97m",
	"lightgrey": "\033[98m",
}

func ColorLog(color string, a ...any) {

	v, ok := m[color]
	if !ok {
		fmt.Println("color not defined")
		return
	}
	fmt.Print(v)
	fmt.Print(a...)
	fmt.Print("\033[0m")

}

func ColorLogln(color string, a ...any) {
	a = append(a, "\n")
	ColorLog(color, a...)
}

func GLog(a ...any) { ColorLog("green", a...) }
func BLog(a ...any) { ColorLog("blue", a...) }
func YLog(a ...any) { ColorLog("yellow", a...) }
func RLog(a ...any) { ColorLog("red", a...) }

func GLogln(a ...any) { ColorLogln("green", a...) }
func BLogln(a ...any) { ColorLogln("blue", a...) }
func YLogln(a ...any) { ColorLogln("yellow", a...) }
func RLogln(a ...any) { ColorLogln("red", a...) }
