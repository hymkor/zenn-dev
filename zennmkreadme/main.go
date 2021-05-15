package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"

	"gopkg.in/yaml.v2"
)

type YamlWithTitle struct {
	Title string `yaml:"title"`
}

var rxPagePattern = regexp.MustCompile(`^(\d+)\..*\.md$`)

func readMdFileHeader(name string) (*YamlWithTitle, error) {
	fd, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer fd.Close()

	var yamlBuf bytes.Buffer

	hrCount := 0
	sc := bufio.NewScanner(fd)
	for sc.Scan() {
		line := sc.Text()
		if line == "---" {
			hrCount++
			if hrCount >= 2 {
				break
			}
		} else if hrCount == 1 {
			yamlBuf.WriteString(line)
			yamlBuf.WriteString("\n")
		}
	}
	y := &YamlWithTitle{}
	err = yaml.Unmarshal(yamlBuf.Bytes(), y)
	return y, err
}

type Chapter struct {
	Index    int
	Title    string
	Filename string
}

func mkreadme(dir string) (string, []Chapter, error) {
	fd, err := os.Open(dir)
	if err != nil {
		return "", nil, err
	}
	defer fd.Close()

	dirs, err := fd.ReadDir(-1)
	if err != nil {
		return "", nil, err
	}

	var bookTitle string
	result := make([]Chapter, 0, 20)
	for _, dir1 := range dirs {
		name := dir1.Name()

		if name == "config.yaml" {
			configYamlBin, err := os.ReadFile(name)
			if err != nil {
				return "", nil, err
			}
			configYaml := &YamlWithTitle{}
			err = yaml.Unmarshal(configYamlBin, configYaml)
			if err != nil {
				return "", nil, err
			}
			bookTitle = configYaml.Title
			continue
		}
		m := rxPagePattern.FindStringSubmatch(name)
		if m == nil {
			continue
		}
		y, err := readMdFileHeader(name)
		if err != nil {
			return "", nil, err
		}
		index, err := strconv.Atoi(m[1])
		if err != nil {
			return "", nil, err
		}
		result = append(result, Chapter{Index: index, Title: y.Title, Filename: name})
	}
	sort.Slice(result, func(i, j int) bool {
		if result[i].Index != result[j].Index {
			return result[i].Index < result[j].Index
		}
		return result[i].Filename < result[j].Filename
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
		fmt.Printf("%d. [%s](./%s)\n", i+1, c.Title, c.Filename)
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
