package models

import (
	"regexp"
	"strings"
)

type GlobalStatus struct {
	ID        string
	Name      string
	Provider  string
	State     string
	Directory string
}

func ResolveGlobalStatus(content string) (globalStatusList []GlobalStatus) {
	for index, line := range strings.Split(content, "\n") {
		if index == 0 || index == 1 {
			continue
		}
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			break
		}
		parts := regexp.MustCompile("\\s+").Split(line, -1)
		if len(parts) == 5 {
			globalStatusList = append(globalStatusList, GlobalStatus{
				ID: parts[0], Name: parts[1], Provider: parts[2],
				State: parts[3], Directory: parts[4],
			})
		}
	}
	return
}

func GetVIDByWorkPath(globalStatusList []GlobalStatus, workpath string) (VID string) {
	for _, globalStatus := range globalStatusList {
		if globalStatus.Directory == workpath {
			VID = globalStatus.ID
			return
		}
	}
	return
}
