package lcu

import (
	"os"
	"runtime"
	"strings"
)

func locfile_path() string {
	var path string
	if runtime.GOOS == "darwin" {
		path = "/Applications/League of Legends.app/Contents/Lol/lockfile"
	}
	if runtime.GOOS == "windows" {
		path = `C:\Riot Games\League of Legends\lockfile`
	}
	return path
}

func open_lockfile() (string, string, error) {
	path := locfile_path()
	file, err := os.ReadFile(path)
	if err != nil {
		return "", "", err
	}
	data := strings.Split(string(file), ":")
	return data[2], data[3], nil
}
