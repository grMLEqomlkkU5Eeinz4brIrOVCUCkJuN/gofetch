package main

import (
	"github.com/grMLEqomlkkU5Eeinz4brIrOVCUCkJuN/gofetch/src/components"
)

func main() {
	info := components.GatherSystemInfo()
	components.Display(info)
}
