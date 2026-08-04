# Sagaz — PyInstaller Cleanup

Removes leftover temp files created by PyInstaller executables on Windows.

**What it deletes:**
- `%LOCALAPPDATA%\Temp\_MEI*` — PyInstaller extraction folders
- `%LOCALAPPDATA%\Packages\...\Temp\mat-debug-*.log` — debug logs

## Use

Download `cleanup.exe` and double-click. No install needed.

## Build

```bash
GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -o cleanup.exe main.go
```

## License

MIT
