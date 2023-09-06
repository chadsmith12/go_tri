package todo

import (
	"encoding/json"
	"os"
	"strconv"
)
type Item struct {
	Text string
	Priority int
	position int
}

func (item *Item) SetPriority(priority int) {
	switch priority {
	case 1: 
		item.Priority = 1
	case 3:
		item.Priority = 3
	default:
		item.Priority = 2
	}
}

func (item *Item) PrettyP() string {
	if item.Priority == 1 {
		return "(1)"
	}
	if item.Priority == 3 {
		return "(3)"
	}

	return " "
}

func (item *Item) Label() string {
	return strconv.Itoa(item.position) + "."
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
	
	for i := range items {
		items[i].position = i + 1
	}
	return items, nil
}
