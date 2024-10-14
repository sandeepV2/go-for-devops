package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	_ "embed"
)

//go:embed hello.txt
var s string

func main() {
    fmt.Println(runtime.GOOS) // linux, darwin, ...
    fmt.Println(runtime.GOARCH) // amd64, arm64, ...

	 wd, err := os.Getwd()
    if err != nil {
        panic(err) 
    }
    fmt.Println(wd)
    content, err := os.ReadFile(filepath.Join(wd, "config", "config.json"))
    if err != nil {
        panic(err) 
    }
    fmt.Println(string(content))


	fmt.Println(s)
}