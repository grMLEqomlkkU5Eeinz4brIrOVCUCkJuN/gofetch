package components

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"
	"syscall"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

// SystemInfo holds all gathered system information.
type SystemInfo struct {
	User       string
	Hostname   string
	Distro     *DistroInfo
	Kernel     string
	Uptime     string
	Packages   string
	Shell      string
	Resolution string
	DE         string
	WM         string
	CPU        string
	GPU        string
	Memory     string
	Disk       string
}

// GatherSystemInfo collects all system information.
func GatherSystemInfo() *SystemInfo {
	info := &SystemInfo{}

	// User@Host
	if u, err := user.Current(); err == nil {
		info.User = u.Username
	}
	if h, err := os.Hostname(); err == nil {
		info.Hostname = h
	}

	// Distro
	info.Distro = DetectDistroInfo()

	// Kernel
	info.Kernel = detectKernel()

	// Uptime
	info.Uptime = detectUptime()

	// Packages
	info.Packages = detectPackages()

	// Shell
	info.Shell = detectShell()

	// Resolution
	info.Resolution = detectResolution()

	// DE
	info.DE = detectDE()

	// WM
	if wm, err := DetectWindowManager(windowManagerRegistry); err == nil {
		info.WM = wm
	}

	// CPU
	info.CPU = detectCPU()

	// GPU
	info.GPU = detectGPU()

	// Memory
	info.Memory = detectMemory()

	// Disk
	info.Disk = detectDisk()

	return info
}

// InfoLines returns the formatted info lines for display.
func (s *SystemInfo) InfoLines() []string {
	titleColor := "\033[1;34m"
	reset := "\033[0m"

	title := fmt.Sprintf("%s%s@%s%s", titleColor, s.User, s.Hostname, reset)
	separator := strings.Repeat("-", len(s.User)+1+len(s.Hostname))

	lines := []string{title, separator}

	add := func(label, value string) {
		if value != "" {
			lines = append(lines, fmt.Sprintf("%s%s:%s %s", titleColor, label, reset, value))
		}
	}

	distroStr := "Unknown"
	if s.Distro != nil {
		if s.Distro.More != "" {
			distroStr = s.Distro.More
		} else {
			distroStr = s.Distro.Distro
		}
	}

	add("OS", distroStr)
	add("Kernel", s.Kernel)
	add("Uptime", s.Uptime)
	add("Packages", s.Packages)
	add("Shell", s.Shell)
	add("Resolution", s.Resolution)
	add("DE", s.DE)
	add("WM", s.WM)
	add("CPU", s.CPU)
	add("GPU", s.GPU)
	add("Memory", s.Memory)
	add("Disk", s.Disk)

	// Color bar
	lines = append(lines, "")
	bar1 := ""
	bar2 := ""
	for i := 0; i < 8; i++ {
		bar1 += fmt.Sprintf("\033[4%dm   ", i)
		bar2 += fmt.Sprintf("\033[10%dm   ", i)
	}
	lines = append(lines, bar1+"\033[0m")
	lines = append(lines, bar2+"\033[0m")

	return lines
}

func detectKernel() string {
	var utsname syscall.Utsname
	if err := syscall.Uname(&utsname); err != nil {
		return ""
	}
	return charsToString(utsname.Release[:])
}

func charsToString(ca []int8) string {
	buf := make([]byte, 0, len(ca))
	for _, c := range ca {
		if c == 0 {
			break
		}
		buf = append(buf, byte(c))
	}
	return string(buf)
}

func detectUptime() string {
	uptime, err := host.Uptime()
	if err != nil {
		return ""
	}
	days := uptime / 86400
	hours := (uptime % 86400) / 3600
	mins := (uptime % 3600) / 60

	parts := []string{}
	if days > 0 {
		parts = append(parts, fmt.Sprintf("%dd", days))
	}
	if hours > 0 {
		parts = append(parts, fmt.Sprintf("%dh", hours))
	}
	parts = append(parts, fmt.Sprintf("%dm", mins))
	return strings.Join(parts, " ")
}

func detectPackages() string {
	var counts []string

	// pacman (Arch)
	if entries, err := os.ReadDir("/var/lib/pacman/local"); err == nil {
		// subtract 1 for ALPM_DB_VERSION file
		count := len(entries) - 1
		if count > 0 {
			counts = append(counts, fmt.Sprintf("%d (pacman)", count))
		}
	}

	// dpkg (Debian/Ubuntu)
	if matches, err := filepath.Glob("/var/lib/dpkg/info/*.list"); err == nil && len(matches) > 0 {
		counts = append(counts, fmt.Sprintf("%d (dpkg)", len(matches)))
	}

	// rpm
	if runtime.GOOS == "linux" {
		out := runCommand("rpm", "-qa", "--last")
		if out != "" {
			lines := splitLines(out)
			if len(lines) > 0 {
				counts = append(counts, fmt.Sprintf("%d (rpm)", len(lines)))
			}
		}
	}

	// flatpak
	out := runCommand("flatpak", "list", "--app")
	if out != "" {
		lines := splitLines(out)
		if len(lines) > 0 {
			counts = append(counts, fmt.Sprintf("%d (flatpak)", len(lines)))
		}
	}

	// snap
	out = runCommand("snap", "list")
	if out != "" {
		lines := splitLines(out)
		if len(lines) > 1 { // first line is header
			counts = append(counts, fmt.Sprintf("%d (snap)", len(lines)-1))
		}
	}

	return strings.Join(counts, ", ")
}

func detectShell() string {
	shell := os.Getenv("SHELL")
	if shell == "" {
		return ""
	}
	base := filepath.Base(shell)
	version := runCommand(shell, "--version")
	if version != "" {
		lines := splitLines(version)
		if len(lines) > 0 {
			first := lines[0]
			for _, field := range strings.Fields(first) {
				if len(field) > 0 && field[0] >= '0' && field[0] <= '9' {
					// Extract just major.minor.patch
					ver := field
					if idx := strings.Index(ver, "("); idx >= 0 {
						ver = ver[:idx]
					}
					return base + " " + ver
				}
			}
		}
	}
	return base
}

func detectResolution() string {
	out := runCommand("xrandr", "--current")
	if out == "" {
		// Try wayland via wlr-randr
		out = runCommand("wlr-randr")
	}
	if out == "" {
		return ""
	}
	var resolutions []string
	for _, line := range splitLines(out) {
		if strings.Contains(line, "*") || strings.Contains(line, "current") {
			fields := strings.Fields(line)
			for i, f := range fields {
				if f == "current" && i+2 < len(fields) {
					resolutions = append(resolutions, fields[i+1]+" x "+fields[i+3])
				}
			}
			// xrandr connected line with resolution
			if strings.Contains(line, " connected") {
				for _, f := range fields {
					if strings.Contains(f, "x") && strings.Contains(f, "+") {
						res := strings.Split(f, "+")[0]
						resolutions = append(resolutions, res)
					}
				}
			}
		}
	}
	if len(resolutions) > 0 {
		return strings.Join(resolutions, ", ")
	}
	return ""
}

func detectDE() string {
	de := os.Getenv("XDG_CURRENT_DESKTOP")
	if de != "" {
		return de
	}
	de = os.Getenv("DESKTOP_SESSION")
	if de != "" {
		return de
	}
	return ""
}

func detectCPU() string {
	infos, err := cpu.Info()
	if err != nil || len(infos) == 0 {
		return ""
	}
	name := strings.TrimSpace(infos[0].ModelName)
	// Clean up common fluff
	name = strings.ReplaceAll(name, "(R)", "")
	name = strings.ReplaceAll(name, "(TM)", "")
	name = strings.ReplaceAll(name, "CPU ", "")
	// Collapse multiple spaces
	fields := strings.Fields(name)
	return strings.Join(fields, " ")
}

func detectGPU() string {
	out := runCommand("lspci", "-mm")
	if out == "" {
		return ""
	}
	var gpus []string
	for _, line := range splitLines(out) {
		lower := strings.ToLower(line)
		if strings.Contains(lower, "vga") || strings.Contains(lower, "3d") || strings.Contains(lower, "display") {
			// Parse the quoted fields
			fields := parseQuotedFields(line)
			if len(fields) >= 4 {
				gpu := strings.TrimSpace(fields[3])
				if len(fields) >= 5 {
					gpu += " " + strings.TrimSpace(fields[4])
				}
				gpus = append(gpus, gpu)
			}
		}
	}
	return strings.Join(gpus, ", ")
}

func parseQuotedFields(s string) []string {
	var fields []string
	for len(s) > 0 {
		s = strings.TrimLeft(s, " \t")
		if len(s) == 0 {
			break
		}
		if s[0] == '"' {
			end := strings.Index(s[1:], "\"")
			if end < 0 {
				fields = append(fields, s[1:])
				break
			}
			fields = append(fields, s[1:end+1])
			s = s[end+2:]
		} else {
			end := strings.IndexAny(s, " \t")
			if end < 0 {
				fields = append(fields, s)
				break
			}
			fields = append(fields, s[:end])
			s = s[end:]
		}
	}
	return fields
}

func detectMemory() string {
	v, err := mem.VirtualMemory()
	if err != nil {
		return ""
	}
	usedMiB := (v.Total - v.Available) / 1024 / 1024
	totalMiB := v.Total / 1024 / 1024
	return fmt.Sprintf("%dMiB / %dMiB", usedMiB, totalMiB)
}

func detectDisk() string {
	usage, err := disk.Usage("/")
	if err != nil {
		return ""
	}
	usedGiB := float64(usage.Used) / 1024 / 1024 / 1024
	totalGiB := float64(usage.Total) / 1024 / 1024 / 1024
	return fmt.Sprintf("%.1fGiB / %.1fGiB (%.0f%%)", usedGiB, totalGiB, usage.UsedPercent)
}
