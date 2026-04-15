package components

// Logo holds ASCII art and color info for a distro.
type Logo struct {
	Art          []string // Lines of ASCII art (with ANSI color codes embedded)
	Width        int      // Visual width of the logo (excluding ANSI codes)
	PrimaryColor string   // ANSI code for the info label color
}

// GetLogo returns the ASCII art logo for the given distro name.
func GetLogo(distro string) *Logo {
	switch distro {
	case "Arch Linux":
		return archLogo()
	case "Ubuntu":
		return ubuntuLogo()
	case "Debian":
		return debianLogo()
	case "Fedora":
		return fedoraLogo()
	case "Mint":
		return mintLogo()
	case "Manjaro":
		return manjaroLogo()
	case "Gentoo":
		return gentooLogo()
	case "CentOS":
		return centosLogo()
	case "openSUSE":
		return opensuseLogo()
	case "Void Linux":
		return voidLogo()
	case "Kali Linux":
		return kaliLogo()
	case "NixOS":
		return nixosLogo()
	case "EndeavourOS":
		return endeavourLogo()
	default:
		return tuxLogo()
	}
}

func archLogo() *Logo {
	c1 := "\033[1;36m"  // light cyan
	c2 := "\033[0;36m"  // cyan
	return &Logo{
		Width:        38,
		PrimaryColor: c1,
		Art: []string{
			c1 + "                   -`                 ",
			c1 + "                  .o+`                ",
			c1 + "                 `ooo/                ",
			c1 + "                `+oooo:               ",
			c1 + "               `+oooooo:              ",
			c1 + "               -+oooooo+:             ",
			c1 + "             `/:-:++oooo+:            ",
			c1 + "            `/++++/+++++++:           ",
			c1 + "           `/++++++++++++++:          ",
			c1 + "          `/+++o" + c2 + "oooooooo" + c1 + "oooo/`        ",
			c2 + "         " + c1 + "./" + c2 + "ooosssso++osssssso" + c1 + "+`       ",
			c2 + "        .oossssso-````/ossssss+`      ",
			c2 + "       -osssssso.      :ssssssso.     ",
			c2 + "      :osssssss/        osssso+++.    ",
			c2 + "     /ossssssss/        +ssssooo/-    ",
			c2 + "   `/ossssso+/:-        -:/+osssso+-  ",
			c2 + "  `+sso+:-`                 `.-/+oso: ",
			c2 + " `++:.                           `-/+/",
			c2 + " .`                                 `/",
		},
	}
}

func ubuntuLogo() *Logo {
	c1 := "\033[1;37m"  // white
	c2 := "\033[1;31m"  // light red
	c3 := "\033[1;33m"  // yellow
	return &Logo{
		Width:        38,
		PrimaryColor: c2,
		Art: []string{
			c2 + "                          ./+o+-      ",
			c1 + "                  yyyyy- " + c2 + "-yyyyyy+     ",
			c1 + "               " + c1 + "://+//////" + c2 + "-yyyyyyo     ",
			c3 + "           .++ " + c1 + ".:/++++++/-" + c2 + ".+sss/`     ",
			c3 + "         .:++o:  " + c1 + "/++++++++/:--:/-     ",
			c3 + "        o:+o+:++." + c1 + "`..```.-/oo+++++/    ",
			c3 + "       .:+o:+o/." + c1 + "          `+sssoo+/   ",
			c1 + "  .++/+:" + c3 + "+oo+o:`" + c1 + "             /sssooo.  ",
			c1 + " /+++//+:" + c3 + "`oo+o" + c1 + "               /::--:.  ",
			c1 + " \\+/+o+++" + c3 + "`o++o" + c2 + "               ++////.  ",
			c1 + "  .++.o+" + c3 + "++oo+:`" + c2 + "             /dddhhh.  ",
			c3 + "       .+.o+oo:." + c2 + "          `oddhhhh+   ",
			c3 + "        \\+.++o+o`" + c2 + "`-```.:ohdhhhhh+    ",
			c3 + "         `:o+++ " + c2 + "`ohhhhhhhhyo++os:     ",
			c3 + "           .o:" + c2 + "`.syhhhhhhh/" + c3 + ".oo++o`     ",
			c2 + "               /osyyyyyyo" + c3 + "++ooo+++/    ",
			c2 + "                   ````` " + c3 + "+oo+++o\\:    ",
			c3 + "                          `oo++.      ",
		},
	}
}

func debianLogo() *Logo {
	c1 := "\033[1;37m"  // white
	c2 := "\033[1;31m"  // light red
	return &Logo{
		Width:        32,
		PrimaryColor: c2,
		Art: []string{
			c1 + "         _,met$$$$$gg.          ",
			c1 + "      ,g$$$$$$$$$$$$$$$P.       ",
			c1 + "    ,g$$P\"\"       \"\"\"Y$$.\".     ",
			c1 + "   ,$$P'              `$$$.     ",
			c1 + "  ',$$P       ,ggs.     `$$b:   ",
			c1 + "  `d$$'     ,$P\"'   " + c2 + "." + c1 + "    $$$    ",
			c1 + "   $$P      d$'     " + c2 + "," + c1 + "    $$P    ",
			c1 + "   $$:      $$.   " + c2 + "-" + c1 + "    ,d$$'    ",
			c1 + "   $$;      Y$b._   _,d$P'     ",
			c1 + "   Y$$.    " + c2 + "`." + c1 + "`\"Y$$$$P\"'         ",
			c1 + "   `$$b      " + c2 + "\"-.__              ",
			c1 + "    `Y$$                        ",
			c1 + "     `Y$$.                      ",
			c1 + "       `$$b.                    ",
			c1 + "         `Y$$b.                 ",
			c1 + "            `\"\"Y$b._             ",
			c1 + "                `\"\"\"           ",
			c1 + "                                ",
		},
	}
}

func fedoraLogo() *Logo {
	c1 := "\033[1;34m"  // light blue
	c2 := "\033[1;37m"  // white
	return &Logo{
		Width:        38,
		PrimaryColor: c1,
		Art: []string{
			c1 + "             .',;::::;,'.             ",
			c1 + "         .';:cccccccccccc:;,.         ",
			c1 + "      .;cccccccccccccccccccccc;.      ",
			c1 + "    .:cccccccccccccccccccccccccc:.    ",
			c1 + "  .;ccccccccccccc;" + c2 + ".:dddl:." + c1 + ";ccccccc;. ",
			c1 + " .:ccccccccccccc;" + c2 + "OWMKOOXMWd" + c1 + ";ccccccc:.",
			c1 + ".:ccccccccccccc;" + c2 + "KMMc" + c1 + ";cc;" + c2 + "xMMc" + c1 + ";ccccccc:.",
			c1 + ",cccccccccccccc;" + c2 + "MMM." + c1 + ";cc;" + c2 + ";WW:" + c1 + ";cccccccc,",
			c1 + ":cccccccccccccc;" + c2 + "MMM." + c1 + ";cccccccccccccccc:",
			c1 + ":ccccccc;" + c2 + "oxOOOo" + c1 + ";" + c2 + "MMM0OOk." + c1 + ";cccccccccccc:",
			c1 + "cccccc;" + c2 + "0MMKxdd:" + c1 + ";" + c2 + "MMMkddc." + c1 + ";cccccccccccc;",
			c1 + "ccccc;" + c2 + "XM0'" + c1 + ";cccc;" + c2 + "MMM." + c1 + ";cccccccccccccccc'",
			c1 + "ccccc;" + c2 + "MMo" + c1 + ";ccccc;" + c2 + "MMW." + c1 + ";ccccccccccccccc; ",
			c1 + "ccccc;" + c2 + "0MNc." + c1 + "ccc" + c2 + ".xMMd" + c1 + ";ccccccccccccccc;  ",
			c1 + "cccccc;" + c2 + "dNMWXXXWM0:" + c1 + ";cccccccccccccc:,  ",
			c1 + "cccccccc;" + c2 + ".:odl:." + c1 + ";cccccccccccccc:,.   ",
			c1 + ":cccccccccccccccccccccccccccc:'.      ",
			c1 + ".:cccccccccccccccccccccc:;,..         ",
			c1 + "  '::cccccccccccccc::;,.              ",
		},
	}
}

func mintLogo() *Logo {
	c1 := "\033[1;37m"  // white
	c2 := "\033[1;32m"  // light green
	return &Logo{
		Width:        38,
		PrimaryColor: c2,
		Art: []string{
			c2 + "                                      ",
			c2 + " MMMMMMMMMMMMMMMMMMMMMMMMMmds+.       ",
			c2 + " MMm----::-://////////////oymNMd+`    ",
			c2 + " MMd      " + c1 + "/++                " + c2 + "-sNMd:   ",
			c2 + " MMNso/`  " + c1 + "dMM    `.::-. .-::.` " + c2 + ".hMN:  ",
			c2 + " ddddMMh  " + c1 + "dMM   :hNMNMNhNMNMNh: " + c2 + "`NMm  ",
			c2 + "     NMm  " + c1 + "dMM  .NMN/-+MMM+-/NMN` " + c2 + "dMM  ",
			c2 + "     NMm  " + c1 + "dMM  -MMm  `MMM   dMM. " + c2 + "dMM  ",
			c2 + "     NMm  " + c1 + "dMM  -MMm  `MMM   dMM. " + c2 + "dMM  ",
			c2 + "     NMm  " + c1 + "dMM  .mmd  `mmm   yMM. " + c2 + "dMM  ",
			c2 + "     NMm  " + c1 + "dMM`  ..`   ...   ydm. " + c2 + "dMM  ",
			c2 + "     hMM- " + c1 + "+MMd/-------...-:sdds  " + c2 + "dMM  ",
			c2 + "     -NMm- " + c1 + ":hNMNNNmdddddddddy/`  " + c2 + "dMM  ",
			c2 + "      -dMNs-" + c1 + "`-::::-------.``    " + c2 + "dMM  ",
			c2 + "       `/dMNmy+/:-------------:/yMMM  ",
			c2 + "          ./ydNMMMMMMMMMMMMMMMMMMMMM  ",
			c2 + "             \\.MMMMMMMMMMMMMMMMMMM    ",
			c2 + "                                      ",
		},
	}
}

func manjaroLogo() *Logo {
	c1 := "\033[1;32m"  // light green
	return &Logo{
		Width:        33,
		PrimaryColor: c1,
		Art: []string{
			c1 + " \u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588  \u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588    ",
			c1 + " \u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588  \u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588    ",
			c1 + " \u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588  \u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588    ",
			c1 + " \u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588  \u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588    ",
			c1 + " \u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588            \u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588    ",
			c1 + " \u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588  \u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588  \u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588    ",
			c1 + " \u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588  \u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588  \u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588    ",
			c1 + " \u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588  \u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588  \u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588    ",
			c1 + " \u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588  \u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588  \u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588    ",
			c1 + " \u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588  \u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588  \u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588    ",
			c1 + " \u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588  \u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588  \u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588    ",
			c1 + " \u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588  \u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588  \u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588    ",
			c1 + " \u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588  \u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588  \u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588    ",
			c1 + " \u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588  \u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588  \u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588    ",
		},
	}
}

func gentooLogo() *Logo {
	c1 := "\033[1;37m"  // white
	c2 := "\033[1;35m"  // light purple
	return &Logo{
		Width:        37,
		PrimaryColor: c2,
		Art: []string{
			c2 + "         -/oyddmdhs+:.               ",
			c2 + "     -o" + c1 + "dNMMMMMMMMNNmhy+" + c2 + "-`            ",
			c2 + "   -y" + c1 + "NMMMMMMMMMMMNNNmmdhy" + c2 + "+-          ",
			c2 + " `o" + c1 + "mMMMMMMMMMMMMNmdmmmmddhhy" + c2 + "/`       ",
			c2 + " om" + c1 + "MMMMMMMMMMMN" + c2 + "hhyyyo" + c1 + "hmdddhhhd" + c2 + "o`     ",
			c2 + ".y" + c1 + "dMMMMMMMMMMd" + c2 + "hs++so/s" + c1 + "mdddhhhhdm" + c2 + "+`   ",
			c2 + " oy" + c1 + "hdmNMMMMMMMN" + c2 + "dyooy" + c1 + "dmddddhhhhyhN" + c2 + "d.  ",
			c2 + "  :o" + c1 + "yhhdNNMMMMMMMNNNmmdddhhhhhyym" + c2 + "Mh  ",
			c2 + "    .:" + c1 + "+sydNMMMMMNNNmmmdddhhhhhhmM" + c2 + "my  ",
			c2 + "       /m" + c1 + "MMMMMMNNNmmmdddhhhhhmMNh" + c2 + "s:  ",
			c2 + "    `o" + c1 + "NMMMMMMMNNNmmmddddhhdmMNhs" + c2 + "+`   ",
			c2 + "  `s" + c1 + "NMMMMMMMMNNNmmmdddddmNMmhs" + c2 + "/.     ",
			c2 + " /N" + c1 + "MMMMMMMMNNNNmmmdddmNMNdso" + c2 + ":`       ",
			c2 + "+M" + c1 + "MMMMMMNNNNNmmmmdmNMNdso" + c2 + "/-          ",
			c2 + "yM" + c1 + "MNNNNNNNmmmmmNNMmhs+/" + c2 + "-`            ",
			c2 + "/h" + c1 + "MMNNNNNNNNMNdhs++/" + c2 + "-`               ",
			c2 + "`/" + c1 + "ohdmmddhys+++/:" + c2 + ".`                  ",
			c2 + "  `-//////:--.                       ",
		},
	}
}

func centosLogo() *Logo {
	c1 := "\033[1;33m"  // yellow
	c2 := "\033[1;32m"  // light green
	c3 := "\033[1;34m"  // light blue
	c4 := "\033[1;35m"  // light purple
	return &Logo{
		Width:        40,
		PrimaryColor: c1,
		Art: []string{
			c1 + "                   ..                   ",
			c1 + "                 .PLTJ.                 ",
			c1 + "                <><><><>                ",
			c2 + "       KKSSV' 4KKK " + c1 + "LJ" + c4 + " KKKL.'VSSKK       ",
			c2 + "       KKV' 4KKKKK " + c1 + "LJ" + c4 + " KKKKAL 'VKK       ",
			c2 + "       V' ' 'VKKKK " + c1 + "LJ" + c4 + " KKKKV' ' 'V       ",
			c2 + "       .4MA.' 'VKK " + c1 + "LJ" + c4 + " KKV' '.4Mb.       ",
			c4 + "     . " + c2 + "KKKKKA.' 'V " + c1 + "LJ" + c4 + " V' '.4KKKKK " + c3 + ".    ",
			c4 + "   .4D " + c2 + "KKKKKKKA.'' " + c1 + "LJ" + c4 + " ''.4KKKKKKK " + c3 + "FA.  ",
			c4 + "  <QDD ++++++++++++  " + c3 + "++++++++++++ GFD>",
			c4 + "   'VD " + c3 + "KKKKKKKK'.. " + c2 + "LJ " + c1 + "..'KKKKKKKK " + c3 + "FV   ",
			c4 + "     ' " + c3 + "VKKKKK'. .4 " + c2 + "LJ " + c1 + "K. .'KKKKKV " + c3 + "'    ",
			c3 + "        'VK'. .4KK " + c2 + "LJ " + c1 + "KKA. .'KV'        ",
			c3 + "       A. . .4KKKK " + c2 + "LJ " + c1 + "KKKKA. . .4       ",
			c3 + "       KKA. 'KKKKK " + c2 + "LJ " + c1 + "KKKKK' .4KK       ",
			c3 + "       KKSSA. VKKK " + c2 + "LJ " + c1 + "KKKV .4SSKK       ",
			c2 + "                <><><><>                ",
			c2 + "                 'MKKM'                 ",
			c2 + "                   ''                   ",
		},
	}
}

func opensuseLogo() *Logo {
	c1 := "\033[1;32m"  // light green
	c2 := "\033[1;37m"  // bold white
	return &Logo{
		Width:        44,
		PrimaryColor: c1,
		Art: []string{
			c2 + "             .;ldkO0000Okdl;.              ",
			c2 + "         .;d00xl:^''''''^:ok00d;.          ",
			c2 + "       .d00l'                'o00d.        ",
			c2 + "     .d0K^'" + c1 + "  Okxoc;:,.          " + c2 + "^O0d.     ",
			c2 + "    .OVV" + c1 + "AK0kOKKKKKKKKKKOxo:,      " + c2 + "lKO.    ",
			c2 + "   ,0VV" + c1 + "AKKKKKKKKKKKKK0P^" + c2 + ",,," + c1 + "^dx:" + c2 + "    ;00,   ",
			c2 + "  .OVV" + c1 + "AKKKKKKKKKKKKKk'" + c2 + ".oOPPb." + c1 + "'0k." + c2 + "   cKO.  ",
			c2 + "  :KV" + c1 + "AKKKKKKKKKKKKKK: " + c2 + "kKx..dd " + c1 + "lKd" + c2 + "   'OK:  ",
			c2 + "  lKl" + c1 + "KKKKKKKKKOx0KKKd " + c2 + "^0KKKO' " + c1 + "kKKc" + c2 + "   lKl  ",
			c2 + "  lKl" + c1 + "KKKKKKKKKK;.;oOKx,.." + c2 + "^" + c1 + "..;kKKK0." + c2 + "  lKl  ",
			c2 + "  :KA" + c1 + "lKKKKKKKKK0o;...^cdxxOK0O/^^'  " + c2 + ".0K:  ",
			c2 + "   kKA" + c1 + "VKKKKKKKKKKKK0x;,,......,;od  " + c2 + "lKP   ",
			c2 + "   '0KA" + c1 + "VKKKKKKKKKKKKKKKKKK00KKOo^  " + c2 + "c00'   ",
			c2 + "    'kKA" + c1 + "VOxddxkOO00000Okxoc;''   " + c2 + ".dKV'    ",
			c2 + "      l0Ko.                    .c00l'      ",
			c2 + "       'l0Kk:.              .;xK0l'        ",
			c2 + "          'lkK0xc;:,,,,:;odO0kl'           ",
			c2 + "              '^:ldxkkkkxdl:^'              ",
		},
	}
}

func voidLogo() *Logo {
	c1 := "\033[0;32m"  // green
	c2 := "\033[1;32m"  // light green
	c3 := "\033[1;30m"  // dark grey
	return &Logo{
		Width:        47,
		PrimaryColor: c2,
		Art: []string{
			c2 + "                 __.;=====;.__                 ",
			c2 + "             _.=+==++=++=+=+===;.             ",
			c2 + "              -=+++=+===+=+=+++++=_           ",
			c1 + "         .     " + c2 + "-=:`     `--==+=++==.          ",
			c1 + "        _vi,    " + c2 + "`            --+=++++:       ",
			c1 + "       .uvnvi.       " + c2 + "_._       -==+==+.     ",
			c1 + "      .vvnvnI`    " + c2 + ".;==|==;.     :|=||=|.   ",
			c3 + " +QmQQm" + c1 + "pvvnv; " + c3 + "_yYsyQQWUUQQQm #QmQ#" + c2 + ":" + c3 + "QQQWUV$QQmL",
			c3 + "  -QQWQW" + c1 + "pvvo" + c3 + "wZ?.wQQQE" + c2 + "==<" + c3 + "QWWQ/QWQW.QQWW" + c2 + "(: " + c3 + "jQWQE",
			c3 + "   -$QQQQmmU'  jQQQ@" + c2 + "+=<" + c3 + "QWQQ)mQQQ.mQQQC" + c2 + "+;" + c3 + "jWQQ@'",
			c3 + "    -$WQ8Y" + c1 + "nI:   " + c3 + "QWQQwgQQWV" + c2 + "`" + c3 + "mWQQ.jQWQQgyyWW@! ",
			c1 + "      -1vvnvv.     " + c2 + "`~+++`        ++|+++   ",
			c1 + "       +vnvnnv,                 " + c2 + "`-|===    ",
			c1 + "        +vnvnvns.           .      " + c2 + ":=-   ",
			c1 + "         -Invnvvnsi..___..=sv=.     " + c2 + "`    ",
			c1 + "           +Invnvnvnnnnnnnnvvnn;.          ",
			c1 + "             ~|Invnvnvvnvvvnnv}+`          ",
			c1 + "                -~\"|{*l}*|\"\"~               ",
		},
	}
}

func kaliLogo() *Logo {
	c1 := "\033[1;34m"  // light blue
	c2 := "\033[0;30m"  // black
	return &Logo{
		Width:        48,
		PrimaryColor: c1,
		Art: []string{
			c1 + "..............                                  ",
			c1 + "            ..,;:ccc,.                          ",
			c1 + "          ......''';lxO.                        ",
			c1 + ".....''''..........,:ld;                        ",
			c1 + "           .';;;:::;,,.x,                      ",
			c1 + "      ..'''.            0Xxoc:,.  ...           ",
			c1 + "  ....                ,ONkc;,;cokOdc',.        ",
			c1 + " .                   OMo           ':" + c2 + "dd" + c1 + "o.    ",
			c1 + "                    dMc               :OO;     ",
			c1 + "                    0M.                 .:o.   ",
			c1 + "                    ;Wd                        ",
			c1 + "                     ;XO,                      ",
			c1 + "                       ,d0Odlc;,..             ",
			c1 + "                           ..',;:cdOOd::,.     ",
			c1 + "                                    .:d;.':;.  ",
			c1 + "                                       'd,  .'",
			c1 + "                                         ;l   ..",
			c1 + "                                          .o    ",
			c1 + "                                            c   ",
			c1 + "                                            .'  ",
			c1 + "                                             .  ",
		},
	}
}

func nixosLogo() *Logo {
	c1 := "\033[0;34m"  // blue
	c2 := "\033[1;34m"  // light blue
	return &Logo{
		Width:        45,
		PrimaryColor: c2,
		Art: []string{
			c1 + "          ::::.    " + c2 + "':::::     ::::'          ",
			c1 + "          ':::::    " + c2 + "':::::.  ::::'           ",
			c1 + "            :::::     " + c2 + "'::::.:::::            ",
			c1 + "      .......:::::..... " + c2 + "::::::::             ",
			c1 + "     ::::::::::::::::::. " + c2 + "::::::    " + c1 + "::::.    ",
			c1 + "    ::::::::::::::::::::: " + c2 + ":::::.  " + c1 + ".::::'   ",
			c2 + "           .....           ::::' " + c1 + ":::::'       ",
			c2 + "          :::::            '::' " + c1 + ":::::'        ",
			c2 + " ........:::::               ' " + c1 + ":::::::::::.  ",
			c2 + ":::::::::::::                 " + c1 + ":::::::::::::  ",
			c2 + " ::::::::::: " + c1 + "..              :::::           ",
			c2 + "     .::::: " + c1 + ".:::            :::::            ",
			c2 + "    .:::::  " + c1 + ":::::          '''''    " + c2 + ".....   ",
			c2 + "    :::::   " + c1 + "':::::.  " + c2 + "......:::::::::::::'   ",
			c2 + "     :::     " + c1 + "::::::. " + c2 + "':::::::::::::::::'    ",
			c1 + "            .:::::::: " + c2 + "'::::::::::              ",
			c1 + "           .::::''::::.     " + c2 + "'::::.             ",
			c1 + "          .::::'   ::::.     " + c2 + "'::::.            ",
			c1 + "         .::::      ::::      " + c2 + "'::::.           ",
		},
	}
}

func endeavourLogo() *Logo {
	c1 := "\033[1;33m"  // yellow
	c3 := "\033[1;35m"  // purple
	c5 := "\033[1;36m"  // cyan
	return &Logo{
		Width:        44,
		PrimaryColor: c3,
		Art: []string{
			c1 + "                  +" + c3 + "I" + c5 + "+                        ",
			c1 + "                 +" + c3 + "777" + c5 + "+                       ",
			c1 + "                +" + c3 + "77777" + c5 + "++                     ",
			c1 + "               +" + c3 + "7777777" + c5 + "++                    ",
			c1 + "              +" + c3 + "7777777777" + c5 + "++                  ",
			c1 + "            ++" + c3 + "7777777777777" + c5 + "++                ",
			c1 + "           ++" + c3 + "777777777777777" + c5 + "+++              ",
			c1 + "         ++" + c3 + "77777777777777777" + c5 + "++++             ",
			c1 + "        ++" + c3 + "7777777777777777777" + c5 + "++++            ",
			c1 + "      +++" + c3 + "777777777777777777777" + c5 + "++++           ",
			c1 + "    ++++" + c3 + "7777777777777777777777" + c5 + "+++++          ",
			c1 + "   ++++" + c3 + "77777777777777777777777" + c5 + "+++++         ",
			c1 + "  +++++" + c3 + "777777777777777777777777" + c5 + "+++++        ",
			c5 + "       +++++++" + c3 + "7777777777777777" + c5 + "++++++        ",
			c5 + "      +++++++++++++++++++++++++++++      ",
			c5 + "     +++++++++++++++++++++++++++         ",
		},
	}
}

func tuxLogo() *Logo {
	c1 := "\033[1;37m"  // white
	c2 := "\033[1;30m"  // dark grey
	c3 := "\033[1;33m"  // yellow
	return &Logo{
		Width:        28,
		PrimaryColor: c1,
		Art: []string{
			c2 + "                            ",
			c2 + "                            ",
			c2 + "                            ",
			c2 + "         #####              ",
			c2 + "        #######             ",
			c2 + "        ##" + c1 + "O" + c2 + "#" + c1 + "O" + c2 + "##             ",
			c2 + "        #" + c3 + "#####" + c2 + "#             ",
			c2 + "      ##" + c1 + "##" + c3 + "###" + c1 + "##" + c2 + "##           ",
			c2 + "     #" + c1 + "##########" + c2 + "##          ",
			c2 + "    #" + c1 + "############" + c2 + "##         ",
			c2 + "    #" + c1 + "############" + c2 + "###        ",
			c3 + "   ##" + c2 + "#" + c1 + "###########" + c2 + "##" + c3 + "#       ",
			c3 + " ######" + c2 + "#" + c1 + "#######" + c2 + "#" + c3 + "######    ",
			c3 + " #######" + c2 + "#" + c1 + "#####" + c2 + "#" + c3 + "#######    ",
			c3 + "   #####" + c2 + "#######" + c3 + "#####      ",
			c2 + "                            ",
			c2 + "                            ",
			c2 + "                            ",
		},
	}
}
