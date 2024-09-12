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
	// 获取 GOPATH
	goPath := os.Getenv("GOPATH")
	if goPath == "" {
		fmt.Println("GOPATH is not set, please set it before running this program.")
		return
	}

	// 构造目标路径
	binDir := filepath.Join(goPath, "bin")
	autoGoPath := filepath.Join(binDir, "AutoGo.exe")

	// 写入 AutoGo.exe 到 GOPATH/bin
	if err := os.WriteFile(autoGoPath, autoGoExe, 0755); err != nil {
		fmt.Println("Failed to write AutoGo.exe:", err)
		return
	}

	fmt.Println("AutoGo.exe has been successfully installed to", autoGoPath)
}
