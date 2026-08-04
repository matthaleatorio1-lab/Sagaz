package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	local := os.Getenv("LOCALAPPDATA")
	removed := 0

	// Remove _MEI* folders left by PyInstaller
	entries, _ := os.ReadDir(filepath.Join(local, "Temp"))
	for _, e := range entries {
		if e.IsDir() && strings.HasPrefix(e.Name(), "_MEI") {
			path := filepath.Join(local, "Temp", e.Name())
			if os.RemoveAll(path) == nil {
				fmt.Println("Removed:", path)
				removed++
			}
		}
	}

	// Remove mat-debug-*.log files
	logs, _ := filepath.Glob(filepath.Join(local,
		`Packages\microsoft.windowscommunicationsapps_8wekyb3d8bbwe\AC\Temp\mat-debug-*.log`))
	for _, f := range logs {
		if os.Remove(f) == nil {
			fmt.Println("Removed:", f)
			removed++
		}
	}

	fmt.Printf("\nDone. %d item(s) removed.\n", removed)
	fmt.Print("Press Enter to exit...")
	os.Stdin.Read(make([]byte, 1))
}
