package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

func runMake(dir string, args ...string) error {
	dirCommand := []string{"-C", dir}
	allArgs := append(dirCommand, args...)
	cmd := exec.Command("make",  allArgs...)

    stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}

    if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

    scannerOut := bufio.NewScanner(stdout)
    scannerOut.Split(bufio.ScanLines)

	scannerErr := bufio.NewScanner(stderr)
    scannerErr.Split(bufio.ScanLines)

	go func() {
		for scannerOut.Scan() {
			m := scannerOut.Text()
			fmt.Println(m)
		}
	}()

	for scannerErr.Scan() {
        m := scannerErr.Text()
        fmt.Println(m)
    }



    if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}

	return nil
}



func checkIfCommandExists(command string, path string) (bool, error) {
	fileContent, err := ioutil.ReadFile(path)

	if os.IsNotExist(err) {
		return false, nil
	} else if err != nil {
		return false, err
	}
	fileContentString := string(fileContent)
	stringMatch := fmt.Sprintf("PHONY:%s\n%s:", command, command)
	if strings.Contains(strings.ReplaceAll(fileContentString, " ", ""), stringMatch) {
		return true, nil
	}

	return false, nil
}


func runRecurse(cwd string) error {
	commandExists, err := checkIfCommandExists(os.Args[1], cwd + "Makefile")
	if err != nil {
		return err
	} else if commandExists {
		err := runMake(cwd, os.Args[1:]...)
		return err
	} else if _, err := os.Stat(cwd + ".git"); os.IsNotExist(err) {
		err := runRecurse("./." + cwd)
		if err != nil {
			return err
		}
	} else {
		return errors.New("Could not find a makefile with target: " + os.Args[1])
	}

	return nil
}

func main() {
	if len(os.Args[1:]) < 1 {
		fmt.Println("You must provide a make target")
        return
	}
	err := runRecurse("./")

	if err != nil {
		fmt.Println("An error occurred while running make:", err.Error())
	}
}