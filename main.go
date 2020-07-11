package main

import (
	"fmt"
    "io/ioutil"
	"os"
	"html/template"
	"log"
	"strings"

)

func templateTo(path string) {
	temp, err := template.ParseFiles(path)
	if err != nil {
		panic(err)
		return
	}

	file, err := os.Create(path)
	if err != nil {
		fmt.Println("Created file: ", err)
		return
	}

	config := "change this text"

	err = temp.Execute(file, config)
	if err != nil {
		log.Print("execute: ", err)
		return
	}
	file.Close()
}

func createFile(name string) {
	d := []byte("")
	err := ioutil.WriteFile(name, d, 0644)
	if err != nil {
		panic(err)
	}
}

type content struct {
	FileData 	string
}

// generate new template
func fileToTemplate(filename string) (string, *os.File) {
	filename = strings.Split(filename, ".")[0] + ".html" // adds .html extension
	file, err := os.Create("project/file.html")
	if err != nil {
		panic(err)
	}
	return filename, file
}

func temp() {
	newtemp := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))
	file := "file"
	_, newfileb := fileToTemplate(file)
	err := newtemp.Execute(newfileb, &content{"Change this text"})
	if err != nil {
		panic(err)
	}
}

func main() {	
	fmt.Print("Welcome to Fetch Project. A static project generator.\n")
	fmt.Print("Your new project was generated in the directory.\n")
    err := os.Mkdir("project", 0755)
    if err != nil {
        panic(err)
    }

    err = os.MkdirAll("project/js", 0755)
    if err != nil {
        panic(err)
	}
	err = os.MkdirAll("project/css", 0755)
    if err != nil {
        panic(err)
    }

    createFile("project/css/main.css")
    createFile("project/js/main.js")
	temp()
}