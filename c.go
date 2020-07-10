package main

import (
	"fmt"
	"io"
    "io/ioutil"
	"os"
	"html/template"
	"log"
    // "path/filepath"
)

// func check(e error) {
//     if err != nil {
//         panic(err)
//     }
// }
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

	// A sample config
	// config := map[string]string{
	// 	"textColor":      "#abcdef",
	// 	"linkColorHover": "#ffaacc",
	// }
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

// func fileToTemplate() {
// 	t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))
// 	newfilea, newfileb := fileToTemplate(filename)
// 	err = t.Execute(newfileb)
// 	if err != nil {
// 		panic(err)
// 	}
// }

func main() {
	var count int
	fmt.Print(` Enter the amount of html pages 
you want to generate: `)
	fmt.Scanf("%s", &count)

	// creating new directory to store files
    err := os.Mkdir("project", 0755)
    if err != nil {
        panic(err)
    }

    // defer os.RemoveAll("subdir")

    // createFile("project/file.html")

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

	file, _ := os.Create("project/file.html")
	// file.Write([]byte("Change this text"))
	file.Close()
	templateTo("template.tmpl")
	file, _ = os.Open("project/file.html")
	io.Copy(os.Stdout, file)
	file.Close()
}