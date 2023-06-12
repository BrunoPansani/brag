package brag

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
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
	bragDocumentPath = "../data/brag.json"
)

func readBragDocument() (*BragDocument, error) {
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

	err = ioutil.WriteFile(bragDocumentPath, byteData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func addEntry(entryText string, eventDate time.Time) error {
	brag, err := readBragDocument()
	if err != nil {
		return err
	}

	entry := Entry{
		ID:        len(brag.Entries) + 1,
		Timestamp: time.Now(),
		Text:      entryText,
	}

	if !eventDate.IsZero() {
		entry.Timestamp = eventDate
	}

	brag.Entries = append(brag.Entries, entry)

	err = writeBragDocument(brag)
	if err != nil {
		return err
	}

	return nil
}

func removeEntry(entryID int) error {
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

func clearEntries() error {
	brag := &BragDocument{
		Entries: []Entry{},
	}

	err := writeBragDocument(brag)
	if err != nil {
		return err
	}

	return nil
}
