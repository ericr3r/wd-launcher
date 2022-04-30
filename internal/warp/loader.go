package warp

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

func Load() (*Projects, error) {
	userHome, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	fileName := filepath.Join(userHome, ".warprc")

	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	names := make([]string, 0)
	namesMap := make(map[string]Project)
	i := 0
	for scanner.Scan() {
		name, path := parse(userHome, scanner.Text())

		if _, err := os.Stat(path); err == nil {
			names = append(names, name)
			namesMap[name] = Project{Path: path, ID: i + 1, Name: name}
			i++
		}
	}

	return &Projects{names, namesMap}, nil
}

func parse(homeDir string, entry string) (string, string) {
	parts := strings.Split(entry, ":")

	path := parts[1]
	if path == "~" {
		path = homeDir
	} else if strings.HasPrefix(path, "~/") {
		path = filepath.Join(homeDir, path[2:])
	}

	return parts[0], path
}
