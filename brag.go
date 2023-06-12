package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Entry struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

type Brag struct {
	Entries []Entry `json:"entries"`
}

var bragFilePath string

func main() {
	initFlags()
	handleCommand()
}

func initFlags() {
	flag.StringVar(&bragFilePath, "file", "brag.json", "Path to the brag document file")
	flag.Parse()
}

func handleCommand() {
	if flag.NArg() == 0 {
		printHelp()
		return
	}

	command := flag.Arg(0)

	switch command {
	case "init":
		initBragDocument()
	case "add":
		addEntry()
	case "list":
		listEntries()
	case "remove":
		removeEntry()
	case "clear":
		clearEntries()
	case "export":
		exportEntries()
	case "help":
		printHelp()
	default:
		fmt.Println("Invalid command. Use 'brag help' to see the available commands.")
	}
}

func initBragDocument() {
	_, err := os.Stat(bragFilePath)
	if err == nil {
		fmt.Println("Brag document already exists.")
		return
	}

	file, err := os.Create(bragFilePath)
	if err != nil {
		fmt.Println("Error creating brag document:", err)
		return
	}
	defer file.Close()

	fmt.Println("Brag document initialized successfully.")
}

func addEntry() {
	brag, err := readBragDocument()
	if err != nil {
		fmt.Println("Error reading brag document:", err)
		return
	}

	text := strings.Join(flag.Args()[1:], " ")

	entry := Entry{
		ID:   len(brag.Entries) + 1,
		Text: text,
	}

	brag.Entries = append(brag.Entries, entry)

	err = writeBragDocument(brag)
	if err != nil {
		fmt.Println("Error adding entry:", err)
		return
	}

	fmt.Println("Entry added successfully.")
}

func listEntries() {
	brag, err := readBragDocument()
	if err != nil {
		fmt.Println("Error reading brag document:", err)
		return
	}

	if len(brag.Entries) == 0 {
		fmt.Println("No entries found.")
		return
	}

	for _, entry := range brag.Entries {
		fmt.Printf("[%d] %s\n", entry.ID, entry.Text)
	}
}

func removeEntry() {
	brag, err := readBragDocument()
	if err != nil {
		fmt.Println("Error reading brag document:", err)
		return
	}

	if len(brag.Entries) == 0 {
		fmt.Println("No entries to remove.")
		return
	}

	if flag.NArg() < 2 {
		fmt.Println("Please provide the ID of the entry to remove.")
		return
	}

	entryID, err := strconv.Atoi(flag.Arg(1))
	if err != nil {
		fmt.Println("Invalid entry ID.")
		return
	}

	var updatedEntries []Entry
	entryRemoved := false

	for _, entry := range brag.Entries {
		if entry.ID != entryID {
			updatedEntries = append(updatedEntries, entry)
		} else {
			entryRemoved = true
		}
	}

	if !entryRemoved {
		fmt.Println("Entry not found.")
		return
	}

	brag.Entries = updatedEntries

	err = writeBragDocument(brag)
	if err != nil {
		fmt.Println("Error removing entry:", err)
		return
	}

	fmt.Println("Entry removed successfully.")
}

func clearEntries() {
	err := os.Remove(bragFilePath)
	if err != nil {
		fmt.Println("Error clearing entries:", err)
		return
	}

	fmt.Println("Entries cleared successfully.")
}

func exportEntries() {
	brag, err := readBragDocument()
	if err != nil {
		fmt.Println("Error reading brag document:", err)
		return
	}

	if len(brag.Entries) == 0 {
		fmt.Println("No entries to export.")
		return
	}

	fileFormat := flag.Arg(1)

	switch fileFormat {
	case "txt":
		exportToTXT(brag)
	case "csv":
		exportToCSV(brag)
	case "json":
		exportToJSON(brag)
	default:
		fmt.Println("Invalid file format.")
	}
}

func exportToTXT(brag Brag) {
	file, err := os.Create("brag.txt")
	if err != nil {
		fmt.Println("Error exporting entries:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	for _, entry := range brag.Entries {
		fmt.Fprintf(writer, "[%d] %s\n", entry.ID, entry.Text)
	}

	writer.Flush()

	fmt.Println("Entries exported to brag.txt successfully.")
}

func exportToCSV(brag Brag) {
	file, err := os.Create("brag.csv")
	if err != nil {
		fmt.Println("Error exporting entries:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	fmt.Fprintln(writer, "ID,Text")

	for _, entry := range brag.Entries {
		fmt.Fprintf(writer, "%d,%s\n", entry.ID, entry.Text)
	}

	writer.Flush()

	fmt.Println("Entries exported to brag.csv successfully.")
}

func exportToJSON(brag Brag) {
	file, err := os.Create("brag.json")
	if err != nil {
		fmt.Println("Error exporting entries:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	err = encoder.Encode(brag)
	if err != nil {
		fmt.Println("Error exporting entries:", err)
		return
	}

	fmt.Println("Entries exported to brag.json successfully.")
}

func readBragDocument() (Brag, error) {
	var brag Brag

	file, err := ioutil.ReadFile(bragFilePath)
	if err != nil {
		return brag, err
	}

	err = json.Unmarshal(file, &brag)
	if err != nil {
		return brag, err
	}

	return brag, nil
}

func writeBragDocument(brag Brag) error {
	data, err := json.MarshalIndent(brag, "", "  ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(bragFilePath, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func printHelp() {
	fmt.Println("Usage: brag <command> [arguments]")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  init            Initializes the brag document")
	fmt.Println("  add <entry>     Adds a new entry to the brag document")
	fmt.Println("  list            Lists all entries in the brag document")
	fmt.Println("  remove <id>     Removes the entry with the specified ID")
	fmt.Println("  clear           Clears all entries from the brag document")
	fmt.Println("  export <format> Exports the brag document to the specified format (txt, csv, json)")
	fmt.Println("  help            Displays help information")
}
