package main

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
)

//go:embed AutoGo.exe
var autoGoExe []byte

func main() {
	goPath := os.Getenv("GOPATH")
	if goPath == "" {
		fmt.Println("GOPATH is not set, please set it before running this program.")
		return
	}

	binDir := filepath.Join(goPath, "bin")
	autoGoPath := filepath.Join(binDir, "AutoGo.exe")

	if err := os.WriteFile(autoGoPath, autoGoExe, 0755); err != nil {
		fmt.Println("Failed to write AutoGo.exe:", err)
		return
	}

	fmt.Println("successfully")
}
