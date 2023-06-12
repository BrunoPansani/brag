package brag

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

// Entry represents an entry in the brag document.
type Entry struct {
	ID        int       `json:"id"`
	Timestamp time.Time `json:"timestamp"`
	Text      string    `json:"text"`
}

// BragDocument represents the brag document.
type BragDocument struct {
	Entries []Entry `json:"entries"`
}

const (
	bragDocumentFilePath = "/data/brag.json"
	dataPath             = "/data/"
)

func readBragDocument() (*BragDocument, error) {
	bragDocumentPath := getBragDocumentPath()

	file, err := os.Open(bragDocumentPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	byteData, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	brag := &BragDocument{}
	err = json.Unmarshal(byteData, brag)
	if err != nil {
		return nil, err
	}

	return brag, nil
}

func writeBragDocument(brag *BragDocument) error {
	byteData, err := json.MarshalIndent(brag, "", "  ")
	if err != nil {
		return err
	}

	bragDocumentPath := getBragDocumentPath()

	err = ioutil.WriteFile(bragDocumentPath, byteData, 0644)
	if err != nil {
		return err
	}

	return nil
}

// fileExists checks if a file exists at the specified path
func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func getBragDocumentPath(fileName ...string) string {
	executablePath, err := os.Executable()
	if err != nil {
		fmt.Println("Error retrieving executable path:", err)
		os.Exit(1)
	}

	if len(fileName) > 0 {
		return filepath.Join(filepath.Dir(executablePath), dataPath, fileName[0])
	}

	return filepath.Join(filepath.Dir(executablePath), bragDocumentFilePath)
}

// InitBragDocument initializes the brag document file if it doesn't exist
func InitBragDocument() error {

	bragDocumentPath := getBragDocumentPath()

	if fileExists(bragDocumentPath) {
		fmt.Println("Brag document already exists.")
		return nil
	}

	file, err := os.Create(bragDocumentPath)

	if err != nil {
		fmt.Println("Error creating brag document:", err)
		return err
	}

	brag := &BragDocument{
		Entries: []Entry{},
	}

	err = writeBragDocument(brag)
	if err != nil {
		fmt.Println("Error writing to brag document:", err)
		return err
	}
	defer file.Close()

	fmt.Println("Brag document created successfully.")
	return nil
}

func AddEntry(entryText string) error {
	brag, err := readBragDocument()
	if err != nil {
		return err
	}

	entry := Entry{
		ID:        len(brag.Entries) + 1,
		Timestamp: time.Now(),
		Text:      entryText,
	}

	brag.Entries = append(brag.Entries, entry)

	err = writeBragDocument(brag)
	if err != nil {
		return err
	}

	return nil
}

func RemoveEntry(entryID int) error {
	brag, err := readBragDocument()
	if err != nil {
		return err
	}

	if entryID <= 0 || entryID > len(brag.Entries) {
		return fmt.Errorf("invalid entry ID")
	}

	brag.Entries = append(brag.Entries[:entryID-1], brag.Entries[entryID:]...)

	err = writeBragDocument(brag)
	if err != nil {
		return err
	}

	return nil
}

func ClearEntries() error {
	brag := &BragDocument{
		Entries: []Entry{},
	}

	err := writeBragDocument(brag)
	if err != nil {
		return err
	}

	return nil
}

func ListEntries() {
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
		fmt.Printf("#%d [%s] %s\n", entry.ID, entry.Timestamp.Format("2006-01-02 15:04:05"), entry.Text)
	}
}

func ExportEntries(fileFormat string) {
	brag, err := readBragDocument()
	if err != nil {
		fmt.Println("Error reading brag document:", err)
		return
	}

	if len(brag.Entries) == 0 {
		fmt.Println("No entries to export.")
		return
	}

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

func exportToTXT(brag *BragDocument) {
	exportPath := getBragDocumentPath("brag.txt")
	file, err := os.Create(exportPath)
	if err != nil {
		fmt.Println("Error exporting entries:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	for _, entry := range brag.Entries {
		fmt.Fprintf(writer, "#%d [%s] %s\n", entry.ID, entry.Timestamp.Format("2006-01-02 15:04:05"), entry.Text)
	}

	writer.Flush()

	fmt.Println("Entries exported to brag.txt successfully.")
}

func exportToCSV(brag *BragDocument) {
	exportPath := getBragDocumentPath("brag.csv")
	file, err := os.Create(exportPath)
	if err != nil {
		fmt.Println("Error exporting entries:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	fmt.Fprintln(writer, "ID,Timestamp,Text")

	for _, entry := range brag.Entries {
		fmt.Fprintf(writer, "%d,%s,%s\n", entry.ID, entry.Timestamp.Format("2006-01-02 15:04:05"), entry.Text)
	}

	writer.Flush()

	fmt.Println("Entries exported to brag.csv successfully.")
}

func exportToJSON(brag *BragDocument) {
	exportPath := getBragDocumentPath("brag.json")
	file, err := os.Create(exportPath)
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
