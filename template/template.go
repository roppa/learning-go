package main

import (
	"bytes"
	"fmt"
	"html/template"
)

// HelloMessage type contains Message string
type HelloMessage struct {
	Message string
}

// Generate an html string from path and Message struct
func getHTMLFromTemplate(path string, data HelloMessage) string {
	tmpl := template.Must(template.ParseFiles(path))
	var templateString bytes.Buffer
	tmpl.Execute(&templateString, data)
	return templateString.String()
}

func main() {
	path := "./templates/hello.html"
	hello1 := HelloMessage{Message: "Hello one!"}
	hello2 := HelloMessage{Message: "Hello two!"}
	fmt.Println(getHTMLFromTemplate(path, hello1))
	fmt.Println(getHTMLFromTemplate(path, hello2))
}
