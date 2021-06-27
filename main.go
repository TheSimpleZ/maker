package main

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
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
	makefilePath := filepath.Join(cwd, "Makefile")
	hasMakefile, err := exists(makefilePath)
	if err != nil {
		return "", err
	}

	if hasMakefile {
		commandExists, err := makefileHasTarget(target, makefilePath)
		if err != nil {
			return "", err
		}

		if commandExists {
			return cwd, nil
		}
	}

	isGitRoot, err := exists(filepath.Join(cwd, ".git"))
	if err != nil {
		return "", err
	}


	if !isGitRoot && cwd != "/" {
		parentDir := filepath.Dir(cwd)
		return findMakeFileWithTarget(parentDir, target)
	} else {
		return "", fmt.Errorf("Stopped at %v and found no make target '%v'", cwd, target)
	}
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

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalln("Could not get current working directory:", err)
	}

	makeFileDir, err := findMakeFileWithTarget(cwd, os.Args[1])
	if err != nil {
		log.Fatalln("Could not find a makefile:", err)
	}

	err = runMake(makeFileDir, os.Args[1:]...)
	if err != nil {
		log.Fatalf("Could not run makefile in %v: %v\n", makeFileDir, err)
	}
}