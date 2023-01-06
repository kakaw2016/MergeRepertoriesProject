package configurationFile

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Configuration details in file that generates all the input we will use to find the files

func ReadConfigurationFile(configurationFilepath string) (string, string, string, error) {

	// Open the config file in read-only mode
	configFile, err := os.Open(configurationFilepath)
	if err != nil {
		fmt.Println(err)

	}
	defer configFile.Close()

	// Create scanner for config.txt
	scanner := bufio.NewScanner(configFile)

	var pathA string
	var pathB string
	var pathC string

	// Create slice to store extensions

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "=")
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			if key == "repertoryA" {
				// Add value to each key
				pathA = value

			} else if key == "repertoryB" {

				pathB = value
			} else if key == "MergedRepertory" {

				pathC = value
			}
		}

	}

	// Check for errors while scanning
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return pathA, pathB, pathC, nil
}
