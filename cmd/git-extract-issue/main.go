package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	argv := os.Args
	argc := len(argv)
	if argc != 2 {
		return
	}
	branchName := argv[1]
	fmt.Println(extractIssueNumberFromBranchName(branchName))
}

// * 123                       => #123
// * 123-suffix                => #123
// * feature/123               => #123
// * feature/123-suffix        => #123
// * feature/prefix-123        => prefix-123
// * feature/prefix-123-suffix => prefix-123
func extractIssueNumberFromBranchName(branchName string) string {
	pattern := regexp.MustCompile(`\A(?:\w+/)?(\w+-)?(\d+)(?:-\w+)*\z`)
	matches := pattern.FindSubmatch([]byte(branchName))
	if len(matches) != 3 {
		return ""
	}
	var prefix string
	if len(matches[1]) == 0 {
		prefix = "#"
	} else {
		prefix = string(matches[1])
	}
	return prefix + string(matches[2])
}
