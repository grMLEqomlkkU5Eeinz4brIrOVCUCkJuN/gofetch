package constants

import (
	"sync"
)

var (
	supportedDistrosOnce sync.Once
	supportedDistros     string

	supportedOtherOnce sync.Once
	supportedOther     string

	supportedDEsOnce sync.Once
	supportedDEs     string

	supportedWMsOnce sync.Once
	supportedWMs     string
)

func GetSupportedDistros() string {
	supportedDistrosOnce.Do(func() {
		supportedDistros = "ALDOS, Alpine Linux, AlmaLinux, Alter Linux, Amazon Linux, Antergos, Arch Linux (Old and Current Logos), " +
			"Arch Linux 32, ArcoLinux, Artix Linux, blackPanther OS, BLAG, BunsenLabs, CentOS, Chakra, Chapeau, Chrome OS, " +
			"Chromium OS, CrunchBang, CRUX, Debian, Deepin, DesaOS, Devuan, Dragora, DraugerOS, elementary OS, EuroLinux, " +
			"Evolve OS, Sulin, Exherbo, Fedora (Old and Current Logos), Frugalware, Fuduntu, Funtoo, Fux, Gentoo, gNewSense, " +
			"Guix System, Hyperbola GNU/Linux-libre, januslinux, Jiyuu Linux, Kali Linux, KaOS, KDE neon, Kogaion, Korora, " +
			"LinuxDeepin, Linux Mint, LMDE, Logos, Mageia, Mandriva/Mandrake, Manjaro, Mer, Netrunner, NixOS, OBRevenge, " +
			"openSUSE, OS Elbrus, Oracle Linux, Parabola GNU/Linux-libre, Pardus, Parrot Security, PCLinuxOS, PeppermintOS, " +
			"Proxmox VE, PureOS, Quirinux, Qubes OS, Raspbian, Red Hat Enterprise Linux, Rocky Linux, ROSA, Sabayon, SailfishOS, " +
			"Scientific Linux, Siduction, Slackware, Solus, Source Mage GNU/Linux, SparkyLinux, SteamOS, SUSE Linux Enterprise, " +
			"SwagArch, TeArch, TinyCore, Trisquel, Ubuntu, Viperr, Void Linux, Zorin OS, and EndeavourOS"
	})
	return supportedDistros
}

func GetSupportedOther() string {
	supportedOtherOnce.Do(func() {
		supportedOther = "Dragonfly/Free/Open/Net BSD, Haiku, macOS, Windows+Cygwin and Windows+MSYS2."
	})
	return supportedOther
}

func GetSupportedDEs() string {
	supportedDEsOnce.Do(func() {
		supportedDEs = "KDE, GNOME, Unity, Xfce, LXDE, Cinnamon, MATE, Deepin, CDE, RazorQt and Trinity."
	})
	return supportedDEs
}

func GetSupportedWMs() string {
	supportedWMsOnce.Do(func() {
		supportedWMs = "2bwm, 9wm, Awesome, Beryl, Blackbox, Cinnamon, chromeos-wm, Compiz, deepin-wm, " +
			"dminiwm, dwm, dtwm, E16, E17, echinus, Emerald, FluxBox, FLWM, FVWM, herbstluftwm, howm, IceWM, KWin, " +
			"Metacity, monsterwm, Musca, Gala, Mutter, Muffin, Notion, OpenBox, PekWM, Ratpoison, Sawfish, ScrotWM, SpectrWM, " +
			"StumpWM, subtle, sway, TWin, WindowMaker, WMFS, wmii, Xfwm4, XMonad and i3."
	})
	return supportedWMs
}