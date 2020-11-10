/*
git-prompt

To use just add $(git-prompt) to PS1, for example.

    PS1='\u@\h:\w $(git-prompt)$ '

Ole Martin Bjorndalen

License: MIT
*/

package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func readCommand(cmdName string, cmdArgs []string) string {
	var (
		cmdOut []byte
		err    error
	)
	if cmdOut, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		return ""
	}
	return string(cmdOut)
}

func getStatusText() string {
	return readCommand("git", []string{"status", "--porcelain", "--branch"})
}

func getHeadHash() string {
	return readCommand("git", []string{"rev-parse", "--verify", "HEAD"})
}

func splitIntoLines(text string) []string {
	return strings.Split(strings.TrimSpace(text), "\n")
}

func getBranchName(line string) string {
	if line == "## No commits yet on master" || line == "## Initial commit on master" {
		return ":initial"
	}

	return strings.Split(strings.Fields(line)[1], "...")[0]
}

func addBranchStatus(status map[string]bool, line string) {
	// We're looking for what's inside the brackets:
	// ## master...origin/master [ahead 1, behind 1]
	if strings.Contains(line, "[") {
		statusPart := strings.Split(line, "[")[1]
		if strings.Contains(statusPart, "ahead") {
			status["ahead"] = true
		}
		if strings.Contains(statusPart, "behind") {
			status["behind"] = true
		}
	}
}

func addFileStatus(status map[string]bool, lines []string) {
	conflictCodes := map[string]bool{
		"DD": true,
		"AU": true,
		"UD": true,
		"UA": true,
		"DU": true,
		"AA": true,
		"UU": true,
	}

	for _, line := range lines {
		xy := line[:2]

		if xy == "??" {
			status["untracked"] = true
		} else if conflictCodes[xy] {
			status["conflict"] = true
		} else {
			status["changed"] = true
		}
	}
}

func parseStatus(text string) (string, map[string]bool) {
	status := make(map[string]bool)
	lines := splitIntoLines(text)

	branch := getBranchName(lines[0])

	addBranchStatus(status, lines[0])
	addFileStatus(status, lines[1:])

	return branch, status
}

func colorText(text string, color string) string {
	colors := map[string]string{
		"green":  "92",
		"yellow": "93",
		"red":    "31",
	}

	return "\001\033[" + colors[color] + "m\002" + text + "\001\033[0m\002"
}

func formatStatus(branch string, status map[string]bool) string {
	flags := ""
	color := "green"

	if status["changed"] {
		flags += "*"
		color = "yellow"
	}

	if status["untracked"] {
		flags += "?"
		color = "yellow"
	}

	if status["conflict"] {
		flags += "!"
		color = "red"
	}

	if status["ahead"] && status["behind"] {
		flags += "↕"
	} else if status["ahead"] {
		flags += "↑"
	} else if status["behind"] {
		flags += "↓"
	}

	if flags != "" {
		flags = " " + flags
	}

	text := fmt.Sprintf("[%s%s]", branch, flags)

	return colorText(text, color)
}

func main() {
	text := getStatusText()
	if text != "" {
		branch, status := parseStatus(text)
		if branch == "HEAD" {
			branch = ":" + getHeadHash()[:6]
		}

		fmt.Print(formatStatus(branch, status))
	}
}
