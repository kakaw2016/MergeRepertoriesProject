package regroupingdirectoriesfiles

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func EraseDataInRepertory(mergedRepertory string) {
	error1 := os.RemoveAll(mergedRepertory)
	if error1 != nil {
		fmt.Println("Error1", error1)
	}

	error2 := os.MkdirAll(mergedRepertory, 0755)
	if error2 != nil {
		fmt.Println("Error2", error2)
	}

}

func BrowsingFirstRepertory(repertory1 string, mergedRepertory string) {

	// Walk the directories and their subdirectories

	filepath.Walk(repertory1, func(pathRepertory1 string, info os.FileInfo, err error) error {

		// Calculate the destination path for the file or directory
		dstPath := filepath.Join(mergedRepertory, pathRepertory1[len(repertory1):])
		if info.IsDir() {
			// If it's a directory, create it in the merged directory
			return os.MkdirAll(dstPath, info.Mode())
		} else {
			// If it's a file, copy it to the merged directory if it's newer or doesn't exist
			srcModTime := info.ModTime()
			dstInfo, err := os.Stat(dstPath)
			if err == nil {
				dstModTime := dstInfo.ModTime()
				if !srcModTime.After(dstModTime) {
					return nil
				}
			}
			return copyFile(pathRepertory1, dstPath)
		}

	})
}

func BrowsingSecondRepertory(repertory2 string, mergedRepertory string) {

	filepath.Walk(repertory2, func(pathRepertory2 string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// Calculate the destination path for the file or directory
		dstPath := filepath.Join(mergedRepertory, pathRepertory2[len(repertory2):])
		if info.IsDir() {
			// If it's a directory, create it in the merged directory
			return os.MkdirAll(dstPath, info.Mode())
		} else {
			// If it's a file, copy it to the merged directory if it's newer or doesn't exist

			_ = copyCheckFile(pathRepertory2, dstPath)
			srcModTime := info.ModTime()
			dstInfo, err := os.Stat(dstPath)
			if err == nil {
				dstModTime := dstInfo.ModTime()
				if !srcModTime.After(dstModTime) {
					return nil
				}
			}
			return copyFile(pathRepertory2, dstPath)
		}

	})

}

//Check the destination file contents and compare similar base files name with the second repository.
//If similar base name have different contents save both under different names in destination repertory.

func copyCheckFile(src2, dstPath string) error {

	var fileNames []string
	dstInfo := make(map[string]string)

	filepath.Walk(dstPath, func(pathDestination string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}
		fileNames = append(fileNames, info.Name())

		dstInfo[info.Name()] = pathDestination

		return nil
	})

	for _, fNames := range fileNames {
		if filepath.Base(src2) == fNames {
			// Read the contents of the first file
			bytes1, err := ioutil.ReadFile(src2)

			if err != nil {
				fmt.Println(err)

			}
			// Read the contents of the second file
			bytes2, err := ioutil.ReadFile(dstInfo[fNames])

			if err != nil {
				fmt.Println(err)

			}
			if len(bytes1) != len(bytes2) {
				for i := 0; i < len(bytes1); i++ {
					if bytes1[i] != bytes2[i] {

						// Rename the first file
						dir1, file1 := filepath.Split(dstInfo[fNames])

						ext1 := filepath.Ext(file1)
						base1 := file1[:len(file1)-len(ext1)]

						newFile1 := base1 + "_new1" + ext1
						_ = os.Rename(dstInfo[fNames], filepath.Join(dir1, newFile1))

						if err != nil {
							fmt.Println(err)

						}

					}
				}
			}

		}

	}
	return nil
}

//Copy the file from one place to the second location

func copyFile(src, dst string) error {
	// Read the file at src
	data, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}

	// Write the data to the file at dst
	return ioutil.WriteFile(dst, data, 0644)
}
