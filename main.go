package main

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"regexp"

	"github.com/creack/pty"
)


func makefileHasTarget(target string, path string) (bool, error) {
	fileContent, err := ioutil.ReadFile(path)
	if err != nil {
		return false, err
	}

	stringMatch := fmt.Sprintf(`\n?%s:`, target)

	return regexp.Match(stringMatch, fileContent)
}

func exists(name string) (bool, error) {
	_, err := os.Stat(name)
	if errors.Is(err, os.ErrNotExist) {
	  return false, nil
	}
	return err == nil, err
}

func findMakeFileWithTarget(cwd, target string) (string, error) {
	hasMakefile, err := exists(cwd + "Makefile")
	if err != nil {
		return "", err
	}

	if hasMakefile {
		commandExists, err := makefileHasTarget(target, cwd + "Makefile")
		if err != nil {
			return "", err
		}

		if commandExists {
			return cwd, nil
		}
	}

	if isGitRoot, _ := exists(cwd + ".git"); !isGitRoot {
		return findMakeFileWithTarget("./." + cwd, target)
	}


	return "", errors.New("Could not find a makefile with target: " + target)
}

func runMake(dir string, args ...string) error {
	dirCommand := []string{"-C", dir}
	allArgs := append(dirCommand, args...)
	cmd := exec.Command("make",  allArgs...)

	f, err := pty.Start(cmd)
    if err != nil {
        return err
    }

    io.Copy(os.Stdout, f)

	return nil
}

func main() {
	if len(os.Args[1:]) < 1 {
		log.Fatalln("You must provide a make target")
	}

	makeFileDir, err := findMakeFileWithTarget("./", os.Args[1])
	if err != nil {
		log.Fatalln("Could not find a makefile:", err.Error())
	}

	err = runMake(makeFileDir, os.Args[1:]...)
	if err != nil {
		log.Fatalln("Could not run makefile:", err.Error())
	}
}