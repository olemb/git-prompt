/*
git-prompt

To use just add $(git-prompt) to PS1, for example.

    PS1='\u@\h:\w $(git-prompt)$ '

Ole Martin Bjorndalen
https://github.com/olemb/git-prompt

License: MIT
*/

package main

import (
	"fmt"
	"os/exec"
	"strings"
	"unicode"
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
	return readCommand("git", []string{"status", "--porcelain=v2", "--branch"})
}

func parseStatus(text string) (string, map[string]bool) {
	oid := ""
	head := ""
	branch := ""
	status := make(map[string]bool)

	lines := strings.Split(strings.TrimSpace(text), "\n")
	for _, line := range lines {
		char := line[0]
		if char == '#' {
			words := strings.Fields(line)
			name := words[1]
			args := words[2:]

			if name == "branch.oid" {
				oid = args[0]
			} else if name == "branch.head" {
				head = args[0]
			} else if name == "branch.ab" {
				if args[0] != "+0" {
					status["ahead"] = true
				}
				if args[1] != "-0" {
					status["behind"] = true
				}
			}
		} else if char == '?' {
			status["untracked"] = true
		} else if char == 'u' {
			status["conflict"] = true
		} else if unicode.IsDigit(rune(char)) {
			status["changed"] = true
		}
	}

	if oid == "(initial)" {
		branch = ":initial"
	} else if head == "(detached)" {
		branch = ":" + oid[:6]
	} else {
		branch = head
	}

	return branch, status
}

func formatStatus(branch string, status map[string]bool) string {
	green := "92"
	yellow := "93"
	red := "31"

	flags := ""
	color := green

	if status["changed"] {
		flags += "*"
		color = yellow
	}

	if status["untracked"] {
		flags += "?"
		color = yellow
	}

	if status["conflict"] {
		flags += "!"
		color = red
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

	return fmt.Sprintf("\033[%sm[%s%s]\033[0m", color, branch, flags)
}

func main() {
	text := getStatusText()
	if text != "" {
		branch, status := parseStatus(text)
		fmt.Print(formatStatus(branch, status))
	}
}
