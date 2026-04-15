package components

import (
	"fmt"
	"sync"
)

func colorizeAsString(colorCode int) string {
	return fmt.Sprintf("\x1b[0m\x1b[38;5;%dm", colorCode)
}

func colorizePrint(colorCode int) {
	fmt.Printf("\x1b[0m\x1b[38;5;%dm", colorCode)
}

var (
	nameToCodeOnce sync.Once
	nameToCodeMap  map[string]string
	
	codeToNameOnce sync.Once
	codeToNameMap  map[string]string
)

func initNameToCodeMap() {
	nameToCodeOnce.Do(func() {
		nameToCodeMap = map[string]string{
			"black":        "\x1b[0m\x1b[30m",
			"red":          "\x1b[0m\x1b[31m",
			"green":        "\x1b[0m\x1b[32m",
			"brown":        "\x1b[0m\x1b[33m",
			"blue":         "\x1b[0m\x1b[34m",
			"purple":       "\x1b[0m\x1b[35m",
			"cyan":         "\x1b[0m\x1b[36m",
			"yellow":       "\x1b[0m\x1b[1;33m",
			"white":        "\x1b[0m\x1b[1;37m",
			"dark grey":    "\x1b[0m\x1b[1;30m",
			"dark gray":    "\x1b[0m\x1b[1;30m",
			"light red":    "\x1b[0m\x1b[1;31m",
			"light green":  "\x1b[0m\x1b[1;32m",
			"light blue":   "\x1b[0m\x1b[1;34m",
			"light purple": "\x1b[0m\x1b[1;35m",
			"light cyan":   "\x1b[0m\x1b[1;36m",
			"light grey":   "\x1b[0m\x1b[37m",
			"light gray":   "\x1b[0m\x1b[37m",
			"orange":       colorizeAsString(202),
			"light orange": colorizeAsString(214),
			"black_haiku":  colorizeAsString(7),
			"rosa_blue":    "\x1b[01;38;05;25m",
			"arco_blue":    "\x1b[1;38;05;111m",
		}
	})
}

func initCodeToNameMap() {
	codeToNameOnce.Do(func() {
		initNameToCodeMap()
		codeToNameMap = make(map[string]string)
		
		preferredNames := []string{
			"black", "red", "green", "brown", "blue", "purple", "cyan",
			"yellow", "white", "dark grey", "light red", "light green",
			"light blue", "light purple", "light cyan", "light grey",
			"orange", "light orange", "black_haiku", "rosa_blue", "arco_blue",
		}
		
		for _, name := range preferredNames {
			if code, ok := nameToCodeMap[name]; ok {
				codeToNameMap[code] = name
			}
		}
	})
}

func GetColorByName(name string) string {
	initNameToCodeMap()
	return nameToCodeMap[name]
}

func GetNameByCode(code string) string {
	initCodeToNameMap()
	return codeToNameMap[code]
}