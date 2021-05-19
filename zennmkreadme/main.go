package main

import (
	"bufio"
	"flag"
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

// Link contains title and url.
type Link struct {
	Title    string
	URL      string
	filename string // for sort
}

// Book contains its title and chapters.
type Book struct {
	Title   string
	Chapter []*Link
}

func readBook(dir string) (*Book, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var bookTitle string
	chapters := make([]*Link, 0, len(entries))
	for _, entry := range entries {
		name := entry.Name()
		thePath := filepath.Join(dir, name)

		if name == "config.yaml" {
			if title, err := grepTitle(thePath); err == nil {
				bookTitle = title
			}
		} else if m := rxPagePattern.FindStringSubmatch(name); m != nil {
			if title, err := grepTitle(thePath); err == nil {
				chapters = append(chapters, &Link{
					Title:    title,
					URL:      filepath.ToSlash(thePath),
					filename: name})
			}
		}
	}
	sort.Slice(chapters, func(i, j int) bool {
		return atoi(chapters[i].filename) < atoi(chapters[j].filename)
	})
	return &Book{Title: bookTitle, Chapter: chapters}, nil
}

func (b *Book) dump(w io.Writer) {
	fmt.Fprintln(w, b.Title)
	fmt.Fprintln(w, "==============")
	fmt.Fprintln(w)
	for i, c := range b.Chapter {
		fmt.Fprintf(w, "%d. [%s](./%s)\n", i+1, c.Title, c.URL)
	}
	fmt.Fprintln(w)
	fmt.Fprintln(w, "![cover](./cover.jpg)")
}

var flagOutput = flag.String("o", "", "output table of contents to")

func mains(args []string) error {
	var tableWriter io.Writer = os.Stdout
	if *flagOutput != "" {
		w, err := os.Create(*flagOutput)
		if err != nil {
			return err
		}
		defer w.Close()
		tableWriter = w
	}
	if len(args) <= 0 {
		b, err := readBook(".")
		if err != nil {
			return err
		}
		b.dump(tableWriter)
	}
	for _, dir := range args {
		b, err := readBook(dir)
		if err != nil {
			return err
		}
		b.dump(tableWriter)
	}
	return nil
}

func main() {
	flag.Parse()
	if err := mains(flag.Args()); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
