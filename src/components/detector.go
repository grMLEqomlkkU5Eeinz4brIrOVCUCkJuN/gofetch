package components

import (
	"bytes"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

// DistroInfo holds information about the detected distribution,
// mirroring the distro/distro_release/distro_codename/distro_more
// variables used by detecfromfile.sh.
type DistroInfo struct {
	Distro   string
	Release  string
	Codename string
	More     string
	WSL      bool
}

// DetectDistro returns the name of the current distribution.
// It is kept as the public entry point for callers that just want a name.
func DetectDistro() string {
	return DetectDistroInfo().Distro
}

// DetectDistroInfo performs distro detection by replicating the logic of
// src/components/detecfromfile.sh as closely as reasonably possible in Go.
func DetectDistroInfo() *DistroInfo {
	info := &DistroInfo{Distro: "Unknown"}

	// MCST / OS Elbrus check
	if data, err := os.ReadFile("/etc/mcst_version"); err == nil {
		info.Distro = "OS Elbrus"
		lines := splitLines(string(data))
		if n := len(lines); n > 0 {
			info.Release = strings.TrimSpace(lines[n-1])
			if info.Release != "" {
				info.More = info.Release
			}
		}
	} else if hasCommand("lsb_release") {
		detectFromLSB(info)
	}

	// Post-lsb fixups on distro_detect matching (applied by the shell script
	// regardless of which case branch was taken)
	applyLSBPostFixups(info)

	// If lsb_release did not yield a distro, try uname -o and friends
	if info.Distro == "Unknown" {
		detectFromUname(info)
	}

	// Cygwin/Msys Windows version detection (best effort)
	if info.Distro == "Cygwin" || info.Distro == "Msys" {
		if ver := runCommand("wmic", "os", "get", "version"); ver != "" {
			for _, line := range splitLines(ver) {
				line = strings.TrimSpace(line)
				if strings.HasPrefix(line, "6.2") || strings.HasPrefix(line, "6.3") || strings.HasPrefix(line, "10") {
					info.More = "Windows - Modern"
				}
			}
		}
	}

	// /etc/os-release based detection with hotfixes
	if info.Distro == "Unknown" {
		detectFromOSRelease(info)
	}

	// Generic /etc/<distro>-release file probing
	if info.Distro == "Unknown" && isLinuxOrGNU() {
		detectFromReleaseFiles(info)
	}

	// Debian/NixOS/Dragora/Slackware/TinyCore/Sabayon + macOS + BSD dmesg
	if info.Distro == "Unknown" {
		detectFromMiscFiles(info)
	}

	// /etc/issue based detection
	if info.Distro == "Unknown" && isLinuxOrGNU() {
		detectFromIssue(info)
	}

	// /etc/system-release and ChromeOS lsb-release
	if info.Distro == "Unknown" && isLinuxOrGNU() {
		detectFromSystemRelease(info)
	}

	// Aggregate "more" string from release/codename
	if info.More == "" {
		if info.Release != "" && info.Release != "n/a" {
			info.More = info.Release
		}
		if info.Codename != "" && info.Codename != "n/a" {
			if info.More != "" {
				info.More += " " + info.Codename
			} else {
				info.More = info.Codename
			}
		}
	}
	if info.More != "" && !strings.HasPrefix(info.More, info.Distro) {
		info.More = info.Distro + " " + info.More
	}

	// Final canonical lowercase-mapped normalization, except Haiku
	if info.Distro != "Haiku" {
		info.Distro = canonicalizeDistro(strings.ToLower(info.Distro))
	}

	// WSL detection
	if containsInFile("/proc/version", "Microsoft") ||
		containsInFile("/proc/sys/kernel/osrelease", "Microsoft") {
		info.WSL = true
	}

	return info
}


func runCommand(name string, args ...string) string {
	cmd := exec.Command(name, args...)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = nil
	if err := cmd.Run(); err != nil {
		return ""
	}
	return strings.TrimSpace(out.String())
}

func hasCommand(name string) bool {
	_, err := exec.LookPath(name)
	return err == nil
}

func splitLines(s string) []string {
	return strings.Split(strings.TrimRight(s, "\n"), "\n")
}

func isLinuxOrGNU() bool {
	return runtime.GOOS == "linux"
}

func containsInFile(path, needle string) bool {
	data, err := os.ReadFile(path)
	if err != nil {
		return false
	}
	return strings.Contains(strings.ToLower(string(data)), strings.ToLower(needle))
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// readKeyValue reads a shell-style KEY=VALUE file and returns the value
// associated with key (quotes stripped).
func readKeyValue(path, key string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		return ""
	}
	prefix := key + "="
	for _, line := range splitLines(string(data)) {
		if strings.HasPrefix(line, prefix) {
			v := strings.TrimPrefix(line, prefix)
			v = strings.TrimSpace(v)
			v = strings.Trim(v, "\"'")
			return v
		}
	}
	return ""
}

// --- detection stages ------------------------------------------------------

func detectFromLSB(info *DistroInfo) {
	detect := runCommand("lsb_release", "-si")
	release := runCommand("lsb_release", "-sr")
	codename := runCommand("lsb_release", "-sc")
	info.Release = release
	info.Codename = codename

	switch detect {
	case "archlinux", "Arch Linux", "arch", "Arch", "archarm":
		info.Distro = "Arch Linux"
		info.Release = "n/a"
		osRelease := ""
		if fileExists("/etc/os-release") {
			osRelease = "/etc/os-release"
		} else if fileExists("/usr/lib/os-release") {
			osRelease = "/usr/lib/os-release"
		}
		if osRelease != "" {
			data, _ := os.ReadFile(osRelease)
			content := strings.ToLower(string(data))
			switch {
			case strings.Contains(content, "antergos"):
				info.Distro = "Antergos"
				info.Release = "n/a"
			case strings.Contains(content, "logos"):
				info.Distro = "Logos"
				info.Release = "n/a"
			case strings.Contains(content, "swagarch"):
				info.Distro = "SwagArch"
				info.Release = "n/a"
			case strings.Contains(content, "obrevenge"):
				info.Distro = "OBRevenge"
				info.Release = "n/a"
			case strings.Contains(content, "alter"):
				info.Distro = "Alter Linux"
				info.Release = "n/a"
			}
		}
	case "ALDOS", "Aldos":
		info.Distro = "ALDOS"
	case "ArcoLinux":
		info.Distro = "ArcoLinux"
		info.Release = "n/a"
	case "artixlinux", "Artix Linux", "artix", "Artix", "Artix release":
		info.Distro = "Artix"
	case "blackPantherOS", "blackPanther", "blackpanther", "blackpantheros":
		info.Distro = readKeyValue("/etc/lsb-release", "DISTRIB_ID")
		info.Release = readKeyValue("/etc/lsb-release", "DISTRIB_RELEASE")
		info.Codename = readKeyValue("/etc/lsb-release", "DISTRIB_CODENAME")
	case "BLAG":
		info.Distro = "BLAG"
		if data, err := os.ReadFile("/etc/fedora-release"); err == nil {
			if lines := splitLines(string(data)); len(lines) > 0 {
				info.More = lines[0]
			}
		}
	case "Chakra":
		info.Distro = "Chakra"
		info.Release = ""
	case "BunsenLabs":
		info.Distro = readKeyValue("/etc/lsb-release", "DISTRIB_ID")
		info.Release = readKeyValue("/etc/lsb-release", "DISTRIB_RELEASE")
		info.Codename = readKeyValue("/etc/lsb-release", "DISTRIB_CODENAME")
	case "Debian":
		switch {
		case fileExists("/etc/crunchbang-lsb-release") || fileExists("/etc/lsb-release-crunchbang"):
			info.Distro = "CrunchBang"
			info.Release = readKeyValue("/etc/lsb-release-crunchbang", "DISTRIB_RELEASE")
			info.Codename = readKeyValue("/etc/lsb-release-crunchbang", "DISTRIB_DESCRIPTION")
		case fileExists("/etc/siduction-version"):
			info.Distro = "Siduction"
			info.Release = "(Debian Sid)"
			info.Codename = ""
		case fileExists("/usr/bin/pveversion"):
			info.Distro = "Proxmox VE"
			info.Codename = "n/a"
			info.Release = extractProxmoxVersion()
		case fileExists("/etc/os-release"):
			content, _ := os.ReadFile("/etc/os-release")
			lower := strings.ToLower(string(content))
			switch {
			case strings.Contains(lower, "raspbian"):
				info.Distro = "Raspbian"
				info.Release = readKeyValue("/etc/os-release", "PRETTY_NAME")
			case strings.Contains(lower, "blankon"):
				info.Distro = "BlankOn"
				info.Release = readKeyValue("/etc/os-release", "PRETTY_NAME")
			case strings.Contains(lower, "quirinux"):
				info.Distro = "Quirinux"
				info.Release = readKeyValue("/etc/os-release", "PRETTY_NAME")
			default:
				info.Distro = "Debian"
			}
		default:
			info.Distro = "Debian"
		}
	case "DraugerOS":
		info.Distro = "DraugerOS"
	case "elementary", "elementary OS":
		info.Distro = "elementary OS"
	case "EvolveOS":
		info.Distro = "Evolve OS"
	case "Sulin":
		info.Distro = "Sulin"
		info.Release = readKeyValue("/etc/os-release", "ID_LIKE")
		info.Codename = "Roolling donkey"
	case "KaOS", "kaos":
		info.Distro = "KaOS"
	case "frugalware":
		info.Distro = "Frugalware"
		info.Codename = ""
		info.Release = ""
	case "Fuduntu":
		info.Distro = "Fuduntu"
		info.Codename = ""
	case "Fux":
		info.Distro = "Fux"
		info.Codename = ""
	case "Gentoo":
		info.Distro = "Gentoo"
		if strings.Contains(runCommand("lsb_release", "-sd"), "Funtoo") {
			info.Distro = "Funtoo"
		}
		info.Release = detectGentooKeywords()
	case "Hyperbola GNU/Linux-libre", "Hyperbola":
		info.Distro = "Hyperbola GNU/Linux-libre"
		info.Codename = "n/a"
		info.Release = "n/a"
	case "januslinux", "janus":
		info.Distro = "januslinux"
	case "LinuxDeepin":
		info.Distro = "LinuxDeepin"
		info.Codename = ""
	case "Uos":
		info.Distro = "Uos"
		info.Codename = ""
	case "Kali", "Debian Kali Linux":
		info.Distro = "Kali Linux"
		if strings.Contains(info.Codename, "kali-rolling") {
			info.Codename = "n/a"
			info.Release = "n/a"
		}
	case "Lunar Linux", "lunar":
		info.Distro = "Lunar Linux"
	case "MandrivaLinux":
		info.Distro = "Mandriva"
		switch info.Codename {
		case "turtle", "Henry_Farman", "Farman", "Adelie", "pauillac":
			info.Distro = "Mandriva-" + info.Release
			info.Codename = ""
		}
	case "ManjaroLinux":
		info.Distro = "Manjaro"
	case "Mer":
		info.Distro = "Mer"
		if fileExists("/etc/os-release") {
			data, _ := os.ReadFile("/etc/os-release")
			if strings.Contains(string(data), "SailfishOS") {
				info.Distro = "SailfishOS"
				info.Release = readKeyValue("/etc/os-release", "VERSION")
				info.Codename = extractParen(info.Release)
			}
		}
	case "neon", "KDE neon":
		info.Distro = "KDE neon"
		info.Codename = "n/a"
		info.Release = "n/a"
		if data, err := os.ReadFile("/etc/issue"); err == nil {
			for _, line := range splitLines(string(data)) {
				if strings.HasPrefix(line, "KDE neon") {
					fields := strings.Fields(line)
					if len(fields) >= 3 {
						info.Release = fields[2]
					}
				}
			}
		}
	case "Ol", "ol", "Oracle Linux":
		info.Distro = "Oracle Linux"
		if data, err := os.ReadFile("/etc/oracle-release"); err == nil {
			info.Release = strings.TrimSpace(strings.Replace(string(data), "Oracle Linux ", "", 1))
		}
	case "LinuxMint":
		info.Distro = "Mint"
		if info.Codename == "debian" {
			info.Distro = "LMDE"
			info.Codename = "n/a"
			info.Release = "n/a"
		} else if strings.Contains(runCommand("lsb_release", "-sd"), "LMDE") {
			info.Distro = "LMDE"
		}
	case "openSUSE", "openSUSE project", "SUSE LINUX", "SUSE":
		info.Distro = "openSUSE"
		if fileExists("/etc/os-release") {
			data, _ := os.ReadFile("/etc/os-release")
			if strings.Contains(strings.ToLower(string(data)), "suse linux enterprise") {
				info.Distro = "SUSE Linux Enterprise"
				info.Codename = "n/a"
				info.Release = readKeyValue("/etc/os-release", "VERSION_ID")
			}
		}
		if info.Codename == "Tumbleweed" {
			info.Release = "n/a"
		}
	case "Parabola GNU/Linux-libre", "Parabola":
		info.Distro = "Parabola GNU/Linux-libre"
		info.Codename = "n/a"
		info.Release = "n/a"
	case "Parrot", "Parrot Security":
		info.Distro = "Parrot Security"
	case "PCLinuxOS":
		info.Distro = "PCLinuxOS"
		info.Codename = "n/a"
		info.Release = "n/a"
	case "Peppermint":
		info.Distro = "Peppermint"
		info.Codename = ""
	case "rhel":
		info.Distro = "Red Hat Enterprise Linux"
	case "RosaDesktopFresh":
		info.Distro = "ROSA"
		info.Release = readKeyValue("/etc/os-release", "VERSION")
		info.Codename = readKeyValue("/etc/os-release", "PRETTY_NAME")
	case "SailfishOS":
		info.Distro = "SailfishOS"
		if fileExists("/etc/os-release") {
			info.Release = readKeyValue("/etc/os-release", "VERSION")
			info.Codename = extractParen(info.Release)
		}
	case "Sparky", "SparkyLinux":
		info.Distro = "SparkyLinux"
	case "TeArch", "TeArchLinux", "":
		if detect != "" {
			info.Distro = "TeArch"
		} else {
			info.Distro = detect
		}
	case "Ubuntu":
		info.Distro = "Ubuntu"
	case "Viperr":
		info.Distro = "Viperr"
		info.Codename = ""
	case "Void", "VoidLinux":
		info.Distro = "Void Linux"
		info.Codename = ""
		info.Release = ""
	case "Zorin":
		info.Distro = "Zorin OS"
		info.Codename = ""
	default:
		info.Distro = detect
	}
}

func applyLSBPostFixups(info *DistroInfo) {
	d := info.Distro
	switch {
	case strings.Contains(d, "CentOSStream"):
		info.Distro = "CentOS Stream"
	case strings.Contains(d, "RedHatEnterprise"):
		info.Distro = "Red Hat Enterprise Linux"
	case strings.Contains(d, "Rocky Linux"):
		info.Distro = "Rocky Linux"
	case strings.Contains(d, "SUSELinuxEnterprise"):
		info.Distro = "SUSE Linux Enterprise"
	}
}

func detectFromUname(info *DistroInfo) {
	osType := runCommand("uname", "-o")
	switch osType {
	case "Cygwin", "FreeBSD", "OpenBSD", "NetBSD":
		info.Distro = osType
	case "DragonFly":
		info.Distro = "DragonFlyBSD"
	case "EndeavourOS":
		info.Distro = "EndeavourOS"
	case "Msys":
		info.Distro = "Msys"
		uname := runCommand("uname", "-r")
		if len(uname) > 0 {
			info.More = "Msys " + string(uname[0])
		}
	case "Haiku":
		info.Distro = "Haiku"
		if v := runCommand("uname", "-v"); v != "" {
			for _, f := range strings.Fields(v) {
				if strings.HasPrefix(f, "hrev") {
					info.More = f
					break
				}
			}
		}
	case "GNU/Linux":
		if hasCommand("crux") {
			info.Distro = "CRUX"
			if out := runCommand("crux"); out != "" {
				fields := strings.Fields(out)
				if len(fields) >= 3 {
					info.More = fields[2]
				}
			}
		}
		if hasCommand("nixos-version") {
			info.Distro = "NixOS"
			info.More = runCommand("nixos-version")
		}
		if hasCommand("sorcery") {
			info.Distro = "SMGL"
		}
		if hasCommand("guix") && hasCommand("herd") {
			info.Distro = "Guix System"
		}
	}
}

func detectFromOSRelease(info *DistroInfo) {
	path := ""
	if fileExists("/etc/os-release") {
		path = "/etc/os-release"
	} else if fileExists("/usr/lib/os-release") {
		path = "/usr/lib/os-release"
	}
	if path == "" {
		return
	}
	id := readKeyValue(path, "ID")
	if id == "" {
		return
	}

	// Title-case each space-separated token (shell does the same with ${i^})
	parts := strings.Fields(id)
	for i, p := range parts {
		if p == "" {
			continue
		}
		parts[i] = strings.ToUpper(p[:1]) + p[1:]
	}
	info.Distro = strings.Join(parts, " ")

	// Hotfixes straight out of the shell script
	switch info.Distro {
	case "Opensuse-tumbleweed":
		info.Distro = "openSUSE"
		info.More = "Tumbleweed"
	case "Opensuse-leap":
		info.Distro = "openSUSE"
	case "Void":
		info.Distro = "Void Linux"
	case "Evolveos":
		info.Distro = "Evolve OS"
	case "Antergos":
		info.Distro = "Antergos"
	case "Logos":
		info.Distro = "Logos"
	case "Alter":
		info.Distro = "Alter Linux"
	case "Arch", "Archarm":
		info.Distro = "Arch Linux"
	case "Elementary":
		info.Distro = "elementary OS"
	case "Fedora":
		if fileExists("/etc/qubes-rpc") {
			info.Distro = "qubes"
		}
	case "Ol":
		info.Distro = "Oracle Linux"
	case "Neon":
		info.Distro = "KDE neon"
	case "Sled", "Sles":
		info.Distro = "SUSE Linux Enterprise"
	}
	if info.Distro == "Oracle Linux" && fileExists("/etc/oracle-release") {
		data, _ := os.ReadFile("/etc/oracle-release")
		info.More = strings.TrimSpace(strings.Replace(string(data), "Oracle Linux ", "", 1))
	}
	if info.Distro == "Rhel" {
		info.Distro = "Red Hat Enterprise Linux"
		if data, err := os.ReadFile("/etc/os-release"); err == nil {
			content := string(data)
			switch {
			case strings.Contains(content, "Scientific"):
				info.Distro = "Scientific Linux"
			case strings.Contains(content, "EuroLinux"):
				info.Distro = "EuroLinux"
			}
		}
	}
	if info.Distro == "SUSE Linux Enterprise" && fileExists("/etc/os-release") {
		info.More = readKeyValue("/etc/os-release", "VERSION_ID")
	}
	if info.Distro == "Debian" && fileExists("/usr/bin/pveversion") {
		info.Distro = "Proxmox VE"
		info.Codename = "n/a"
		info.Release = extractProxmoxVersion()
	}
	if info.Distro == "Almalinux" && fileExists("/etc/almalinux-release") {
		info.Distro = "AlmaLinux"
		data, _ := os.ReadFile("/etc/almalinux-release")
		s := strings.TrimSpace(string(data))
		s = strings.Replace(s, "AlmaLinux release ", "", 1)
		fields := strings.Fields(s)
		if len(fields) > 0 {
			info.Release = fields[0]
		}
		if open := strings.Index(s, "("); open >= 0 {
			if close := strings.Index(s[open:], ")"); close >= 0 {
				info.Codename = s[open+1 : open+close]
			}
		}
	}
}

func detectFromReleaseFiles(info *DistroInfo) {
	candidates := []string{
		"arch", "chakra", "crunchbang-lsb", "evolveos", "exherbo", "fedora",
		"frugalware", "fux", "gentoo", "kogaion", "mageia", "obarun", "oracle",
		"pardus", "pclinuxos", "redhat", "redstar", "rosa", "SuSe",
	}
	for _, di := range candidates {
		if fileExists("/etc/" + di + "-release") {
			info.Distro = di
			break
		}
	}
	switch info.Distro {
	case "crunchbang-lsb":
		info.Distro = "Crunchbang"
	case "gentoo":
		if containsInFile("/etc/gentoo-release", "Funtoo") {
			info.Distro = "Funtoo"
		}
	case "mandrake", "mandriva":
		if containsInFile("/etc/"+info.Distro+"-release", "PCLinuxOS") {
			info.Distro = "PCLinuxOS"
		}
	case "fedora":
		if containsInFile("/etc/fedora-release", "Korora") {
			info.Distro = "Korora"
		}
		if containsInFile("/etc/fedora-release", "BLAG") {
			info.Distro = "BLAG"
			if data, err := os.ReadFile("/etc/fedora-release"); err == nil {
				if lines := splitLines(string(data)); len(lines) > 0 {
					info.More = lines[0]
				}
			}
		}
	case "oracle":
		if data, err := os.ReadFile("/etc/oracle-release"); err == nil {
			info.More = strings.TrimSpace(strings.Replace(string(data), "Oracle Linux ", "", 1))
		}
	case "SuSe":
		info.Distro = "openSUSE"
		if fileExists("/etc/os-release") {
			data, _ := os.ReadFile("/etc/os-release")
			if strings.Contains(strings.ToLower(string(data)), "suse linux enterprise") {
				info.Distro = "SUSE Linux Enterprise"
				info.More = readKeyValue("/etc/os-release", "VERSION_ID")
			}
		}
		if strings.Contains(info.More, "Tumbleweed") {
			info.More = "Tumbleweed"
		}
	case "redstar":
		if data, err := os.ReadFile("/etc/redstar-release"); err == nil {
			info.More = extractDigitsAndDots(string(data))
		}
	case "redhat":
		data, _ := os.ReadFile("/etc/redhat-release")
		content := string(data)
		switch {
		case strings.Contains(content, "CentOS"):
			info.Distro = "CentOS"
		case strings.Contains(content, "Rocky Linux"):
			info.Distro = "Rocky Linux"
		case strings.Contains(strings.ToLower(content), "almalinux"):
			info.Distro = "AlmaLinux"
		case strings.Contains(content, "Scientific"):
			info.Distro = "Scientific Linux"
		case strings.Contains(content, "EuroLinux"):
			info.Distro = "EuroLinux"
		case strings.Contains(content, "PCLinuxOS"):
			info.Distro = "PCLinuxOS"
		}
	}
}

func detectFromMiscFiles(info *DistroInfo) {
	if isLinuxOrGNU() {
		switch {
		case fileExists("/etc/debian_version"):
			if data, err := os.ReadFile("/etc/issue"); err == nil {
				issue := string(data)
				switch {
				case strings.Contains(strings.ToLower(issue), "gnewsense"):
					info.Distro = "gNewSense"
				case strings.Contains(issue, "KDE neon"):
					info.Distro = "KDE neon"
					if fields := strings.Fields(issue); len(fields) >= 3 {
						info.More = fields[2]
					}
				default:
					info.Distro = "Debian"
				}
			} else {
				info.Distro = "Debian"
			}
			if containsInFile("/etc/debian_version", "Kali") {
				info.Distro = "Kali Linux"
			}
		case fileExists("/etc/NIXOS"):
			info.Distro = "NixOS"
		case fileExists("/etc/dragora-version"):
			info.Distro = "Dragora"
			if data, err := os.ReadFile("/etc/dragora-version"); err == nil {
				parts := strings.SplitN(string(data), ",", 2)
				info.More = strings.TrimSpace(parts[0])
			}
		case fileExists("/etc/slackware-version"):
			info.Distro = "Slackware"
		case fileExists("/usr/share/doc/tc/release.txt"):
			info.Distro = "TinyCore"
			if data, err := os.ReadFile("/usr/share/doc/tc/release.txt"); err == nil {
				info.More = strings.TrimSpace(string(data))
			}
		case fileExists("/etc/sabayon-edition"):
			info.Distro = "Sabayon"
		}
		return
	}

	// Non-Linux branches
	if hasCommand("/usr/bin/sw_vers") {
		out := runCommand("/usr/bin/sw_vers")
		lower := strings.ToLower(out)
		switch {
		case strings.Contains(lower, "mac os x"):
			info.Distro = "Mac OS X"
		case strings.Contains(lower, "macos"):
			info.Distro = "macOS"
		}
	}
	if info.Distro == "Unknown" && fileExists("/var/run/dmesg.boot") {
		data, _ := os.ReadFile("/var/run/dmesg.boot")
		for _, line := range splitLines(string(data)) {
			switch {
			case strings.Contains(line, "DragonFly"):
				info.Distro = "DragonFlyBSD"
				return
			case strings.Contains(line, "FreeBSD"):
				info.Distro = "FreeBSD"
				return
			case strings.Contains(line, "NetBSD"):
				info.Distro = "NetBSD"
				return
			case strings.Contains(line, "OpenBSD"):
				info.Distro = "OpenBSD"
				return
			}
		}
	}
}

func detectFromIssue(info *DistroInfo) {
	data, err := os.ReadFile("/etc/issue")
	if err != nil {
		return
	}
	content := string(data)
	switch {
	case strings.Contains(content, "Hyperbola GNU/Linux-libre"):
		info.Distro = "Hyperbola GNU/Linux-libre"
	case strings.Contains(content, "LinuxDeepin"):
		info.Distro = "LinuxDeepin"
	case strings.Contains(content, "Obarun"):
		info.Distro = "Obarun"
	case strings.Contains(content, "Parabola GNU/Linux-libre"):
		info.Distro = "Parabola GNU/Linux-libre"
	case strings.Contains(content, "Solus"):
		info.Distro = "Solus"
	case strings.Contains(content, "ALDOS"):
		info.Distro = "ALDOS"
	}
}

func detectFromSystemRelease(info *DistroInfo) {
	switch {
	case fileExists("/etc/system-release"):
		if containsInFile("/etc/system-release", "Scientific Linux") {
			info.Distro = "Scientific Linux"
		} else if containsInFile("/etc/system-release", "Oracle Linux") {
			info.Distro = "Oracle Linux"
		}
	case fileExists("/etc/lsb-release"):
		if name := readKeyValue("/etc/lsb-release", "CHROMEOS_RELEASE_NAME"); name != "" {
			info.Distro = name
			info.More = readKeyValue("/etc/lsb-release", "CHROMEOS_RELEASE_VERSION")
		}
	}
}


func extractProxmoxVersion() string {
	out := runCommand("/usr/bin/pveversion")
	idx := strings.Index(out, "pve-manager/")
	if idx < 0 {
		return ""
	}
	rest := out[idx+len("pve-manager/"):]
	end := 0
	seenDot := false
	for end < len(rest) {
		c := rest[end]
		if c >= '0' && c <= '9' {
			end++
			continue
		}
		if c == '.' && !seenDot {
			seenDot = true
			end++
			continue
		}
		break
	}
	return rest[:end]
}

func detectGentooKeywords() string {
	files := []string{"/etc/portage/make.conf"}
	if info, err := os.Stat("/etc/portage/make.conf"); err == nil && info.IsDir() {
		entries, _ := os.ReadDir("/etc/portage/make.conf")
		files = files[:0]
		for _, e := range entries {
			files = append(files, "/etc/portage/make.conf/"+e.Name())
		}
	}
	keywords := ""
	for _, f := range files {
		if v := readKeyValue(f, "ACCEPT_KEYWORDS"); v != "" {
			keywords = v
		}
	}
	switch {
	case strings.HasPrefix(keywords, "~"):
		return "testing"
	case keywords == "**":
		return "experimental"
	case keywords != "":
		return "stable"
	}
	return ""
}

func extractParen(s string) string {
	open := strings.Index(s, "(")
	if open < 0 {
		return ""
	}
	close := strings.Index(s[open:], ")")
	if close < 0 {
		return ""
	}
	return s[open+1 : open+close]
}

func extractDigitsAndDots(s string) string {
	var b strings.Builder
	for _, r := range s {
		if (r >= '0' && r <= '9') || r == '.' {
			b.WriteRune(r)
		}
	}
	return b.String()
}

// canonicalizeDistro mirrors the final case statement in the shell script
// that maps lowercase identifiers to their canonical display names.
func canonicalizeDistro(d string) string {
	switch {
	case d == "aldos":
		return "ALDOS"
	case d == "alpine":
		return "Alpine Linux"
	case d == "almalinux":
		return "AlmaLinux"
	case matches(d, "alter*linux") || d == "alter":
		return "Alter Linux"
	case d == "amzn" || d == "amazon" || matches(d, "amazon*linux"):
		return "Amazon Linux"
	case d == "antergos":
		return "Antergos"
	case matches(d, "arch*linux*old"):
		return "Arch Linux - Old"
	case d == "arch" || matches(d, "arch*linux"):
		return "Arch Linux"
	case d == "arch32":
		return "Arch Linux 32"
	case d == "arcolinux" || strings.HasPrefix(d, "arcolinux"):
		return "ArcoLinux"
	case d == "artix" || matches(d, "artix*linux"):
		return "Artix Linux"
	case d == "blackpantheros" || matches(d, "black*panther*"):
		return "blackPanther OS"
	case d == "blag":
		return "BLAG"
	case d == "bunsenlabs":
		return "BunsenLabs"
	case d == "centos":
		return "CentOS"
	case matches(d, "centos*stream"):
		return "CentOS Stream"
	case d == "chakra":
		return "Chakra"
	case d == "chapeau":
		return "Chapeau"
	case strings.HasPrefix(d, "chrome") || strings.HasPrefix(d, "chromium"):
		return "Chrome OS"
	case d == "crunchbang":
		return "CrunchBang"
	case d == "crux":
		return "CRUX"
	case d == "cygwin":
		return "Cygwin"
	case d == "debian":
		return "Debian"
	case d == "devuan":
		return "Devuan"
	case d == "deepin":
		return "Deepin"
	case d == "uos":
		return "Uos"
	case d == "desaos":
		return "DesaOS"
	case d == "dragonflybsd":
		return "DragonFlyBSD"
	case d == "dragora":
		return "Dragora"
	case strings.HasPrefix(d, "drauger"):
		return "DraugerOS"
	case d == "elementary" || d == "elementary os":
		return "elementary OS"
	case d == "eurolinux":
		return "EuroLinux"
	case d == "evolveos":
		return "Evolve OS"
	case d == "sulin":
		return "Sulin"
	case d == "exherbo" || matches(d, "exherbo*linux"):
		return "Exherbo"
	case d == "fedora":
		return "Fedora"
	case matches(d, "fedora*old"):
		return "Fedora - Old"
	case d == "freebsd":
		return "FreeBSD"
	case matches(d, "freebsd*old"):
		return "FreeBSD - Old"
	case d == "frugalware":
		return "Frugalware"
	case d == "fuduntu":
		return "Fuduntu"
	case d == "funtoo":
		return "Funtoo"
	case d == "fux":
		return "Fux"
	case d == "gentoo":
		return "Gentoo"
	case d == "gnewsense":
		return "gNewSense"
	case matches(d, "guix*system"):
		return "Guix System"
	case d == "haiku":
		return "Haiku"
	case d == "hyperbolagnu" || d == "hyperbolagnu/linux-libre" || d == "hyperbola gnu/linux-libre" || d == "hyperbola":
		return "Hyperbola GNU/Linux-libre"
	case d == "januslinux":
		return "januslinux"
	case matches(d, "kali*linux"):
		return "Kali Linux"
	case d == "kaos":
		return "KaOS"
	case matches(d, "kde*neon") || d == "neon":
		return "KDE neon"
	case d == "kogaion":
		return "Kogaion"
	case d == "korora":
		return "Korora"
	case d == "linuxdeepin":
		return "LinuxDeepin"
	case d == "lmde":
		return "LMDE"
	case d == "logos":
		return "Logos"
	case d == "lunar" || matches(d, "lunar*linux"):
		return "Lunar Linux"
	case matches(d, "mac*os*x") || matches(d, "os*x"):
		return "Mac OS X"
	case d == "macos":
		return "macOS"
	case d == "manjaro":
		return "Manjaro"
	case d == "mageia":
		return "Mageia"
	case d == "mandrake":
		return "Mandrake"
	case d == "mandriva":
		return "Mandriva"
	case d == "mer":
		return "Mer"
	case d == "mint" || matches(d, "linux*mint"):
		return "Mint"
	case d == "msys" || d == "msys2":
		return "Msys"
	case d == "netbsd":
		return "NetBSD"
	case d == "netrunner":
		return "Netrunner"
	case d == "nix" || matches(d, "nix*os"):
		return "NixOS"
	case d == "obarun":
		return "Obarun"
	case d == "obrevenge":
		return "OBRevenge"
	case d == "ol" || matches(d, "oracle*linux"):
		return "Oracle Linux"
	case d == "openbsd":
		return "OpenBSD"
	case d == "opensuse":
		return "openSUSE"
	case matches(d, "os*elbrus"):
		return "OS Elbrus"
	case d == "parabolagnu" || d == "parabolagnu/linux-libre" || d == "parabola gnu/linux-libre" || d == "parabola":
		return "Parabola GNU/Linux-libre"
	case d == "pardus":
		return "Pardus"
	case d == "parrot" || matches(d, "parrot*security"):
		return "Parrot Security"
	case d == "pclinuxos" || d == "pclos":
		return "PCLinuxOS"
	case d == "peppermint":
		return "Peppermint"
	case d == "proxmox" || matches(d, "proxmox*ve"):
		return "Proxmox VE"
	case d == "pureos":
		return "PureOS"
	case d == "quirinux":
		return "Quirinux"
	case d == "qubes":
		return "Qubes OS"
	case d == "raspbian":
		return "Raspbian"
	case matches(d, "red*hat*") || d == "rhel":
		return "Red Hat Enterprise Linux"
	case d == "rosa":
		return "ROSA"
	case matches(d, "red*star") || matches(d, "red*star*os"):
		return "Red Star OS"
	case d == "rocky":
		return "Rocky Linux"
	case d == "sabayon":
		return "Sabayon"
	case d == "sailfish" || matches(d, "sailfish*os"):
		return "SailfishOS"
	case strings.HasPrefix(d, "scientific"):
		return "Scientific Linux"
	case d == "siduction":
		return "Siduction"
	case d == "slackware":
		return "Slackware"
	case d == "smgl" || matches(d, "source*mage") || matches(d, "source*mage*gnu*linux"):
		return "Source Mage GNU/Linux"
	case d == "solus":
		return "Solus"
	case d == "sparky" || matches(d, "sparky*linux"):
		return "SparkyLinux"
	case d == "steam" || matches(d, "steam*os"):
		return "SteamOS"
	case matches(d, "suse*linux*enterprise"):
		return "SUSE Linux Enterprise"
	case d == "swagarch":
		return "SwagArch"
	case strings.HasPrefix(d, "tearch"):
		return "TeArch"
	case d == "tinycore" || matches(d, "tinycore*linux"):
		return "TinyCore"
	case d == "trisquel":
		return "Trisquel"
	case d == "grombyangos":
		return "GrombyangOS"
	case d == "ubuntu":
		return "Ubuntu"
	case d == "viperr":
		return "Viperr"
	case matches(d, "void*linux"):
		return "Void Linux"
	case strings.HasPrefix(d, "zorin"):
		return "Zorin OS"
	case strings.HasPrefix(d, "endeavour"):
		return "EndeavourOS"
	}
	// Unknown stays Unknown (title-cased); otherwise return as provided.
	if d == "" {
		return "Unknown"
	}
	return strings.ToUpper(d[:1]) + d[1:]
}

// matches implements a minimal glob matcher supporting only the '*' wildcard,
// which is all that the canonicalization table needs.
func matches(s, pattern string) bool {
	parts := strings.Split(pattern, "*")
	if len(parts) == 1 {
		return s == pattern
	}
	// Anchor first and last segments
	if !strings.HasPrefix(s, parts[0]) {
		return false
	}
	if !strings.HasSuffix(s, parts[len(parts)-1]) {
		return false
	}
	pos := len(parts[0])
	for _, p := range parts[1 : len(parts)-1] {
		idx := strings.Index(s[pos:], p)
		if idx < 0 {
			return false
		}
		pos += idx + len(p)
	}
	return true
}
