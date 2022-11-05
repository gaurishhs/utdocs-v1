package config

import (
	"encoding/json"
	"os"
)

var (
	Configuration = Config{
		SiteName:  "UTDocs",
		CodeTheme: "atom-one-dark",
	}
)

func ReadConfig(filePath string) {
	// Reading JSON file
	file, err := os.ReadFile(filePath)
	if err != nil {
		return
	}

	// Unmarshalling JSON file
	err = json.Unmarshal(file, &Configuration)
	if err != nil {
		return
	}

}
