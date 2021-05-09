package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

const (
	pref_st    = "├───"
	pref_last  = "└───"
	pref_tab   = "│\t"
	pref_space = "\t"
)

type Line interface {
	String() string
}

func dirTree(out io.Writer, path string, printFiles bool) error {
	err, lines := readDir(path, []Line{}, printFiles)
	formatTree(out, lines, []string{})

	// for _, line := range lines {
	// 	fmt.Printf("%s\n", line)
	// }
	return err
}

func formatTree(out io.Writer, lines []Line, prefixes []string) {
	for idx, line := range lines {
		if len(lines)-1 == idx {
			fmt.Fprintf(out, "%s%s%s\n", strings.Join(prefixes, ""), pref_last, line)
			if directory, ok := line.(Directory); ok {
				formatTree(out, directory.children, append(prefixes, pref_space))
			}
		} else {
			fmt.Fprintf(out, "%s%s%s\n", strings.Join(prefixes, ""), pref_st, line)
			if directory, ok := line.(Directory); ok {
				formatTree(out, directory.children, append(prefixes, pref_tab))
			}
		}
	}
}

func readDir(path string, lines []Line, printFiles bool) (error, []Line) {
	file, err := os.Open(path)
	if err != nil {
		return err, nil
	}
	files, err := file.Readdir(0)
	if err != nil {
		return err, nil
	}
	file.Close()

	sort.Slice(files, func(i, j int) bool {
		return files[i].Name() < files[j].Name()
	})

	for _, file := range files {
		if !(file.IsDir() || printFiles) {
			continue
		}

		var line Line
		if file.IsDir() {
			_, children := readDir(filepath.Join(path, file.Name()), []Line{}, printFiles)
			line = Directory{file.Name(), children}
		} else {
			line = File{file.Name(), file.Size()}
		}

		lines = append(lines, line)
	}

	return err, lines
}

type Directory struct {
	name     string
	children []Line
}

func (directory Directory) String() string {
	return directory.name
}

type File struct {
	name string
	size int64
}

func (file File) String() string {
	if file.size == 0 {
		return file.name + " (empty)"
	} else {
		return file.name + " (" + strconv.FormatInt(file.size, 10) + "b)"
	}
}
func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}
