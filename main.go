package main

import (
	"os"
	"path/filepath"

	"github.com/kakaw2016/MergeRepertoriesProject/configurationFile"
	"github.com/kakaw2016/MergeRepertoriesProject/regroupingdirectoriesfiles"
)

func main() {

	currentDir, _ := os.Getwd()
	configPath := filepath.Join(currentDir, "configurationfile.txt")
	relativePath, _ := filepath.Rel(currentDir, configPath)

	repertory1, repertory2, mergedRepertory, _ := configurationFile.ReadConfigurationFile(relativePath)

	regroupingdirectoriesfiles.EraseDataInRepertory(mergedRepertory)

	regroupingdirectoriesfiles.BrowsingFirstRepertory(repertory1, mergedRepertory)

	regroupingdirectoriesfiles.BrowsingSecondRepertory(repertory2, mergedRepertory)

}
