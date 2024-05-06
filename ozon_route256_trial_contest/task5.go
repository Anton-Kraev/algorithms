package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Folder struct {
	Dir     string   `json:"dir"`
	Files   []string `json:"files"`
	Folders []Folder `json:"folders"`
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscanln(in, &t)
	for i := 0; i < t; i++ {
		var n int
		fmt.Fscanln(in, &n)
		var jsonDirectories []byte
		for j := 0; j < n; j++ {
			line, _ := in.ReadBytes('\n')
			jsonDirectories = append(jsonDirectories, line...)
		}
		var directories Folder
		json.Unmarshal(jsonDirectories, &directories)
		_, infected := searchInfected(directories)
		fmt.Fprintln(out, infected)
	}
}

func searchInfected(directories Folder) (int, int) {
	dirInfected := false
	files, infected := 0, 0
	for _, file := range directories.Files {
		files++
		if strings.HasSuffix(file, ".hack") {
			dirInfected = true
		}
	}
	for _, folder := range directories.Folders {
		folderFiles, folderInfected := searchInfected(folder)
		files += folderFiles
		infected += folderInfected
	}
	if dirInfected {
		return files, files
	}
	return files, infected
}
