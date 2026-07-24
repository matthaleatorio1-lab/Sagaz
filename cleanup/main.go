package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	localAppData := os.Getenv("LOCALAPPDATA")
	if localAppData == "" {
		fmt.Println("[ERROR] LOCALAPPDATA environment variable not found. Are you running this on Windows?")
		pause()
		return
	}

	var totalRemoved int
	var totalErrors int

	fmt.Println("=== PyInstaller Cleanup Utility ===")
	fmt.Println()

	// 1. Remove _MEI* folders in %LOCALAPPDATA%\Temp
	tempDir := filepath.Join(localAppData, "Temp")
	fmt.Printf("Scanning: %s\\_MEI*\n", tempDir)

	entries, err := os.ReadDir(tempDir)
	if err != nil {
		fmt.Printf("[WARNING] Could not read temp dir: %v\n", err)
	} else {
		for _, entry := range entries {
			if entry.IsDir() && strings.HasPrefix(entry.Name(), "_MEI") {
				fullPath := filepath.Join(tempDir, entry.Name())
				err := os.RemoveAll(fullPath)
				if err != nil {
					fmt.Printf("  [FAIL] %s — %v\n", entry.Name(), err)
					totalErrors++
				} else {
					fmt.Printf("  [OK]   Removed %s\n", entry.Name())
					totalRemoved++
				}
			}
		}
	}

	// 2. Remove mat-debug-*.log files in WindowsCommunicationsApps temp
	mailTemp := filepath.Join(
		localAppData,
		"Packages",
		"microsoft.windowscommunicationsapps_8wekyb3d8bbwe",
		"AC", "Temp",
	)
	fmt.Printf("\nScanning: %s\\mat-debug-*.log\n", mailTemp)

	logMatches, err := filepath.Glob(filepath.Join(mailTemp, "mat-debug-*.log"))
	if err != nil {
		fmt.Printf("[WARNING] Could not scan mail temp dir: %v\n", err)
	} else {
		for _, logFile := range logMatches {
			err := os.Remove(logFile)
			if err != nil {
				fmt.Printf("  [FAIL] %s — %v\n", filepath.Base(logFile), err)
				totalErrors++
			} else {
				fmt.Printf("  [OK]   Removed %s\n", filepath.Base(logFile))
				totalRemoved++
			}
		}
	}

	// Summary
	fmt.Println()
	fmt.Printf("Done. Removed: %d item(s)  |  Errors: %d\n", totalRemoved, totalErrors)
	pause()
}

func pause() {
	fmt.Print("\nPress Enter to exit...")
	buf := make([]byte, 1)
	os.Stdin.Read(buf)
}
