package gutil

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"
)

func ExpandPath(p string) string {
	var splitter string
	if runtime.GOOS == "windows" {
		splitter = `\`
	} else {
		splitter = `/`
	}
	if strings.HasPrefix(p, fmt.Sprintf(`~%s`, splitter)) {
		return filepath.Join(HomeDir(), p[2:])
	}
	absDir, err := filepath.Abs(p)
	if err != nil {
		return p
	}
	return absDir
}

func HomeDir() string {
	if home := os.Getenv("HOME"); home != "" {
		return home
	}
	if usr, err := user.Current(); err == nil {
		return usr.HomeDir
	}
	if currentDir, err := os.Getwd(); err == nil {
		return currentDir
	}
	return ""
}
