package path

import "github.com/adrg/xdg"

// Gets the file path that the data file will be saved to for the application
func DataFilePath() (string, error) {
	return xdg.DataFile("go_tri/.tridos.json")
}
