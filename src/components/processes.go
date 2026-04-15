package components

import (
	"fmt"
	"regexp"
	"github.com/shirou/gopsutil/process"
)

// isProcessRunning checks if any running process matches the given regex pattern
func isProcessRunning(nameRegex string) (bool, error) {
	// compile the regex
	re, err := regexp.Compile(nameRegex)
	if err != nil {
		return false, fmt.Errorf("invalid regex pattern %q: %w", nameRegex, err)
	}

	// get all the processes
	processes, err := process.Processes()
	if err != nil {
		return false, fmt.Errorf("failed to get processes: %w", err)
	}

	// iterate and search
	for _, p := range processes {
		name, err := p.Name()
		if err != nil {
			// Some processes might not be accessible, skip them
			continue
		}

		if re.MatchString(name) {
			return true, nil
		}
	}

	return false, nil
}
