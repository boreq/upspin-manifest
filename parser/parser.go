// Package parser parses `upspin share` output.
package parser

import (
	"bufio"
	"bytes"
	"path"
	"strings"
)

// Parse parses the `upspin share` output. maps a list of files that can be
// accessed by each user to users.
func Parse(data []byte) (map[string][]string, map[string][]string, error) {
	buf := bytes.NewBuffer(data)
	scanner := bufio.NewScanner(buf)

	// Parse the output to produce userFiles
	userFiles := make(map[string][]string)
	var p parser = parserStart
	var err error
	for p != nil {
		p, err = p(scanner, userFiles)
		if err != nil {
			return nil, nil, err
		}
	}

	// Process the userFiles to create userDirectories (as of now directory
	// output is not supported by upspin share)
	userDirectories := make(map[string][]string)
	for user, files := range userFiles {
		for _, file := range files {
			directory, filename := path.Split(file)
			if filename == "Access" {
				directories, ok := userDirectories[user]
				if !ok {
					directories = make([]string, 0)
				}
				directories = append(directories, directory)
				userDirectories[user] = directories
			}
		}
	}

	return userFiles, userDirectories, nil
}

type parser func(scanner *bufio.Scanner, data map[string][]string) (parser, error)

func parserStart(scanner *bufio.Scanner, data map[string][]string) (parser, error) {
	for scanner.Scan() {
		if scanner.Text() == "files readable by:" {
			return parserUsers, nil
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return nil, nil
}

func parserUsers(scanner *bufio.Scanner, data map[string][]string) (parser, error) {
	ok := scanner.Scan()
	if ok {
		users := strings.Fields(strings.Trim(scanner.Text(), ":"))
		return createParserFiles(users, scanner, data), nil
	}
	return nil, scanner.Err()
}

func createParserFiles(users []string, scanner *bufio.Scanner, data map[string][]string) parser {
	return func(scanner *bufio.Scanner, data map[string][]string) (parser, error) {
		for scanner.Scan() {
			if scanner.Text() == "" {
				return parserStart, nil
			} else {
				file := strings.TrimSpace(scanner.Text())
				for _, user := range users {
					files, ok := data[user]
					if !ok {
						files = make([]string, 0)
					}
					data[user] = append(files, file)
				}
			}
		}
		if err := scanner.Err(); err != nil {
			return nil, err
		}
		return nil, nil
	}
}
