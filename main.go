package main

import (
	"os"
	"fmt"
	"io/ioutil"
	"html/template"
	"flag"
	"strings"
	"github.com/labstack/gommon/color"
)

// data where filecontents is stored, dictionary
type Content struct {
	// Title 		string
	FileData 	string
}

func thisDirectory(directory string) []os.FileInfo {
	allFiles, err := ioutil.ReadDir(directory)
	if err != nil {
		panic(err)
	}
	return allFiles
}

//	check if existent file -good
func checkFile(filename string) bool {
	if strings.Contains(filename, "."){
		fmt.Println("file already exists")
		return strings.Split(filename, ".")[1] == "txt"
	} else {
		return false
	}
}

func extensionAdd(filename string) string {
	filext := strings.Split(filename, ".")[0] + ".html" // adds .html extension
	return filext
}

// generate new template -good
func fileToTemplate(filename string) (string, *os.File) {
	filename = extensionAdd(filename)
	// bytesToWrite := []byte(temp)
	// con := content{FileData: readFile(cont)}
	// err := ioutil.WriteFile(filename, bytesToWrite, 0644)
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	
	return filename, file
}

// get contents from file to format -good
func readFile(filename string) string {
	fileContents, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(fileContents)
}

func renderTemplate(filename string) string {
	var err error
	var filecont Content
	filecont.FileData = readFile(filename)
	// bytesToWrite := []byte(filename)
	// buf := new(bytes.Buffer)
	// con := Content{FileData: cont}

	t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))
	// err := ioutil.WriteFile("new-file.txt", bytesToWrite, 0644)
	// err = t.Execute(os.Stdout, con)
	newfilea, newfileb := fileToTemplate(filename)
	err = t.Execute(newfileb, Content: filecont)
	if err != nil {
		panic(err)
	}

	return newfilea
	// fmt.Println(err)
	// return fmt.Sprintf("%s file created", err)
}


// hopefully will be adding more flags
// func allFlags(flg string) bool {
// 	runit := false
// 	flag.Visit(func(flgg *flag.Flag) {
// 		if flgg.Name == flg {
// 			runit = true
// 		}
// 	})
// 	return runit
// }

func dirFlagfunc(directory string) []os.FileInfo {
	// var allFiles string
	allfiles, err := ioutil.ReadDir(directory)
	if err != nil {
		panic(err)
	}
	return allfiles
}


func save() {
	// var err error
	// var content string
	// var filename *string
	// var dirname *string
	// FLAGS
	// var filename string
	// var dirname string

	fileFlag := flag.String("file", "", " name of file to render")
	dirFlag := flag.String("dir", ".", " all .txt files")

	flag.Parse()

	if *dirFlag != "" {
		allfiles := dirFlagfunc(*dirFlag)

		for _, file := range allfiles {
			filename := file.Name()
			if checkFile(filename) == true {
				temp := renderTemplate(filename)
				fmt.Sprintf("%b rendered", temp)
				fmt.Sprintf("%s file created", filename)
			}
		}
	} else {
		temp := renderTemplate(*fileFlag)
		fmt.Println("hello dere")
	}

	if *fileFlag != "" {
		allfiles := dirFlagfunc(*dirFlag)
		for _, file := range allfiles {
			filename := file.Name()
			fmt.Println(filename)
		}
	}

	// if dirFlagfunc("dir") {
	// 	allfiles := checkFile(*dirFlag)

	// 	for _, file := range files {
	// 		con := readFile(file)
	// 		temp := renderTemplate(con)
	// 		filename := strings.SplitN(file, ".", 2)[0] + ".html"
	// 		fileToTemplate(temp, filename)
	// 		fmt.Sprintf("%s file created", filename)
	// 	}
	// }
	
	// if dirFlagfunc("file") {
	// 	con := readFile(file)
	// 	temp := renderTemplate(con)
	// 	filename := strings.SplitN(*fileFlag, ".", 2)[0] + ".html"
	// 	fileToTemplate(temp, filename)
	// 	fmt.Sprintf("%s file created", filename)
	// }
}

	// fmt.Println("file: ", *filename)
// 	if *dirFlag != ""{
// 		var allFiles = 0
// 		var dirr := directory string
// 		files, err := ioutil.ReadDir(dirr)

// 		if err != nil {
// 			panic(err)
// 		}

// 		for _, f := range files {
// 			fname := f.Name()

// 			// check if extension already exists
// 			if checkFile(fname) == true {
// 				renderTemplate("template.tmpl", readFile(fname))
// 				fileToTemplate("template.tmpl", fname)
// 				allFiles += 1
// 				// fmt.Println(*dirname)
// 			}
// 		}

// 		fmt.Println("new files created")

// 	}
// 	if *fileFlag != "" {
// 		renderTemplate("template.tmpl", readFile(*fileFlag))
// 		fmt.Println("file: ", *fileFlag)
// 	} 
// 	// else {
// 	// 	err := filepath.Walk(".",
// 	// 		func(path string, info os.FileInfo, err error) error {
// 	// 			if err != nil {
// 	// 				return err
// 	// 			}
// 	// 			fmt.Println(path, info.Size())
// 	// 			return nil
// 	// 		})
// 	// 	fmt.Println("directory", *dirname)
// 	// }

// 	// if *dirname != ""{
// 	// 	var allFiles = 0
// 	// 	files, err := ioutil.ReadDir(*dirname)
// 	// 	if err != nil {
// 	// 		panic(err)
// 	// 	}
// 	// 	for _, f := range files {
// 	// 		filename := f.Name()
// 	// 		renderTemplate("template.tmpl", readFile(name))
// 	// 		fmt.Println(*dirname)
// 	// 		}
// 	// 	}
// 		// err := filepath.Walk(".",
// 		// 	func(path string, info os.FileInfo, err error) error {
// 		// 		if err != nil {
// 		// 			return err
// 		// 		}
// 		// 		fmt.Println(path, info.Size())
// 		// 		return nil
// 		// 	})
// 		// 	if err != nil {
// 		// 		log.Fatal(err)
// 		// 	}
// 		// allFiles, _ := ioutil.ReadDir(".")
// 		// // if err != nil {
// 		// // 	log.Fatal(err)
// 		// // }
// 		// for _, f := range allFiles {
// 		// 	fmt.Println(f)
// 		// }
// 	}

// 	/////////////////////////
// 	// fileContents, err := ioutil.ReadFile("onepost.txt")
// 	// fmt.Println(string(fileContents))

// 	// content := FileData {
// 	// 	Title: "first-post",
// 	// 	Content: string(fileContents),
// 	// }
// 	// t := template.Must(template.ParseFiles("template.tmpl"))

// 	// newfile, err := os.Create("onepost.html")
// 	// err = t.Execute(os.Stdout, content)
// 	// err = t.Execute(newfile, content)
// 	// if err != nil {
// 	// 	panic(err)
// 	// }
// // }

func main() {
	color.Println(color.Red("Welcome to the Static Site Generator!"))
	save()
	// return FileData{Content: string(content)}
}