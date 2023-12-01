package test

import (
	"embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"testing"
)

//go:embed version.txt
var version string

func TestString(t *testing.T) {
	fmt.Println(version)
}

//go:embed logo.png
var logo []byte

func TestByte(t *testing.T) {
	err := ioutil.WriteFile("logo_new.png", logo, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}

//go:embed file/a.txt
//go:embed file/b.txt
//go:embed file/c.txt
var files embed.FS

func TestMultipleFile(t *testing.T) {
	a, _ := files.ReadFile("file/a.txt")
	b, _ := files.ReadFile("file/b.txt")
	c, _ := files.ReadFile("file/c.txt")

	fmt.Println(string(a))
	fmt.Println(string(b))
	fmt.Println(string(c))
}

//go:embed file/*.txt
var path embed.FS

func TestPatchMatcher(t *testing.T) {
	dirEntries, _ := path.ReadDir("file")
	for _, entry := range dirEntries {
		if !entry.IsDir() {
			fmt.Println(entry.Name())

			file, _ := path.ReadFile("file/" + entry.Name())
			fmt.Println(string(file))
		}
	}
}
