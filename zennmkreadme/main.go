package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

func grepTitle(fname string) (string, error) {
	fd, err := os.Open(fname)
	if err != nil {
		return "", err
	}
	defer fd.Close()

	sc := bufio.NewScanner(fd)
	for sc.Scan() {
		line := sc.Text()
		const header = "title: "
		if strings.HasPrefix(line, header) {
			return strings.Trim(line[len(header):], " \"\r\n"), nil
		}
	}
	return "", io.EOF
}

var rxPagePattern = regexp.MustCompile(`^(\d+)\..*\.md$`)

func atoi(s string) int {
	n := 0
	for len(s) > 0 {
		i := strings.IndexByte("0123456789", s[0])
		if i < 0 {
			break
		}
		n = n*10 + i
		s = s[1:]
	}
	return n
}

func mkreadme(dir string) (string, [][2]string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return "", nil, err
	}

	var bookTitle string
	result := make([][2]string, 0, len(entries))
	for _, entry := range entries {
		name := entry.Name()
		thePath := filepath.Join(dir, name)

		if name == "config.yaml" {
			if title, err := grepTitle(thePath); err == nil {
				bookTitle = title
			}
		} else if m := rxPagePattern.FindStringSubmatch(name); m != nil {
			if title, err := grepTitle(thePath); err == nil {
				result = append(result, [2]string{
					title, filepath.ToSlash(name)})
			}
		}
	}
	sort.Slice(result, func(i, j int) bool {
		return atoi(result[i][1]) < atoi(result[j][1])
	})
	return bookTitle, result, nil
}

func mains() error {
	title, pages, err := mkreadme(".")
	if err != nil {
		return err
	}
	fmt.Println(title)
	fmt.Println("==============")
	fmt.Println()
	for i, c := range pages {
		fmt.Printf("%d. [%s](./%s)\n", i+1, c[0], c[1])
	}
	fmt.Println()
	fmt.Println("![cover](./cover.jpg)")
	return nil
}

func main() {
	if err := mains(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
