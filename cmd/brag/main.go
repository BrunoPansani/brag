package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/BrunoPansani/brag/internal/brag"
)

func main() {
	// Parse command-line flags
	flag.Parse()

	// Get the command and arguments
	args := flag.Args()
	if len(args) < 1 {
		printUsage()
		os.Exit(1)
	}

	// Check the command and execute the corresponding function
	switch args[0] {
	case "init":
		initBragDocument()
	case "add":
		if len(args) < 2 {
			printUsage()
			os.Exit(1)
		}
		entryText := args[1]
		err := addEntry(entryText)
		if err != nil {
			fmt.Println("Error adding entry:", err)
			os.Exit(1)
		}
		fmt.Println("Entry added successfully!")
	case "list":
		listEntries()
	case "remove":
		if len(args) < 2 {
			printUsage()
			os.Exit(1)
		}
		entryIDStr := args[1]
		entryID, err := strconv.Atoi(entryIDStr)
		if err != nil {
			fmt.Println("Invalid entry ID:", entryIDStr)
			os.Exit(1)
		}
		err = removeEntry(entryID)
		if err != nil {
			fmt.Println("Error removing entry:", err)
			os.Exit(1)
		}
		fmt.Println("Entry removed successfully!")
	case "clear":
		err := clearEntries()
		if err != nil {
			fmt.Println("Error clearing entries:", err)
			os.Exit(1)
		}
		fmt.Println("All entries cleared!")
	case "export":
		if len(args) < 2 {
			printUsage()
			os.Exit(1)
		}
		format := args[1]
		exportEntries(format)
	case "help":
		printUsage()
	default:
		fmt.Println("Invalid command. Use 'help' to see available commands.")
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("Usage: brag <command> [arguments]")
	fmt.Println()
	fmt.Println("Available commands:")
	fmt.Println("  init                 Initialize the brag document")
	fmt.Println("  add <entry> 			Add a new entry to the brag document")
	fmt.Println("  list                 List all entries in the brag document")
	fmt.Println("  remove <id>          Remove the entry with the specified ID")
	fmt.Println("  clear                Clear all entries from the brag document")
	fmt.Println("  export <format>      Export the brag document to the specified format (txt, csv, json)")
	fmt.Println("  help                 Display help information")
}

func initBragDocument() {
	fmt.Println("Initializing the brag document...")
	brag.InitBragDocument()
}

func addEntry(entryText string) error {
	return brag.AddEntry(entryText)
}

func listEntries() {
	brag.ListEntries()
}

func removeEntry(entryID int) error {
	return brag.RemoveEntry(entryID)
}

func clearEntries() error {
	return brag.ClearEntries()
}

func exportEntries(format string) {
	brag.ExportEntries(format)
}
