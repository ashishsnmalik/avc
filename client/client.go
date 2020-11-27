package main

import (
	"fmt"
	"os"
)

// This will validate various command line options from the user.
// Work according to the commands selected.
// Initially decided to use the flag option but that is not working for Unix wildcard expansion like *.
// I need a command to be able to interpret values like avc add * .. (that should include all the files in the current directory).

func main() {
	mainCommand := os.Args[1]
	fmt.Printf("Main Command is [%s]\n", mainCommand)

	if mainCommand == "" {

		mainCommand = "help"
	}
	switch mainCommand {
	case "add":
		fmt.Println("Main command is add.")
		fmt.Println(os.Args[2:])
		addFile(os.Args[2:])
	case "status":
		fmt.Println("Main command is status.")
	case "show-ref":
		fmt.Println("Main command is show-ref.")
	case "cat-file":
		fmt.Println("Main command is cat-file.")
	case "checkout":
		fmt.Println("Main command is checkout.")
	case "commit":
		fmt.Println("Main command is commit.")
	case "init":
		fmt.Println("Main command is init.")
	case "hash-object":
		fmt.Println("Main command is hash-object.")
	case "log":
		fmt.Println("Main command is log.")
	case "ls-tree":
		fmt.Println("Main command is ls-tree.")
	case "merge":
		fmt.Println("Main command is merge.")
	case "rebase":
		fmt.Println("Main command is rebase.")
	case "rev-parse":
		fmt.Println("Main command is rev-parse.")
	case "rm":
		fmt.Println("Main command is rm.")
	case "tag":
		fmt.Println("Main command is tag.")
	case "help":
		fmt.Println("Main command is help.")
		printHelp()
	default:
		printHelp()
	}

}

func printHelp() {
	fmt.Println("This is the HELP PRINT")
}

func addFile(str []string) {
	fmt.Println("\n\nIn Function addFile")
	fmt.Println(str)

	fp, err := os.Open(str[0])
}
