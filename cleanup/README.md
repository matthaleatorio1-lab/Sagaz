# PyInstaller Cleanup Utility

A small Windows utility that removes leftover temp files created by PyInstaller-bundled executables.

## What it cleans

| Location | Pattern | Why it exists |
|---|---|---|
| `%LOCALAPPDATA%\Temp\` | `_MEI*` folders | PyInstaller extracts its bundle here on every run and doesn't always clean up |
| `%LOCALAPPDATA%\Packages\microsoft.windowscommunicationsapps_8wekyb3d8bbwe\AC\Temp\` | `mat-debug-*.log` | Debug logs triggered by certain PyInstaller apps |

## Usage

1. Download `cleanup.exe` from [Releases](../../releases)
2. Double-click it — no installation needed
3. It lists every item removed, then waits for you to press Enter

No admin rights required (everything is inside your own user profile).

## Building from source

Requires [Go 1.21+](https://go.dev/dl/).

```bash
# Windows
go build -o cleanup.exe main.go

# Cross-compile from Linux/macOS
GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -o cleanup.exe main.go
```

## Contributing

Pull requests welcome! Some ideas for improvement:

- [ ] Add a `--dry-run` flag to preview what would be deleted
- [ ] Scan for other common PyInstaller artefacts
- [ ] Add a tray icon / silent mode for scheduled-task use
- [ ] Localisation / translated output

## License

MIT
