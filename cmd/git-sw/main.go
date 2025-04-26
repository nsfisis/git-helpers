package main

import (
	"log"
	"os"
	"os/exec"
	"strings"
	"unicode"
)

func main() {
	gitArgs := []string{"switch"}
	if requiresDetachFlag(os.Args) {
		gitArgs = append(gitArgs, "--detach")
	}
	firstPositionalArg := true
	for i, argv := range os.Args {
		if i == 0 {
			continue // argv[0] is a program name.
		}
		if firstPositionalArg && !strings.HasPrefix(argv, "-") {
			if isInt(argv) {
				argv = "feature/" + argv
			}
			firstPositionalArg = false
		}
		gitArgs = append(gitArgs, argv)
	}

	cmd := exec.Command("git", gitArgs...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		switch err.(type) {
		case *exec.ExitError:
			// Do nothing here because Git has already reported the error.
		default:
			log.Fatal(err)
		}
	}

	os.Exit(cmd.ProcessState.ExitCode())
}

func requiresDetachFlag(argv []string) bool {
	argc := len(argv)
	if argc == 1 {
		return false
	}
	firstArg := argv[1]

	// Example: origin/main, upstream/develop
	if strings.HasPrefix(firstArg, "origin/") || strings.HasPrefix(firstArg, "upstream/") {
		return true
	}

	// Example: 1234, cafebabe
	if len(firstArg) >= 4 {
		for _, c := range firstArg {
			if !unicode.Is(unicode.ASCII_Hex_Digit, c) {
				return false
			}
		}
		return true
	}

	return false
}

func isInt(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}
