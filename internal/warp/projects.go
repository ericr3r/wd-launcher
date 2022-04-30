package warp

import (
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/lithammer/fuzzysearch/fuzzy"
)

type Project struct {
	Name string
	Path string
	ID   int
}

type Projects struct {
	names    []string
	namesMap map[string]Project
}

func (p Projects) Search(str string, logger *log.Logger) []Project {
	term := str
	parts := strings.Split(str, " ")
	if len(parts) > 1 {
		term = strings.Join(parts[1:], " ")
	}

	names := fuzzy.Find(term, p.names)

	projects := make([]Project, 0)

	if len(names) > 0 {
		for _, name := range names {
			projects = append(projects, p.namesMap[name])
		}
	}

	return projects
}

func (p Projects) Activate(index int, logger *log.Logger) error {
	name := p.names[index-1]
	project := p.namesMap[name]

	flag := fmt.Sprintf("--working-directory=%s", project.Path)
	cmd := exec.Command("tilix", flag)

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
