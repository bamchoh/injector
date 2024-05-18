package main

import (
	"fmt"
	"io/fs"
	"os"
	"os/exec"

	"github.com/joho/godotenv"
)

const (
	ExitCodeOK = iota
	ExitCodeInvalidArguments
	ExitCodeDotenvLoadFailed
	ExitCodeRunFailed
)

func run() int {
	if len(os.Args) < 2 {
		fmt.Println("You run me without any command line options. To run, please specify one or more parameters.")
		return ExitCodeInvalidArguments
	}

	if err := godotenv.Load(); err != nil {
		switch err.(type) {
		case *fs.PathError:
			break
		default:
			fmt.Println(err)
			return ExitCodeDotenvLoadFailed
		}
	}

	cmdname := os.Args[1]
	var args []string
	if len(os.Args[2:]) > 0 {
		args = append(args, os.Args[2:]...)
	}
	cmd := exec.Command(cmdname, args...)
	fmt.Printf("Command: %q\n", cmd.Args)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("Execution Failed:", err)
		return ExitCodeRunFailed
	}

	return ExitCodeOK
}

func main() {
	os.Exit(run())
}
