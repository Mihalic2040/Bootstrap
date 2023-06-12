package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/Mihalic2040/Hub/src/types"
)

func create() {
	config := types.Config{
		Host:       "0.0.0.0",
		Port:       "6666",
		DHTServer:  true,
		ProtocolId: "/hub/0.0.1",
		Rendezvous: "Hub",
		Secret:     "BOOTSTRAP",
	}

	// Convert the config to JSON format
	jsonData, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		log.Println("[*]", "Failed to marshal JSON: %v", err)
		return
	}

	// Write the JSON data to a file
	err = ioutil.WriteFile("config.json", jsonData, 0644)
	if err != nil {
		log.Println("[*]", "Failed to write file: %v", err)
		return
	}

	log.Println("[*]", "Configuration file created successfully.")
}

func Read() types.Config {
	configFilePath := "config.json"
	_, err := os.Stat(configFilePath)
	if os.IsNotExist(err) {
		create()
	}

	// Read the JSON file
	data, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		fmt.Printf("Failed to read file: %v", err)
		return types.Config{}
	}

	config := types.Config{}
	// Unmarshal JSON data into the Config struct
	err = json.Unmarshal(data, &config)
	if err != nil {
		fmt.Printf("Failed to unmarshal JSON: %v", err)
		return types.Config{}
	}

	return config
}
