package components

type DisplayFields uint16

// was not initially a fan of this but i wanted to explore. at the same time
// i should say that I am considering of migrating it, no clue if this is better
const (
	DisplayDistro   DisplayFields = 1 << iota
	DisplayHost
	DisplayKernel
	DisplayUptime
	DisplayPkgs
	DisplayShell
	DisplayRes
	DisplayDE
	DisplayWM
	DisplayWMTheme
	DisplayGTK
	DisplayDisk
	DisplayCPU
	DisplayGPU
	DisplayMem
)

type DisplayConfig struct {
	Fields DisplayFields
}

func (d *DisplayConfig) Has(field DisplayFields) bool {
	return d.Fields&field != 0
}