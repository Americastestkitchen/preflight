package main

import (
	"bytes"
	"os/exec"
	"path/filepath"
	"strings"
)

type Language int

const (
	_            = iota
	CSS Language = iota
	CoffeeScript
	HTML
	JavaScript
	Ruby
	SCSS
	SQL
)

var extensions = map[string]Language{
	// Ruby
	".rb": Ruby,
	".rake": Ruby,
	".ru":   Ruby,
}

type File struct {
	Path     string
	Language Language
}

func NewFile(filename string) *File {
	f := &File{Path: filename}
	ext := filepath.Ext(f.Path)
	f.Language = extensions[ext]
	return f
}

func FindChanges() []*File {
	var out bytes.Buffer

	cmd := exec.Command("git", "diff", "HEAD", "--name-only", "--diff-filter=MA")
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		panic(err)
	}
	lines := strings.Split(out.String(), "\n")
	files := []*File{}
	for _, line := range lines {
		files = append(files, NewFile(line))
	}
	return files
}
