package components

import (
	"fmt"
)

type WindowManager string

const (
	FluxBox      WindowManager = "fluxbox"
	OpenBox      WindowManager = "openbox"
	BlackBox     WindowManager = "blackbox"
	XFWM4        WindowManager = "xfwm4"
	Metacity     WindowManager = "metacity"
	KWin         WindowManager = "kwin"
	Twin         WindowManager = "twin"
	IceWM        WindowManager = "icewm"
	Pekwm        WindowManager = "pekwm"
	FLWM         WindowManager = "flwm"
	FLWMTopside  WindowManager = "flwm_topside"
	FVWM         WindowManager = "fvwm"
	DWM          WindowManager = "dwm"
	Awesome      WindowManager = "awesome"
	WMaker       WindowManager = "wmaker"
	StumpWM      WindowManager = "stumpwm"
	Musca        WindowManager = "musca"
	XMonad       WindowManager = "xmonad.*"
	I3           WindowManager = "i3"
	RatPoison    WindowManager = "ratpoison"
	Scrotwm      WindowManager = "scrotwm"
	Spectrwm     WindowManager = "spectrwm"
	WMFS         WindowManager = "wmfs"
	WMII         WindowManager = "wmii"
	Beryl        WindowManager = "beryl"
	Subtle       WindowManager = "subtle"
	E16          WindowManager = "e16"
	Enlightenment WindowManager = "enlightenment"
	Sawfish      WindowManager = "sawfish"
	Emerald      WindowManager = "emerald"
	MonsterWM    WindowManager = "monsterwm"
	DMiniWM      WindowManager = "dminiwm"
	Compiz       WindowManager = "compiz"
	Finder       WindowManager = "Finder"
	HerbstluftwM WindowManager = "herbstluftwm"
	Howm         WindowManager = "howm"
	Notion       WindowManager = "notion"
	BspWM        WindowManager = "bspwm"
	Cinnamon     WindowManager = "cinnamon"
	TwoBWM       WindowManager = "2bwm"
	Echinus      WindowManager = "echinus"
	SWM          WindowManager = "swm"
	BudgieWM     WindowManager = "budgie-wm"
	DTWM         WindowManager = "dtwm"
	NineWM       WindowManager = "9wm"
	ChromeOSWM   WindowManager = "chromeos-wm"
	DeepaWM      WindowManager = "deepin-wm"
	Sway         WindowManager = "sway"
)

// Metadata about a window manager
type WindowManagerInfo struct {
	ProcessName string // What to search for in running processes
	DisplayName string // What to show to user
	IsRegex     bool   // Whether ProcessName is a regex
}

// Ordered list for detection priority
var windowManagerRegistry = []WindowManagerInfo{
	{ProcessName: "fluxbox", DisplayName: "FluxBox"},
	{ProcessName: "openbox", DisplayName: "OpenBox"},
	{ProcessName: "blackbox", DisplayName: "BlackBox"},
	{ProcessName: "gala", DisplayName: "Gala"},
	{ProcessName: "mutter", DisplayName: "Mutter"},
	{ProcessName: "muffin", DisplayName: "Muffin"},
	{ProcessName: "xfwm4", DisplayName: "Xfwm4"},
	{ProcessName: "metacity", DisplayName: "Metacity"},
	{ProcessName: "kwin", DisplayName: "KWin"},
	{ProcessName: "twin", DisplayName: "TWin"},
	{ProcessName: "icewm", DisplayName: "IceWM"},
	{ProcessName: "pekwm", DisplayName: "PekWM"},
	{ProcessName: "flwm", DisplayName: "FLWM"},
	{ProcessName: "flwm_topside", DisplayName: "FLWM"},
	{ProcessName: "fvwm", DisplayName: "FVWM"},
	{ProcessName: "dwm", DisplayName: "dwm"},
	{ProcessName: "awesome", DisplayName: "Awesome"},
	{ProcessName: "wmaker", DisplayName: "WindowMaker"},
	{ProcessName: "stumpwm", DisplayName: "StumpWM"},
	{ProcessName: "musca", DisplayName: "Musca"},
	{ProcessName: "xmonad.*", DisplayName: "XMonad", IsRegex: true},
	{ProcessName: "i3", DisplayName: "i3"},
	{ProcessName: "ratpoison", DisplayName: "Ratpoison"},
	{ProcessName: "scrotwm", DisplayName: "ScrotWM"},
	{ProcessName: "spectrwm", DisplayName: "SpectrWM"},
	{ProcessName: "wmfs", DisplayName: "WMFS"},
	{ProcessName: "wmii", DisplayName: "wmii"},
	{ProcessName: "beryl", DisplayName: "Beryl"},
	{ProcessName: "subtle", DisplayName: "subtle"},
	{ProcessName: "e16", DisplayName: "E16"},
	{ProcessName: "enlightenment", DisplayName: "E17"},
	{ProcessName: "sawfish", DisplayName: "Sawfish"},
	{ProcessName: "emerald", DisplayName: "Emerald"},
	{ProcessName: "monsterwm", DisplayName: "monsterwm"},
	{ProcessName: "dminiwm", DisplayName: "dminiwm"},
	{ProcessName: "compiz", DisplayName: "Compiz"},
	{ProcessName: "Finder", DisplayName: "Quartz Compositor"},
	{ProcessName: "herbstluftwm", DisplayName: "herbstluftwm"},
	{ProcessName: "howm", DisplayName: "howm"},
	{ProcessName: "notion", DisplayName: "Notion"},
	{ProcessName: "bspwm", DisplayName: "bspwm"},
	{ProcessName: "cinnamon", DisplayName: "Muffin"},
	{ProcessName: "2bwm", DisplayName: "2bwm"},
	{ProcessName: "echinus", DisplayName: "echinus"},
	{ProcessName: "swm", DisplayName: "swm"},
	{ProcessName: "budgie-wm", DisplayName: "BudgieWM"},
	{ProcessName: "dtwm", DisplayName: "dtwm"},
	{ProcessName: "9wm", DisplayName: "9wm"},
	{ProcessName: "chromeos-wm", DisplayName: "chromeos-wm"},
	{ProcessName: "deepin-wm", DisplayName: "Deepin WM"},
	{ProcessName: "sway", DisplayName: "sway"},
	{ProcessName: "mwm", DisplayName: "MWM"},
}

// DetectWindowManager checks running processes for a known window manager
func DetectWindowManager(registry []WindowManagerInfo) (string, error) {
	for _, wm := range registry {
		found, err := isProcessRunning(wm.ProcessName)
		if err != nil {
			// Log the error but continue checking other WMs
			// In production, you might want: log.Printf("error checking %s: %v", wm.DisplayName, err)
			continue
		}

		if found {
			return wm.DisplayName, nil
		}
	}

	return "", fmt.Errorf("no window manager detected")
}