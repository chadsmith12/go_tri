package path

import "github.com/adrg/xdg"

// Gets the file path that the data file will be saved to for the application
func DataFilePath() (string, error) {
	return xdg.DataFile("go_tri/.tridos.json")
}

func ConfigFilePath() (string, error) {
	return xdg.ConfigFile("go_tri/.config.yml")
}

// This will be the path that the user saves the data file at.
var DataFile string
var ConfigFile string
