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

// Link contains title and url.
type Link struct {
	Title    string
	URL      string
	filename string // for sort
}

// Book contains its title and chapters.
type Book struct {
	Title      string
	Chapter    []*Link
	urlBaseDir string
}

func readBook(dir string, urlBaseDir string) (*Book, error) {
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
					URL:      filepath.ToSlash(filepath.Join(urlBaseDir, name)),
					filename: name})
			}
		}
	}
	sort.Slice(chapters, func(i, j int) bool {
		return atoi(chapters[i].filename) < atoi(chapters[j].filename)
	})
	return &Book{Title: bookTitle, Chapter: chapters, urlBaseDir: urlBaseDir}, nil
}

func (b *Book) dump(w io.Writer) {
	fmt.Fprintln(w, b.Title)
	fmt.Fprintln(w, "==============")
	fmt.Fprintln(w)
	for i, c := range b.Chapter {
		fmt.Fprintf(w, "%d. [%s](%s)\n", i+1, c.Title, c.URL)
	}
	fmt.Fprintln(w)
	fmt.Fprintf(w, "![cover](%s)\n",
		filepath.ToSlash(filepath.Join(b.urlBaseDir, "cover.jpg")))
}

func mains() error {
	books, err := os.ReadDir("./books")
	if err != nil {
		return err
	}
	for _, bookDir1 := range books {
		if !bookDir1.IsDir() {
			continue
		}
		b, err := readBook(filepath.Join("./books", bookDir1.Name()), bookDir1.Name())
		if err != nil {
			return err
		}
		bookIndexPath := filepath.Join("./books", bookDir1.Name()+".md")
		fd, err := os.Create(bookIndexPath)
		if err != nil {
			return err
		}
		b.dump(fd)
		fd.Close()

		fmt.Printf("* [%s](%s)\n", b.Title, filepath.ToSlash(bookIndexPath))
	}
	return nil
}

func main() {
	if err := mains(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
