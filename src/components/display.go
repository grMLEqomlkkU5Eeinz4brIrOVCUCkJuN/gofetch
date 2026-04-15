package components

import (
	"fmt"
	"regexp"
	"strings"
)

// ansiRegex strips ANSI escape sequences for width calculation.
var ansiRegex = regexp.MustCompile(`\x1b\[[0-9;]*m`)

// visibleLen returns the visible (non-ANSI) length of a string.
func visibleLen(s string) int {
	return len(ansiRegex.ReplaceAllString(s, ""))
}

// padRight pads a string to the given visible width.
func padRight(s string, width int) string {
	vl := visibleLen(s)
	if vl >= width {
		return s
	}
	return s + strings.Repeat(" ", width-vl)
}

// Display prints the logo and system info side by side.
func Display(info *SystemInfo) {
	logo := GetLogo(info.Distro.Distro)
	infoLines := info.InfoLines()

	// Override the info label color with the logo's primary color
	if logo.PrimaryColor != "" {
		for i, line := range infoLines {
			// Skip the title line (index 0) and separator (index 1)
			if i >= 2 {
				infoLines[i] = strings.Replace(line, "\033[1;34m", logo.PrimaryColor, 1)
			}
		}
	}

	// Determine how many lines we need
	total := len(logo.Art)
	if len(infoLines) > total {
		total = len(infoLines)
	}

	logoWidth := logo.Width + 4 // padding between logo and info

	for i := 0; i < total; i++ {
		logoLine := ""
		if i < len(logo.Art) {
			logoLine = logo.Art[i]
		}

		infoLine := ""
		if i < len(infoLines) {
			infoLine = infoLines[i]
		}

		fmt.Printf("%s%s\n", padRight(logoLine, logoWidth), infoLine)
	}
	fmt.Print("\033[0m") // reset
}
