package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"time"
)

// Generates a string with format MMDDYYYYHHMMSS
func generateFileName() string {
	name := time.Now().Format("01022006150405")
	return name
}

// Runs Neovim with the buffer of the file name
func runNeovim(filename string) error {
	cmd := exec.Command("nvim", filename)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	return cmd.Run()

}

// Set the default path for a zet
// Currently only sets the env for the current process
// TODO: write $env:ZET_DIR = "path" to $PROFILE
func setZetDir(path string) error {
	//checks the path given to see if the directory exists and
	//if the path is a directory
	if s, err := os.Stat(path); !s.IsDir() || os.IsNotExist(err) {
		return fmt.Errorf(
			`Either the path given is not a directory or the \n
			path does not exist: %v`, err)
	}

	cmd :=
		os.Setenv("ZET_DIR", path)
	return nil
}

// Gets the env ZET_DIR
func getZetDir() (string, error) {

	dirPath := os.Getenv("ZET_DIR")

	if dirPath == "" {
		return dirPath, errors.New("No env ZET_DIR set")
	}
	//gets the Stat from the path and checks if it is a directory
	if s, err := os.Stat(dirPath); !s.IsDir() || err != nil {
		return dirPath, fmt.Errorf("Error getting defualt directory: %v", err)
	}
	return dirPath, nil
}

func main() {
	//Setting up a flag for setting the defualt env variable
	setZetFlag := flag.String("sd", "",
		"Sets the env varible 'ZET_DIR' to the path given")

	if *setZetFlag != "" {
		setZetDir(*setZetFlag)
	}

	//Get the env ZET_DIR
	zetDir, err := getZetDir()

	if err != nil {
		fmt.Printf("Error getting ZET_DIR: %v", err)
	}

	zetName := generateFileName()
	fmt.Printf("Generated File Name: %s", zetName)

	//open Neovim with buffer name zetName
	if err := runNeovim(zetName); err != nil {
		fmt.Printf("Error opening neovim: %s", err)
	}

}
