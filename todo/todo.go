package todo

import (
	"encoding/json"
	"os"
)
type Item struct {
	Text string
}

func SaveItems(filename string, items []Item) error {
	data, jsonErr := json.Marshal(items)
	if jsonErr != nil {
		return jsonErr
	}

	fileErr := os.WriteFile(filename, data, 0644)
	if fileErr != nil {
		return fileErr
	}

	return nil
}

func ReadItems(filename string) ([]Item, error) {
	data, readErr := os.ReadFile(filename)
	if readErr != nil {
		return []Item{}, readErr
	}

	var items []Item
	if err := json.Unmarshal(data, &items); err != nil {
		return []Item{}, err
	}

	return items, nil
}
