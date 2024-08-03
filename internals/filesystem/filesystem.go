package filesystem

import (
	"os"
	"strings"
)

type File struct {
	Path     string
	Name     string
	Children []*File
}

func CopyRecursive(src string, dst string) {
	os.RemoveAll(dst)
	os.MkdirAll(dst, 0755)
	var files File
	GetFileNames(&files, src, "")
	createFiles(&files, src, dst)
}

func GetFileNames(f *File, src string, currentFolder string) {
	entries, _ := os.ReadDir(currentFolder + src)
	currentFolder += src + "/"
	for _, entry := range entries {
		var f1 File
		f.Children = append(f.Children, &f1)
		f1.Path = currentFolder
		if entry.IsDir() {
			GetFileNames(&f1, entry.Name(), currentFolder)
		} else {
			f1.Name = entry.Name()
		}
	}
}

func createFiles(files *File, src string, dst string) {
	for _, f := range files.Children {
		os.MkdirAll(strings.Replace(f.Path, src, dst, -1), 0755)
		if f.Name != "" {
			oldname := f.Path + f.Name
			newname := strings.Replace(oldname, src, dst, -1)
			os.Link(oldname, newname)
		} else {
			createFiles(f, src, dst)
		}
	}
}
