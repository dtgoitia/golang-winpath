/*
%SystemRoot%\system32\WindowsPowerShell\v1.0\;C:\Program Files (x86)\EasyPHP-DevServer-14.1VC11\binaries\php\php_runningversion;C:\ProgramData\Oracle\Java\javapath;C:\Program Files (x86)\Intel\iCLS Client\;C:\Program Files\Intel\iCLS Client\;C:\windows\system32;C:\windows;C:\windows\System32\Wbem;C:\windows\System32\WindowsPowerShell\v1.0\;C:\Program Files\Intel\Intel(R) Management Engine Components\DAL;C:\Program Files (x86)\Intel\Intel(R) Management Engine Components\DAL;C:\Program Files\Intel\Intel(R) Management Engine Components\IPT;C:\Program Files (x86)\Intel\Intel(R) Management Engine Components\IPT;c:\Program Files (x86)\Hewlett-Packard\HP Performance Advisor;C:\Program Files (x86)\Calibre2\;C:\Program Files\MiKTeX 2.9\miktex\bin\x64\;C:\Users\davidt\AppData\Local\atom\bin;C:\Python27;C:\Program Files;C:\Winnt;C:\Winnt\System32;C:\Program Files (x86)\Microsoft VS Code\bin;C:\Users\davidt\AppData\Roaming\npm;C:\Users\davidt\projects\toggl-parser
*/
package main

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// Clear command line
	ClearScreen()
	Test()

	// Get "PATH" environment variable
	//pathArray := GetPath()
	//log.Println("first path = ", pathArray[1])

	// Print current paths in PATH
	//ListPath()

	// Add path
	// - Doesn't let me set the new path!
	//AddPath("C:\\Fakefolder")

	// RemovePath - TODO

	// FindPath - TODO
}

// Test : testing function
func Test() {
	var (
		cmdOut []byte
		err    error
	)
	cmdExec := "C:\\Windows\\system32\\WindowsPowerShell\\v1.0\\powershell.exe"
	//cmdName := "$Env:path"
	cmdName := "[Environment]::SetEnvironmentVariable(\"TestVariable\", \"Test value\", \"User\")"
	if cmdOut, err = exec.Command(cmdExec, cmdName).Output(); err != nil {
		log.Fatal("ERROR:", err)
	}
	log.Println("\n\nOUTPUT:\n", string(cmdOut))
}

// GetPath : return array with paths in PATH
func GetPath() []string {
	pathVar := os.Getenv("PATH")
	pathArray := strings.Split(pathVar, ";")
	return pathArray
}

// ListPath : list current paths in PATH
func ListPath() {
	pathArray := strings.Split(os.Getenv("PATH"), ";")
	for i := range pathArray {
		log.Println(i, " ", pathArray[i])
	}
}

// AddPath : add path to PATH environment variable
func AddPath(newPath string) {
	// - Add path to array
	// - Concatenate array elements
	// - Update "PATH"
	var newPathVar string
	pathArray := strings.Split(os.Getenv("PATH"), ";")
	log.Println("Current path Array:\n", pathArray)
	newPathArray := append(pathArray, newPath)
	for i := range newPathArray {
		if i > 1 {
			newPathVar = newPathVar + ";" + newPathArray[i]
		} else {
			newPathVar = newPathArray[i]
		}
	}
	log.Print("Adding \"", newPath, "\" to PATH...")
	err := os.Setenv("PATH", newPathVar)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("NO ERRORs in os.Setenv")
	}
	log.Println("Done.")
	log.Println("PATH:\n", os.Getenv("PATH"))
	log.Println("\n\nnewPATH:\n", newPathVar)
}

// ClearScreen : clear command line
func ClearScreen() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
