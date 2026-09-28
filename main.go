package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	dir := flag.String("dir", ".", "Directory to organize")
	dryRun := flag.Bool("dry-run", false, "Preview changes without moving files")
	flag.Parse()

	info, err := os.Stat(*dir)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	if !info.IsDir() {
		fmt.Println("Error: path is not a directory")
		os.Exit(1)
	}

	fmt.Println("Smart File Organizer")
	fmt.Println("---------------------")
	fmt.Println("Directory:", *dir)

	if *dryRun {
		fmt.Println("Mode: Preview")
	} else {
		fmt.Println("Mode: Organize")
	}

	err = Organize(*dir, *dryRun)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println("\nDone!")
}
