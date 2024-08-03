package main

import (
	"bufio"
	"bytes"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/thneutral/static-site-generator/internals/block"
	"github.com/thneutral/static-site-generator/internals/filesystem"
)

const (
	TEMPLATE_HTML = "./template.html"

	STATIC_NAME  = "./static"
	PUBLIC_NAME  = "./public"
	CONTENT_NAME = "./content"

	MD_EXTENSION   = ".md"
	HTML_EXTENSION = ".html"
)

type HTMLContent struct {
	Title   string
	Content template.HTML
}

func mdToHTML(md string) *HTMLContent {
	var htmlcontent HTMLContent
	contentString := ""
	hasSetTitle := false
	for _, node := range block.SplitBlocks(md) {
		if !hasSetTitle && node.NodeType == block.HEADING1_TYPE {
			htmlcontent.Title = node.Text[0].Text
			hasSetTitle = true
		}
		contentString += node.ToHTML()
	}
	contentString += ""
	htmlcontent.Content = template.HTML(contentString)
	return &htmlcontent
}

func processContentFolder(files *filesystem.File) {
	os.MkdirAll(strings.Replace(files.Path, CONTENT_NAME, PUBLIC_NAME, -1), 0755)
	if files.Name != "" {
		createHTMLFromMD(files)
	}
	for _, file := range files.Children {
		processContentFolder(file)
	}
}

func createHTMLFromMD(file *filesystem.File) {
	var buf bytes.Buffer
	in, _ := os.OpenFile(file.Path+file.Name, os.O_RDONLY, 0644)
	io.Copy(&buf, in)
	in.Close()

	htmlcontent := mdToHTML(buf.String())
	buf.Reset()
	templFile, _ := os.OpenFile(TEMPLATE_HTML, os.O_RDONLY, 0644)
	io.Copy(&buf, templFile)
	templFile.Close()

	out, _ := os.Create(strings.Replace(strings.Replace(file.Path+file.Name, CONTENT_NAME, PUBLIC_NAME, -1), MD_EXTENSION, HTML_EXTENSION, -1))
	writer := bufio.NewWriter(out)
	templ := template.Must(template.New("template").Parse(buf.String()))
	templ.Execute(writer, htmlcontent)
	writer.Flush()
	out.Close()
}

func main() {
	filesystem.CopyRecursive(STATIC_NAME, PUBLIC_NAME)
	var files filesystem.File
	filesystem.GetFileNames(&files, CONTENT_NAME, "")
	processContentFolder(&files)

	http.Handle("/", http.FileServer(http.Dir(PUBLIC_NAME)))

	log.Println("Listening on 8888")
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
