package todo

import (
	"encoding/json"
	"errors"
	"os"
	"strconv"
)
type Item struct {
	Text string
	Priority int
	Done bool
	position int
}

// ByPriority implements the sort.Interface for []Item based 
// on the Priority and position field
type ByPriority []Item

func (items ByPriority) Len() int {
	return len(items)
}

func (items ByPriority) Swap(i, j int) {
	items[i], items[j] = items[j], items[i]
}

func (items ByPriority) Less(i, j int) bool {
	if items[i].Done != items[j].Done {
		return items[i].Done
	}

	if items[i].Priority == items[j].Priority {
		return items[i].Priority < items[j].Priority
	}
	return items[i].position < items[j].position
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

func (item *Item) PrettyDone() string {
	if item.Done {
		return "X"
	}

	return ""
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
	if errors.Is(readErr, os.ErrNotExist) {
		// files that don't exist might be new. ignore this error and return empty
		return []Item{}, nil
	}	
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
