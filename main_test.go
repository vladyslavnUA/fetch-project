package main

import (
	"testing"
	"github.com/Flaque/filet"
	"fmt"
)

func mainTest(t *testing.T) {
	defer filet.CleanUp(t)
	filet.TmpDir(t, "project")
	filet.File(t, "project/file.html", "")
}

func TestFile(t *testing.T) {  
	defer filet.CleanUp(t)
	
	fileBool := filet.Exists(t, "file.html")
	if fileBool {
		fmt.Println("file.html exists")
	}
}

func TestDir(t *testing.T) {
	defer filet.CleanUp(t)
	cssBoo := filet.Exists(t, "project/css")
	cssBool := filet.DirContains(t, "project/css", "main.css")
	if cssBoo && cssBool {
		fmt.Println("css directory with main.css exists")
	}
	jsBoo := filet.Exists(t, "project/js")
	jsBool := filet.DirContains(t, "project/js", "main.js")
	if jsBoo && jsBool{
		fmt.Println("js directory with main.js exists")
	}
}
