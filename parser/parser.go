// Parser implements a parser that parses `upspin share` output.
package parser

import (
	"bufio"
	"bytes"
	"strings"
)

// Parse parses the `upspin share` output. maps a list of files that can be
// accessed by each user to users.
func Parse(data []byte) (map[string][]string, error) {
	buf := bytes.NewBuffer(data)
	scanner := bufio.NewScanner(buf)

	rv := make(map[string][]string)
	var p parser = parserStart
	var err error
	for p != nil {
		p, err = p(scanner, rv)
		if err != nil {
			return nil, err
		}
	}

	return rv, nil
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
		users := strings.Fields(scanner.Text())
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
